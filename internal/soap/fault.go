package soap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Fault represents a SOAP fault.
type Fault struct {
	Code   string
	String string
	Detail string
}

func (f *Fault) Error() string {
	if f.Detail != "" {
		return fmt.Sprintf("SOAP Fault: %s — %s", f.String, f.Detail)
	}
	return fmt.Sprintf("SOAP Fault: %s", f.String)
}

// ExtractFault checks if the SOAP response contains a fault and returns it.
func ExtractFault(data []byte) *Fault {
	type soapFault struct {
		XMLName     xml.Name `xml:"Envelope"`
		FaultCode   string   `xml:"Body>Fault>faultcode"`
		FaultString string   `xml:"Body>Fault>faultstring"`
		Detail      string   `xml:"Body>Fault>detail"`
	}

	var env soapFault
	if err := xml.Unmarshal(data, &env); err != nil {
		return nil
	}

	if env.FaultString == "" && env.FaultCode == "" {
		return nil
	}

	return &Fault{
		Code:   env.FaultCode,
		String: strings.TrimSpace(env.FaultString),
		Detail: strings.TrimSpace(env.Detail),
	}
}
