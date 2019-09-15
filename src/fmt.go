package main

import "fmt"

func main() {
	var err error
	// START OMIT
	// returns a new error which implements Unwrap()
	// to retrieve the underlying error
	fmt.Errorf("unable to open database: %w", err) // HL
	// END OMIT
}
