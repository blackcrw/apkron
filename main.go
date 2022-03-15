/* Lista de Tarefas:

1. Criar um signal para SHIFT+C
2. Testar as novas funções

*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/blackcrw/apkron/cli"
)

func main() {
	go func(){
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, os.Interrupt)
	
		<-sc
	
		fmt.Printf("Your press CTRL+C\r\n")
		syscall.Exit(0)	
	} ()

	cli.OnInitialize()
}