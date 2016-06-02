package stdout

import (
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

func isTerminal() bool {
	return terminal.IsTerminal(int(os.Stdout.Fd()))
}
