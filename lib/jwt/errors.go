package jwt

import "fmt"

type ErrDIDIsNotSet struct {
	details string
}

func NewErrDIDIsNotSet(details string) *ErrDIDIsNotSet {
	return &ErrDIDIsNotSet{
		details: details,
	}
}

func (e *ErrDIDIsNotSet) Error() string {
	return fmt.Sprintf("did isn't set: %v", e.details)
}
