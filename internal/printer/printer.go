package printer

import (
	"fmt"
	"os"
)

// Color Coder's
const (
	RED        string = "\u001b[31;1m"
	BLUE       string = "\u001b[34;1m"
	GREEN      string = "\u001b[32;1m"
	BLACK      string = "\u001b[30;1m"
	WHITE      string = "\u001b[37;1m"
	YELLOW     string = "\u001b[33;1m"
	MAGENTA    string = "\u001b[35;1m"
	CYAN       string = "\u001b[36;1m"
	RESET      string = "\u001b[0m"
	BOLD       string = "\u001b[1m"
	UUNDERLINE string = "\u001b[4m"
	REVERSED   string = "\u001b[7m"
)

var (
	stdin     = *os.Stdin
	stdout    = *os.Stdout
	stderr    = *os.Stderr
	Required  = RED + "(Required)" + RESET
	Warningx  = YELLOW + "(Warning)" + RESET
)

func Println(a ...interface{}) {
	fmt.Fprintln(&stdout, a...)
}

func Printf(format string, a ...interface{}) {
	fmt.Fprintf(&stdout, format, a...)
}