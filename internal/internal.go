package internal

import (
	"github.com/blackcrw/apkron/internal/apktool"
	"github.com/blackcrw/apkron/internal/printer"
)

const (
	PROJECT_NAME  = "APKRON"
	PROJECT_VERSION = "0.0.0.0"
	PROJECT_BANNER = `
	——————————————————————
	▒▒▒▒▒▒▒▒▄▄▄▄▄▄▄▄▒▒▒▒▒▒
	▒▒█▒▒▒▄██████████▄▒▒▒▒
	▒█▐▒▒▒████████████▒▒▒▒
	▒▌▐▒▒██▄▀██████▀▄██▒▒▒
	▐┼▐▒▒██▄▄▄▄██▄▄▄▄██▒▒▒
	▐┼▐▒▒██████████████▒▒▒
	▐▄▐████─▀▐▐▀█─█─▌▐██▄▒
	▒▒█████──────────▐███▌
	▒▒█▀▀██▄█─▄───▐─▄███▀▒
	▒▒█▒▒███████▄██████▒▒▒
	▒▒▒▒▒██████████████▒▒▒
	▒▒▒▒▒█████████▐▌██▌▒▒▒
	▒▒▒▒▒▐▀▐▒▌▀█▀▒▐▒█▒▒▒▒▒
	▒▒▒▒▒▒▒▒▒▒▒▐▒▒▒▒▌▒▒▒▒▒
	——————————————————————
	    	APKRON
`
)

func InitApkron() {
	printer.Println(PROJECT_BANNER)
	printer.Printf("%s", "——————————————————————————————————————\n")
	printer.Info("APKRON Version:  0.0.0.0")
	
	{	printer.CError(apktool.Exists())
		var version, err = apktool.Version(); printer.CError(err)
		printer.Info("APKTOOL Version:", version)		}
}
