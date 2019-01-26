package utils

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", errors.WithMessage(err, "error"))
	os.Exit(1)
}
