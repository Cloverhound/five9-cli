package output

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func captureStdout(fn func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	fn()
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	buf.ReadFrom(r)
	return buf.String()
}

func TestCamelToHeader(t *testing.T) {
	tests := []struct {
		input, want string
	}{
		{"displayName", "Display Name"},
		{"firstName", "First Name"},
		{"isActive", "Is Active"},
		{"id", "ID"},
		{"email", "Email"},
		{"ivrScripts", "IVR Scripts"},
		{"vccConfiguration", "VCC Configuration"},
		{"crmSettings", "CRM Settings"},
		{"dncList", "DNC List"},
		{"csvData", "CSV Data"},
	}
	for _, tt := range tests {
		got := camelToHeader(tt.input)
		if got != tt.want {
			t.Errorf("camelToHeader(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestClassifyValue(t *testing.T) {
	text, kind := classifyValue("hello")
	if kind != kindScalar || text != "hello" {
		t.Errorf("string: got %q %v", text, kind)
	}

	text, kind = classifyValue(float64(42))
	if kind != kindScalar || text != "42" {
		t.Errorf("int: got %q %v", text, kind)
	}

	text, kind = classifyValue(true)
	if kind != kindScalar || text != "true" {
		t.Errorf("bool: got %q %v", text, kind)
	}

	text, kind = classifyValue(nil)
	if kind != kindScalar || text != "" {
		t.Errorf("nil: got %q %v", text, kind)
	}

	text, kind = classifyValue([]any{"a", "b", "c"})
	if kind != kindSimpleArray || text != "a, b, c" {
		t.Errorf("string array: got %q %v", text, kind)
	}

	text, kind = classifyValue([]any{float64(1), float64(2)})
	if kind != kindSimpleArray || text != "1, 2" {
		t.Errorf("number array: got %q %v", text, kind)
	}

	_, kind = classifyValue(map[string]any{"key": "val"})
	if kind != kindComplex {
		t.Errorf("object: got kind %v, want kindComplex", kind)
	}

	_, kind = classifyValue([]any{map[string]any{"key": "val"}})
	if kind != kindComplex {
		t.Errorf("array of objects: got kind %v, want kindComplex", kind)
	}
}

func TestPrintTableArray(t *testing.T) {
	SetFormat("table")
	data := []byte(`[{"name":"Sales","id":1,"description":"Sales skill"},{"name":"Support","id":2,"description":"Support skill"}]`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "Name") {
		t.Errorf("expected 'Name' header, got:\n%s", out)
	}
	if !strings.Contains(out, "Sales") {
		t.Errorf("expected 'Sales' in output, got:\n%s", out)
	}
	if !strings.Contains(out, "|") {
		t.Errorf("expected table borders '|', got:\n%s", out)
	}
}

func TestPrintTableSingleObject(t *testing.T) {
	SetFormat("table")
	data := []byte(`{"name":"Charlie","email":"charlie@example.com","isActive":true}`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "Name") {
		t.Errorf("expected 'Name' header, got:\n%s", out)
	}
	if !strings.Contains(out, "Charlie") {
		t.Errorf("expected 'Charlie' in output, got:\n%s", out)
	}
}

func TestPrintTableExcludesComplexColumns(t *testing.T) {
	SetFormat("table")
	data := []byte(`[{"name":"Alice","details":{"nested":"obj"}}]`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "Name") {
		t.Errorf("expected 'Name' header, got:\n%s", out)
	}
	if strings.Contains(out, "Details") {
		t.Errorf("should not contain complex column 'Details', got:\n%s", out)
	}
}

func TestPrintTableEmpty(t *testing.T) {
	SetFormat("table")
	data := []byte(`{"items":[]}`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "(empty)") {
		t.Errorf("expected '(empty)' for empty items, got:\n%s", out)
	}
}

func TestPrintTableSimpleArrayIncluded(t *testing.T) {
	SetFormat("table")
	data := []byte(`[{"name":"Alice","skills":["Sales","Support"]}]`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "Skills") {
		t.Errorf("expected 'Skills' header for simple array, got:\n%s", out)
	}
	if !strings.Contains(out, "Sales, Support") {
		t.Errorf("expected joined skills, got:\n%s", out)
	}
}

// --- CSV tests ---

func TestPrintCSVArray(t *testing.T) {
	SetFormat("csv")
	data := []byte(`[{"name":"Sales","id":1},{"name":"Support","id":2}]`)
	out := captureStdout(func() {
		Print(data)
	})
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines (header + 2 rows), got %d:\n%s", len(lines), out)
	}
	if !strings.Contains(lines[0], "Name") {
		t.Errorf("expected 'Name' in header, got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "Sales") {
		t.Errorf("expected 'Sales' in first data row, got: %s", lines[1])
	}
}

func TestPrintCSVSingleObject(t *testing.T) {
	SetFormat("csv")
	data := []byte(`{"name":"Charlie","email":"charlie@example.com","isActive":true}`)
	out := captureStdout(func() {
		Print(data)
	})
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines (header + 1 row), got %d:\n%s", len(lines), out)
	}
	if !strings.Contains(lines[0], "Name") {
		t.Errorf("expected 'Name' in header, got: %s", lines[0])
	}
	if !strings.Contains(lines[1], "Charlie") {
		t.Errorf("expected 'Charlie' in data row, got: %s", lines[1])
	}
}

func TestPrintCSVExcludesComplexColumns(t *testing.T) {
	SetFormat("csv")
	data := []byte(`[{"name":"Alice","details":{"nested":"obj"}}]`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, "Name") {
		t.Errorf("expected 'Name' header, got:\n%s", out)
	}
	if strings.Contains(out, "Details") {
		t.Errorf("should not contain complex column 'Details', got:\n%s", out)
	}
}

func TestPrintCSVEmpty(t *testing.T) {
	SetFormat("csv")
	data := []byte(`{"items":[]}`)
	out := captureStdout(func() {
		Print(data)
	})
	if out != "" {
		t.Errorf("expected empty output for empty items, got:\n%s", out)
	}
}

func TestPrintCSVQuoting(t *testing.T) {
	SetFormat("csv")
	data := []byte(`[{"name":"O'Brien, James","note":"He said \"hello\""}]`)
	out := captureStdout(func() {
		Print(data)
	})
	if !strings.Contains(out, `"O'Brien, James"`) {
		t.Errorf("expected quoted field with comma, got:\n%s", out)
	}
	if !strings.Contains(out, `"He said ""hello"""`) {
		t.Errorf("expected escaped double quotes, got:\n%s", out)
	}
}
