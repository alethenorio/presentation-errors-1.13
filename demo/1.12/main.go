package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if _, err := FetchCat(); err != nil {
		catError, ok := err.(*CatError)
		if !ok {
			panic(err)
		}

		err := catError.Err
		if err == bytes.ErrTooLarge {
			fmt.Println("Too large")
			return
		}

		pathError, ok := err.(*os.PathError)
		if !ok {
			panic(err)
		}

		fmt.Println("Operation:", pathError.Op)

	}
}

type CatError struct {
	Reason string
	Err    error
}

func (e *CatError) Error() string {
	return e.Reason + ":" + e.Err.Error()
}

func FetchCat() (*Cat, error) {
	b, err := ioutil.ReadFile("cat.db")
	if err != nil {
		return nil, &CatError{Reason: "cannot open cat database", Err: err}
	}

	return NewCat(b), nil
}
