package errors

import "fmt"

// NewControllersError returns a new specific Error for Controllers errors.
// It return only one InnerError in the `Errors` silce.
func NewControllersError(k string, metadata M) error {
	key := fmt.Sprintf("controllers-%s", k)
	return &Error{
		Status:     status(key),
		StatusText: statusText(key),
		Errors: []InnerError{{
			Code:     code(key),
			Kind:     "controllers",
			Metadata: appendReasonTo(key, metadata),
		}},
	}
}
