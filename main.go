/* Lista de Tarefas:

1. Criar um signal para SHIFT+C
2. Testar as novas funções

*/

package main

import (
	"sync"
	"time"

	"github.com/blackcrw/apkron/internal"
	"github.com/blackcrw/apkron/internal/apktool"
	"github.com/blackcrw/apkron/internal/printer"
	temp_path "github.com/blackcrw/apkron/internal/temp"
	"github.com/blackcrw/apkron/tools/interesting"
)

var wg sync.WaitGroup

func init() {
	internal.InitApkron()
}

func main() {	
	var temporary_directory, err = temp_path.CreateTempDirectory()
	defer temp_path.DeleteTempDirectory(temporary_directory)

	{	printer.CError(err)
		printer.Info("Temporary Directory Created", temporary_directory)	}

	wg.Add(1)

	go func() {
		{	var _, err = apktool.Decompile("decompile_test.apk", temporary_directory)
			printer.CError(err)	}
		
		defer wg.Done()
	}()

	printer.Done("Decompiling application!\n")
	
	wg.Wait()
	
	printer.Info("Application Permissions:\n")
	
	if perms, err := interesting.AndroidManifest(temporary_directory); err != nil {
		printer.CError(err)
	} else if len(perms.Permissions) > 0 {
		for _, x := range perms.Permissions {
			printer.Done("Name:", x.Name)
			printer.Printf("%s   -%s Description: %s\n", printer.MAGENTA, printer.RESET, x.Description)
			printer.Printf("%s   -%s Match: %s\n\n", printer.MAGENTA, printer.RESET, x.Match)
		}
	} else {
		printer.Danger("No permissions found")
	}

	time.Sleep(5*time.Second)
}