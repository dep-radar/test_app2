package main

import (
	stderr "errors"
	"fmt"

	"github.com/pkg/errors"
	"github.com/pkg/profile"
	"github.com/pkg/sftp"
)

func main() {
	defer profile.Start().Stop()
	err := stderr.New("my")
	err = errors.Wrap(err, "my wrap")
	fmt.Printf("Error: %s %d", err.Error(), sftp.SftpServerWorkerCount)
}
