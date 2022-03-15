package printer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const (
	RED        = "\u001b[31;1m"
	BLUE       = "\u001b[34;1m"
	GREEN      = "\u001b[32;1m"
	BLACK      = "\u001b[30;1m"
	WHITE      = "\u001b[37;1m"
	YELLOW     = "\u001b[33;1m"
	MAGENTA    = "\u001b[35;1m"
	CYAN       = "\u001b[36;1m"
	RESET      = "\u001b[0m"
	BOLD       = "\u001b[1m"
	UUNDERLINE = "\u001b[4m"
	REVERSED   = "\u001b[7m"

	PREFIX_DONE    = GREEN   + "[+]"+RESET
	PREFIX_DANGER  = RED     + "[-]"+RESET
	PREFIX_INFO    = MAGENTA + "[i]"+RESET
	PREFIX_SCAN    = YELLOW  + "[?]"+RESET

	MSG_REQUIRED  = RED    + "(Required)" +RESET
	MSG_WARNING   = YELLOW + "(Warning)"  +RESET

	endl = "\n"
)

var (
	stdin     = *os.Stdin
	stdout    = *os.Stdout
	stderr    = *os.Stderr
)

func doPrintbs(a ...interface{}) (str string){
	for arg_num, arg := range a {
		if arg_num > 0 { str += " " }
		str += fmt.Sprint(arg)
	}

	return
}

func Println(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...))
}

func Printf(format string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf(format, a...))
}

func Fatalln(a ...interface{}) {
	stderr.WriteString(RED + "[ERROR]" + doPrintbs(a...) +endl)
	syscall.Exit(0)
}

// This function is to avoid the constant use of if's for error checking
func CError(err error) { if err != nil { Fatalln(err) } }

func Done(a ...interface{}) {
	stdout.WriteString(PREFIX_DONE +" "+ doPrintbs(a...) +endl)
}

func Danger(a ...interface{}) {
	stdout.WriteString(PREFIX_DANGER +" "+ doPrintbs(a...) +endl)
}

func Info(a ...interface{}) {
	stdout.WriteString(PREFIX_INFO +" "+ doPrintbs(a...) +endl)
}

func PrinterPrefixColor(prefix string, color string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf("%s%s %s%s\n", prefix, color, RESET, doPrintbs(a...)))
}

func Scan(a ...interface{}) string {
	stdout.WriteString(PREFIX_SCAN +" "+ doPrintbs(a...))

	var response, err = bufio.NewReader(&stdin).ReadString('\n')

	if err != nil {
		Fatalln(err)
	}

	response = strings.ToLower(response)

	if response == "\n" {
		return response
	}

	response = strings.Replace(response, "\n", "", -1)

	return response
}