package main

import "os"

func main() {

	_, err := os.Open("")
	// START OMIT
	if errPath, ok := err.(*os.PathError); ok {
		// not PathError
	}
	// OMIT END
}

// end main OMIT
