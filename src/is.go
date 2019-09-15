package main

import (
	"errors"
	"io"
)

func main() {
	var err error
	// START BEFORE OMIT
	if err == io.ErrUnexpectedEOF {
		// do something
	}
	// END BEFORE OMIT

	// START AFTER OMIT
	if errors.Is(err, io.ErrUnexpectedEOF) { // HL
		// do something
	}
	// END AFTER OMIT
}
