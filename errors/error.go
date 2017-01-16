package errors

import (
	"bytes"
	"fmt"
	"net/http"
	"text/template"

	"github.com/mdouchement/lss/utils"
)

// HTTPCoder interface is implemented by application errors
type HTTPCoder interface {
	// HTTPCode return the HTTP status code for the given error
	HTTPCode() int
}

// M is metadata structure
type M map[string]interface{}

// Error describe all errors occurred when a license is asked.
type Error struct {
	Status     int          `json:"status"`
	StatusText string       `json:"status_text"`
	Errors     []InnerError `json:"errors"`
}

// InnerErrors holds all errors raised during a request.
type InnerErrors []InnerError

// InnerError details an error occurred during a request.
type InnerError struct {
	Code     string `json:"code"`
	Kind     string `json:"type"`
	Metadata M      `json:"metadata,omitempty"`
}

var codeList = map[string]map[string]string{
	"controllers-copy": {
		"code":        "500-000",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"engines-open": {
		"code":        "500-001",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"engines-create": {
		"code":        "500-002",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"engines-delete": {
		"code":        "500-003",
		"status":      "500",
		"status_text": "Internal Server Error",
	},
	"controllers-invalid_path": {
		"code":        "422-000",
		"status":      "422",
		"status_text": "Unprocessable Entity",
		"reason":      "Path `{{.path}}` is not valid.",
	},
}

// StatusCode returns the HTTP status code from the given err.
func StatusCode(err error) int {
	if hc, ok := err.(HTTPCoder); ok {
		return hc.HTTPCode()
	}
	return http.StatusInternalServerError
}

// Error returns a well formated string of the current error.
func (e *InnerError) Error() string {
	return fmt.Sprintf("%s-%s: %s", e.Kind, e.Code, e.Metadata["reason"])
}

// Error concats all InnerError in a single string.
func (e *Error) Error() string {
	var errf bytes.Buffer
	errf.WriteString("[")
	for i, err := range e.Errors {
		errf.WriteString("\"")
		errf.WriteString(err.Error())
		if i < len(e.Errors)-1 {
			errf.WriteString("\",")
		} else {
			errf.WriteString("\"]")
		}
	}
	return errf.String()
}

// HTTPCode returns the HHTP status code of the current error.
func (e *Error) HTTPCode() int {
	return e.Status
}

func code(key string) string {
	return codeList[key]["code"]
}

func statusText(key string) string {
	return codeList[key]["status_text"]
}

func status(key string) int {
	return utils.MustAtoi(codeList[key]["status"])
}

func appendReasonTo(key string, metadata M) M {
	if reasonTemplate, ok := codeList[key]["reason"]; ok {
		t := template.Must(template.New("reason").Parse(reasonTemplate))
		reason := &bytes.Buffer{}
		t.Execute(reason, metadata)
		metadata["reason"] = reason.String()
	}
	return metadata
}
