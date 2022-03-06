package main

import (
	"os"
	"sync"
	"time"

	"github.com/blackcrw/apkron/internal/apktool"
	"github.com/blackcrw/apkron/internal/printer"
	temp_path "github.com/blackcrw/apkron/internal/temp"
	"github.com/blackcrw/apkron/tools/interesting"
)

var wg sync.WaitGroup

func check_error(err error) {
	if err != nil {
		printer.Printf("%s[X]%s %s\n", printer.RED, printer.RESET, err)
		
		defer os.Exit(0)
	}
}

func main() {
	check_error(apktool.Exists())
	printer.Printf("%s[+]%s %s\n", printer.GREEN, printer.RESET, "Apk Tool Exists\n")
	
	var temporary_directory, err = temp_path.CreateTemporaryDirectory()
	defer temp_path.DeleteTemporaryDirectory(temporary_directory)

	check_error(err)
	printer.Printf("%s[+]%s %s: %s\n", printer.GREEN, printer.RESET, "Temporary Directory Created", temporary_directory)

	wg.Add(1)

	go func() {
		var _, err = apktool.Decompile("decompile_test.apk", temporary_directory)
		check_error(err)
		
		defer wg.Done()
	}()

	printer.Printf("%s[+]%s %s\n", printer.GREEN, printer.RESET, "Decompiling application!")
	
	wg.Wait()
	
	interesting.AndroidManifest(temporary_directory)

	time.Sleep(5*time.Second)
}