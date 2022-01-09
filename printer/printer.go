// printer/printer.go

package printer

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type calculator struct {
	display string
	lastVal int
	currOp  string
}

func newCalc() calculator {
	return calculator{display: "", lastVal: noVal, currOp: "None"}
}

func errCalc() calculator {
	return calculator{display: "ERR", lastVal: noVal, currOp: "None"}
}

const noVal = 999999999

func ReceiveAndPrint() {
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	calc := newCalc()
	for {

		curr := make([]byte, 1)
		os.Stdin.Read(curr)

		if curr[0] == 'E' {
			break
		}

		if calc.display == "ERR" {
			if curr[0] == 'A' {
				calc = newCalc()
			} else {
				continue
			}
		}

		if curr[0] == 'C' {
			calc.display = ""
		} else if curr[0] == 'A' {
			calc = newCalc()
		} else if isDigit(curr) {
			if len(calc.display) < 8 {
				calc.display += string(curr)
			}
		} else if isOperator(curr) {
			if curr[0] == '=' {
				if calc.display != "" &&
					calc.currOp != "None" &&
					calc.lastVal != noVal {
					newVal := performOp(&calc)
					calc.display = strconv.Itoa(newVal)
					calc.currOp = "None"
					calc.lastVal = noVal
					if newVal > 99999999 {
						calc = errCalc()
					}
				} else if calc.display != "" &&
					calc.currOp == "None" &&
					calc.lastVal == noVal {
					numDisplay, _ := strconv.ParseInt(calc.display, 10, 64)
					calc.lastVal = int(numDisplay)
					calc.display = ""
				} else {
					calc = errCalc()
				}
			} else {
				if calc.display != "" &&
					calc.currOp == "None" {
					numDisplay, _ := strconv.ParseInt(calc.display, 10, 64)
					calc.lastVal = int(numDisplay)
					calc.currOp = string(curr)
					calc.display = string(curr)
					printCalc(calc)
					calc.display = ""
					continue
				} else if calc.display != "" &&
					calc.currOp == "None" &&
					calc.lastVal != noVal {
					calc.currOp = string(curr)
					newVal := performOp(&calc)
					calc.display = strconv.Itoa(newVal)
					calc.currOp = "None"
					calc.lastVal = noVal
					if newVal > 99999999 {
						calc = errCalc()
					} else {
						printCalc(calc)
						calc.display = ""
						calc.lastVal = newVal
						continue
					}
				} else {
					calc = errCalc()
				}
			}
		} else {
			calc = errCalc()
		}

		printCalc(calc)
	}

	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}

func performOp(calc *calculator) int {
	numDisplay, _ := strconv.ParseInt(calc.display, 10, 64)
	var newVal int
	switch calc.currOp {
	case "+":
		newVal = int(numDisplay) + calc.lastVal
	case "-":
		newVal = calc.lastVal - int(numDisplay)
	case "*":
		newVal = int(numDisplay) * calc.lastVal
	case "/":
		newVal = calc.lastVal / int(numDisplay)
	}
	return newVal
}

func printCalc(calc calculator) {
	if calc.lastVal != noVal {
		fmt.Printf("Display: %s  Operator: %s  lastVal: %d\n", calc.display, calc.currOp, calc.lastVal)
	} else {
		fmt.Printf("Display: %s  Operator: %s  lastVal: None\n", calc.display, calc.currOp)
	}
}

func isDigit(curr []byte) bool {
	return '0' <= curr[0] && curr[0] <= '9'
}

func isOperator(curr []byte) bool {
	return curr[0] == '=' ||
		curr[0] == '+' ||
		curr[0] == '-' ||
		curr[0] == '*' ||
		curr[0] == '/'
}
