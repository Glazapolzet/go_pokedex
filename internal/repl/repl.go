package repl

import (
	"bufio"
	"os"
)

func GetUserPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	return scanner.Text()
}
