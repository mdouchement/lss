package errors

import "fmt"

// NewEnginesError returns a new specific Error for Engines errors.
// It return only one InnerError in the `Errors` silce.
func NewEnginesError(k string, metadata M) error {
	key := fmt.Sprintf("engines-%s", k)
	return &Error{
		Status:     status(key),
		StatusText: statusText(key),
		Errors: []InnerError{{
			Code:     code(key),
			Kind:     "engines",
			Metadata: appendReasonTo(key, metadata),
		}},
	}
}
