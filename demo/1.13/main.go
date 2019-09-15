package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	_, err := FetchCat()
	if err != nil {
		if errors.Is(err, bytes.ErrTooLarge) {
			fmt.Println("File too large")
			return
		}

		var errPath *os.PathError
		if !errors.As(err, &errPath) {
			panic(err)
		}

		fmt.Println("Error on file:", errPath.Path)
	}
}

func FetchCat() (*Cat, error) {
	b, err := ioutil.ReadFile("cat.db")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}

	return NewCat(b), nil
}
