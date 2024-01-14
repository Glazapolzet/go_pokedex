package commands

import (
	"fmt"
	"os"
)

func exit() error {
	fmt.Printf("\nBye-bye!\n\n")

	os.Exit(0)

	return nil
}
