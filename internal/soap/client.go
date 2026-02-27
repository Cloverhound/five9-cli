package soap

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Cloverhound/five9-cli/internal/config"
)

// ErrDryRun is returned when a write operation is intercepted by --dry-run mode.
var ErrDryRun = errors.New("dry run: no changes made")

// Do executes a SOAP request and returns the parsed JSON response.
func Do(req *Request) ([]byte, error) {
	envelope, err := BuildEnvelope(req.Method, req.Params)
	if err != nil {
		return nil, fmt.Errorf("building envelope: %w", err)
	}

	endpoint := config.Endpoint()

	if config.Debug() {
		fmt.Fprintf(os.Stderr, "DEBUG: POST %s\n", endpoint)
		fmt.Fprintf(os.Stderr, "DEBUG: SOAPAction: %s\n", req.Method)
		fmt.Fprintf(os.Stderr, "DEBUG: Body:\n%s\n", truncate(envelope, 2000))
	}

	if config.DryRun() && isWriteMethod(req.Method) {
		fmt.Fprintf(os.Stderr, "[DRY RUN] POST %s\n", endpoint)
		fmt.Fprintf(os.Stderr, "[DRY RUN] Method: %s\n", req.Method)
		fmt.Fprintf(os.Stderr, "[DRY RUN] Body:\n%s\n", envelope)
		return nil, ErrDryRun
	}

	httpReq, err := http.NewRequest("POST", endpoint, strings.NewReader(envelope))
	if err != nil {
		return nil, fmt.Errorf("building HTTP request: %w", err)
	}

	// Basic auth
	auth := base64.StdEncoding.EncodeToString([]byte(config.Username() + ":" + config.Password()))
	httpReq.Header.Set("Authorization", "Basic "+auth)
	httpReq.Header.Set("Content-Type", "text/xml; charset=utf-8")
	httpReq.Header.Set("SOAPAction", "")

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if config.Debug() {
		fmt.Fprintf(os.Stderr, "DEBUG: Response %d (%d bytes)\n", resp.StatusCode, len(body))
		fmt.Fprintf(os.Stderr, "DEBUG: Body:\n%s\n", truncate(string(body), 2000))
	}

	// Check for SOAP fault
	if fault := ExtractFault(body); fault != nil {
		return nil, fault
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, truncate(string(body), 500))
	}

	// Parse SOAP response to JSON
	return ParseResponse(body, req.Method)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// isWriteMethod returns true for SOAP methods that modify data.
func isWriteMethod(method string) bool {
	prefixes := []string{
		"create", "modify", "delete", "add", "remove", "set",
		"start", "stop", "reset", "force", "close", "rename",
		"update", "async",
	}
	lower := strings.ToLower(method)
	for _, p := range prefixes {
		if strings.HasPrefix(lower, p) {
			return true
		}
	}
	return false
}
