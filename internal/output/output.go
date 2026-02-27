package output

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"

	"github.com/olekukonko/tablewriter"
	"golang.org/x/term"
)

var format = "json"

func SetFormat(f string) {
	if f != "" {
		format = f
	}
}

func Format() string { return format }

// Print formats and prints response data based on the output format.
func Print(data []byte) error {
	if len(data) == 0 {
		return nil
	}

	switch format {
	case "raw":
		_, err := os.Stdout.Write(data)
		return err
	case "table":
		return printTable(data)
	case "csv":
		return printCSV(data)
	default:
		return printJSON(data)
	}
}

func printJSON(data []byte) error {
	var buf bytes.Buffer
	if err := json.Indent(&buf, data, "", "  "); err != nil {
		_, err := os.Stdout.Write(data)
		return err
	}
	buf.WriteByte('\n')
	_, err := buf.WriteTo(os.Stdout)
	return err
}

type valueKind int

const (
	kindScalar      valueKind = iota
	kindSimpleArray
	kindComplex
)

func classifyValue(v any) (string, valueKind) {
	switch val := v.(type) {
	case string:
		return val, kindScalar
	case float64:
		if val == float64(int64(val)) {
			return fmt.Sprintf("%d", int64(val)), kindScalar
		}
		return fmt.Sprintf("%g", val), kindScalar
	case bool:
		if val {
			return "true", kindScalar
		}
		return "false", kindScalar
	case nil:
		return "", kindScalar
	case []any:
		if len(val) == 0 {
			return "", kindScalar
		}
		parts := make([]string, 0, len(val))
		for _, elem := range val {
			switch e := elem.(type) {
			case string:
				parts = append(parts, e)
			case float64:
				if e == float64(int64(e)) {
					parts = append(parts, fmt.Sprintf("%d", int64(e)))
				} else {
					parts = append(parts, fmt.Sprintf("%g", e))
				}
			default:
				return "", kindComplex
			}
		}
		return strings.Join(parts, ", "), kindSimpleArray
	default:
		return "", kindComplex
	}
}

var acronyms = map[string]bool{
	"id": true, "url": true, "api": true, "sip": true,
	"mac": true, "ip": true, "uri": true,
	"esn": true, "pstn": true, "uuid": true, "dn": true,
	"did": true, "ani": true, "dnis": true, "ata": true,
	"http": true, "https": true, "html": true,
	"ivr": true, "vcc": true, "crm": true, "dnc": true,
	"csv": true, "ftp": true, "tts": true, "wav": true,
}

func camelToHeader(s string) string {
	if s == "" {
		return s
	}

	var words []string
	runes := []rune(s)
	start := 0
	for i := 1; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) && unicode.IsLower(runes[i-1]) {
			words = append(words, string(runes[start:i]))
			start = i
		} else if unicode.IsUpper(runes[i-1]) && unicode.IsUpper(runes[i]) && i+1 < len(runes) && unicode.IsLower(runes[i+1]) {
			words = append(words, string(runes[start:i]))
			start = i
		}
	}
	words = append(words, string(runes[start:]))

	for i, w := range words {
		lower := strings.ToLower(w)
		if acronyms[lower] {
			words[i] = strings.ToUpper(w)
		} else {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}

	return strings.Join(words, " ")
}

func getTerminalWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w <= 0 {
		return 120
	}
	return w
}

func extractRows(data []byte) ([]string, [][]string, bool, error) {
	var items []map[string]any

	if err := json.Unmarshal(data, &items); err != nil {
		var obj map[string]any
		if err2 := json.Unmarshal(data, &obj); err2 != nil {
			return nil, nil, false, fmt.Errorf("not tabular JSON")
		}

		found := false
		foundEmpty := false

		sortedKeys := make([]string, 0, len(obj))
		for k := range obj {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)

		for _, k := range sortedKeys {
			v := obj[k]
			if raw, err3 := json.Marshal(v); err3 == nil {
				var arr []map[string]any
				if err3 = json.Unmarshal(raw, &arr); err3 == nil {
					items = arr
					found = true
					if len(arr) == 0 {
						foundEmpty = true
					}
					break
				}
			}
		}

		if foundEmpty {
			return nil, nil, true, nil
		}

		if !found {
			items = []map[string]any{obj}
		}
	}

	if len(items) == 0 {
		return nil, nil, true, nil
	}

	keySet := map[string]bool{}
	for _, item := range items {
		for k := range item {
			keySet[k] = true
		}
	}

	allKeys := make([]string, 0, len(keySet))
	for k := range keySet {
		allKeys = append(allKeys, k)
	}
	sort.Strings(allKeys)

	type cellInfo struct {
		text string
		kind valueKind
	}
	rowCells := make([]map[string]cellInfo, len(items))
	columnHasScalar := map[string]bool{}

	for i, item := range items {
		rowCells[i] = make(map[string]cellInfo, len(allKeys))
		for _, k := range allKeys {
			val, ok := item[k]
			if !ok {
				rowCells[i][k] = cellInfo{"", kindScalar}
				continue
			}
			text, kind := classifyValue(val)
			rowCells[i][k] = cellInfo{text, kind}
			if kind != kindComplex {
				columnHasScalar[k] = true
			}
		}
	}

	var cols []string
	for _, k := range allKeys {
		if columnHasScalar[k] {
			cols = append(cols, k)
		}
	}

	if len(cols) == 0 {
		return nil, nil, false, nil
	}

	headers := make([]string, len(cols))
	for i, k := range cols {
		headers[i] = camelToHeader(k)
	}

	rows := make([][]string, len(items))
	for i, cells := range rowCells {
		row := make([]string, len(cols))
		for j, k := range cols {
			row[j] = cells[k].text
		}
		rows[i] = row
	}

	return headers, rows, false, nil
}

func printTable(data []byte) error {
	headers, rows, empty, err := extractRows(data)
	if err != nil {
		return printJSON(data)
	}
	if empty {
		fmt.Println("(empty)")
		return nil
	}
	if headers == nil {
		return printJSON(data)
	}

	termWidth := getTerminalWidth()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetColWidth(termWidth)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)

	borderOverhead := len(headers) + 1 + 2*len(headers)
	availWidth := max(termWidth-borderOverhead, len(headers)*4)
	maxColWidth := max(availWidth/len(headers), 4)

	for i, row := range rows {
		for j, cell := range row {
			if len(cell) > maxColWidth {
				rows[i][j] = cell[:maxColWidth-3] + "..."
			}
		}
	}

	table.AppendBulk(rows)
	table.Render()

	return nil
}

func printCSV(data []byte) error {
	headers, rows, empty, err := extractRows(data)
	if err != nil {
		return printJSON(data)
	}
	if empty {
		return nil
	}
	if headers == nil {
		return printJSON(data)
	}

	w := csv.NewWriter(os.Stdout)
	if err := w.Write(headers); err != nil {
		return err
	}
	for _, row := range rows {
		if err := w.Write(row); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}
