package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/integrii/flaggy"
	"github.com/wfrodriguez/console"
	"github.com/wfrodriguez/templet/internal/core"
)

var ext string
var name string
var isJson bool
var paramsCmd *flaggy.Subcommand

func configureCommands() {

	flaggy.String(&ext, "e", "ext", "Extension del template a procesar (Sin el punto)")
	flaggy.String(&name, "n", "name", "Nombre del template a procesar")

	paramsCmd = flaggy.NewSubcommand("params")
	paramsCmd.Bool(&isJson, "j", "json", "Devuelve el resultado en formato JSON, si no se indica devuelve el resultado"+
		" en formato texto plano separados por saltos de linea")

	flaggy.AttachSubcommand(paramsCmd, 1)
}

func emptyStrings(args ...string) bool {
	for _, arg := range args {
		if arg == "" {
			return true
		}
	}
	return false
}

func params(a *core.Application) {
	str, err := a.ExtractVariables(ext, name)
	if err != nil {
		console.Error(err.String())
		os.Exit(3)
	}
	if isJson {
		d, e := json.Marshal(str)
		if e != nil {
			console.Error(err.String())
			os.Exit(3)
		}
		fmt.Print(string(d))
	} else {
		p := strings.Join(str, "\n")
		fmt.Print(p)
	}
}

func main() {
	configureCommands()
	flaggy.Parse()

	if emptyStrings(ext, name) {
		console.Error("Las opciones `--ext` y `--name` son obligatorias")
		os.Exit(2)
	}

	app, err := core.NewApplication()
	if err != nil {
		console.Error(err.String())
	}
	if paramsCmd.Used {
		params(app)
	}
}
