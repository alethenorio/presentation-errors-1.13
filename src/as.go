package main

import (
	"errors"
	"os"
)

func main() {
	var err error
	// START BEFORE OMIT
	if pathError, ok := err.(*os.PathError); ok {
		// do something
	}
	// END BEFORE OMIT
	// START AFTER OMIT
	var pathError *os.PathError
	if errors.As(err, &pathError) { // HL
		// handle pathError
	}
	// END AFTER OMIT
}
