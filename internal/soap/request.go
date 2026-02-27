package soap

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// Request holds everything needed for a SOAP call.
type Request struct {
	Method string
	Params []Param
}

// NewRequest creates a new SOAP request for the given method.
func NewRequest(method string) *Request {
	return &Request{Method: method}
}

// SetParam adds a simple string parameter.
func (r *Request) SetParam(name, value string) {
	if value != "" {
		r.Params = append(r.Params, Param{Name: name, Value: value})
	}
}

// SetParamAlways adds a string parameter even if empty.
func (r *Request) SetParamAlways(name, value string) {
	r.Params = append(r.Params, Param{Name: name, Value: value})
}

// SetComplexParam adds a parameter with children.
func (r *Request) SetComplexParam(name string, children []Param) {
	r.Params = append(r.Params, Param{Name: name, Children: children})
}

// SetBodyFile reads a JSON file and converts it to nested SOAP params under the given wrapper element.
func (r *Request) SetBodyFile(wrapperName, filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading body file %s: %w", filePath, err)
	}
	return r.SetBodyRaw(wrapperName, data)
}

// SetBodyRaw converts JSON bytes to nested SOAP params under the given wrapper element.
func (r *Request) SetBodyRaw(wrapperName string, data []byte) error {
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("parsing JSON body: %w", err)
	}

	children := jsonToParams(obj)
	r.Params = append(r.Params, Param{Name: wrapperName, Children: children})
	return nil
}

// jsonToParams recursively converts a JSON object to a slice of Params.
func jsonToParams(obj map[string]any) []Param {
	keys := make([]string, 0, len(obj))
	for k := range obj {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var params []Param
	for _, k := range keys {
		v := obj[k]
		switch val := v.(type) {
		case map[string]any:
			params = append(params, Param{Name: k, Children: jsonToParams(val)})
		case []any:
			for _, elem := range val {
				switch e := elem.(type) {
				case map[string]any:
					params = append(params, Param{Name: k, Children: jsonToParams(e)})
				default:
					params = append(params, Param{Name: k, Value: fmt.Sprintf("%v", e)})
				}
			}
		default:
			params = append(params, Param{Name: k, Value: fmt.Sprintf("%v", val)})
		}
	}
	return params
}
