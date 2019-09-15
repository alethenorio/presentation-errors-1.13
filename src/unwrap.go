package main

import (
	"errors"
	"os"
)

// START ERRORS UNWRAP OMIT
// Unwrap returns the result of calling the Unwrap method on err,
// if err's type contains an Unwrap method returning error.
// Otherwise, Unwrap returns nil.
// END ERRORS UNWRAP OMIT

type CatError struct {
	Reason string
	Err    error
}

func (e *CatError) Error() string {
	return e.Reason + ": " + e.Err.Error()
}

func (e *CatError) Unwrap() error { // HL
	return e.Err // HL
} // HL
// end unwrap OMIT

func main() { // OMIT
	errCat := &CatError{
		Err:    os.ErrNotExist,
		Reason: "Cat is undisposed",
	}

	// os.ErrNotExist
	err := errors.Unwrap(errCat) // HL
	_ = err                      // end main OMIT
}
