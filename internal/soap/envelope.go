package soap

import (
	"encoding/xml"
	"strings"
)

const soapNS = "http://service.admin.ws.five9.com/"

// Param represents a SOAP parameter, which can be a leaf (Value set) or a branch (Children set).
type Param struct {
	Name     string
	Value    string
	Children []Param
}

// BuildEnvelope constructs a SOAP XML envelope for the given method and params.
func BuildEnvelope(method string, params []Param) (string, error) {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	b.WriteString(`<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ser="`)
	b.WriteString(soapNS)
	b.WriteString(`">`)
	b.WriteString(`<soapenv:Header/>`)
	b.WriteString(`<soapenv:Body>`)
	b.WriteString(`<ser:`)
	b.WriteString(method)
	b.WriteString(`>`)

	for _, p := range params {
		writeParam(&b, p)
	}

	b.WriteString(`</ser:`)
	b.WriteString(method)
	b.WriteString(`>`)
	b.WriteString(`</soapenv:Body>`)
	b.WriteString(`</soapenv:Envelope>`)

	return b.String(), nil
}

func writeParam(b *strings.Builder, p Param) {
	b.WriteString(`<`)
	b.WriteString(p.Name)
	b.WriteString(`>`)
	if len(p.Children) > 0 {
		for _, child := range p.Children {
			writeParam(b, child)
		}
	} else {
		xml.EscapeText(b, []byte(p.Value))
	}
	b.WriteString(`</`)
	b.WriteString(p.Name)
	b.WriteString(`>`)
}
