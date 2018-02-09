package main

import (
	stderr "errors"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	err := stderr.New("my")
	err = errors.Wrap(err, "my wrap")
	fmt.Printf("Error: %s", err.Error())
}
