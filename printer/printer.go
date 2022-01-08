// printer/printer.go

package printer

import (
	"fmt"
	"os"
	"os/exec"
)

type calculator struct {
	display string
	lastVal int
	currOp  string
}

func ReceiveAndPrint() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	calc := calculator{display: "", lastVal: 0, currOp: "None"}
	for {

		curr := make([]byte, 1)
		os.Stdin.Read(curr)
		if curr[0] == 'c' || curr[0] == 'C' {
			calc.display = ""
			continue
		}

		calc.display += string(curr)

		fmt.Println(calc.display)
	}
}

func isDigit(curr []byte) bool {
	return '0' <= curr[0] && curr[0] <= 9
}

func isOperator(curr []byte) bool {
	return curr[0] == '=' ||
		curr[0] == '+' ||
		curr[0] == '-' ||
		curr[0] == '*' ||
		curr[0] == '/'
}
