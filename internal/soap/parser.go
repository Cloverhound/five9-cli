package soap

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// ParseResponse extracts the response body from a SOAP XML response,
// navigating to Envelope > Body > <method>Response > return,
// and converts it to JSON bytes.
func ParseResponse(data []byte, method string) ([]byte, error) {
	responseTag := method + "Response"

	decoder := xml.NewDecoder(strings.NewReader(string(data)))

	// Navigate into Envelope > Body > <method>Response
	if err := seekElement(decoder, "Envelope"); err != nil {
		return nil, fmt.Errorf("no SOAP Envelope: %w", err)
	}
	if err := seekElement(decoder, "Body"); err != nil {
		return nil, fmt.Errorf("no SOAP Body: %w", err)
	}
	if err := seekElement(decoder, responseTag); err != nil {
		return nil, fmt.Errorf("no %s element: %w", responseTag, err)
	}

	// Parse the content inside the response element.
	// There may be zero, one, or multiple <return> elements.
	result, err := parseChildren(decoder, responseTag)
	if err != nil {
		return nil, err
	}

	if result == nil {
		return json.Marshal(struct{}{})
	}

	// If there's a "return" key, use its value directly
	if ret, ok := result["return"]; ok {
		return json.Marshal(ret)
	}

	return json.Marshal(result)
}

// seekElement advances the decoder to the start of the named element.
func seekElement(decoder *xml.Decoder, localName string) error {
	for {
		tok, err := decoder.Token()
		if err != nil {
			return err
		}
		if se, ok := tok.(xml.StartElement); ok {
			if se.Name.Local == localName {
				return nil
			}
		}
	}
}

// parseChildren parses all child elements within the current element (up to its end tag).
// Returns a map of child element names to their values.
// If multiple children share the same name, their values are collected into an array.
func parseChildren(decoder *xml.Decoder, parentLocal string) (map[string]any, error) {
	result := make(map[string]any)
	var textBuf strings.Builder
	hasChildren := false

	for {
		tok, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			hasChildren = true
			childName := t.Name.Local
			childVal, err := parseElement(decoder, childName)
			if err != nil {
				return nil, err
			}

			if existing, ok := result[childName]; ok {
				// Already have this key — convert to or append to array
				switch arr := existing.(type) {
				case []any:
					result[childName] = append(arr, childVal)
				default:
					result[childName] = []any{existing, childVal}
				}
			} else {
				result[childName] = childVal
			}

		case xml.CharData:
			textBuf.Write(t)

		case xml.EndElement:
			if t.Name.Local == parentLocal {
				if !hasChildren {
					// Leaf element — return nil to signal text content
					return nil, nil
				}
				if len(result) == 0 {
					return nil, nil
				}
				return result, nil
			}
		}
	}

	return result, nil
}

// parseElement parses a single element and its content (recursive).
func parseElement(decoder *xml.Decoder, localName string) (any, error) {
	var textBuf strings.Builder
	hasChildren := false
	children := make(map[string]any)

	for {
		tok, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := tok.(type) {
		case xml.StartElement:
			hasChildren = true
			childName := t.Name.Local
			childVal, err := parseElement(decoder, childName)
			if err != nil {
				return nil, err
			}

			if existing, ok := children[childName]; ok {
				switch arr := existing.(type) {
				case []any:
					children[childName] = append(arr, childVal)
				default:
					children[childName] = []any{existing, childVal}
				}
			} else {
				children[childName] = childVal
			}

		case xml.CharData:
			textBuf.Write(t)

		case xml.EndElement:
			if t.Name.Local == localName {
				if hasChildren {
					return children, nil
				}
				return autoType(strings.TrimSpace(textBuf.String())), nil
			}
		}
	}
}

// autoType converts string values to Go types where appropriate.
func autoType(s string) any {
	if s == "" {
		return s
	}
	if s == "true" {
		return true
	}
	if s == "false" {
		return false
	}
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return s
}
