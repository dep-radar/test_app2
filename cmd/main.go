package main

import (
	stderr "errors"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := stderr.New("my")
	err = errors.Wrap(err, "my wrap")
	fmt.Printf("Error: %s", err.Error())
}
