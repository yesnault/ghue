package internal

import (
	"fmt"
	"os"

	"github.com/yesnault/ghue/sdk/common"
)

// Exit func display an error message on stderr and exit 1
func Exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

// Check checks e and panic if not nil
func Check(err error) {
	if err != nil {
		if Verbose {
			panic(err)
		}
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}

// CheckErrors checks err, panic if not nil, then checks errHUE
// returns Exit Code 1 if error
func CheckErrors(err error, errHUE *common.ErrorHUE) {
	Check(err)
	if errHUE != nil {
		fmt.Fprintf(os.Stderr, "HUE Error: %s\n", errHUE.Error.Description)
		os.Exit(1)
	}
}
