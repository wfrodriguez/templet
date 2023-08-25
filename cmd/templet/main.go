package main

import (
	"fmt"

	"github.com/integrii/flaggy"
)

var ext string
var name string
var paramsCmd *flaggy.Subcommand

func configureCommands() {

	flaggy.String(&ext, "e", "ext", "Extension del template a procesar (Sin el punto)")
	flaggy.String(&name, "n", "name", "Nombre del template a procesar")

	paramsCmd = flaggy.NewSubcommand("params")

	flaggy.AttachSubcommand(paramsCmd, 1)
}

func main() {
	// fmt.Println(template.Logo)
	configureCommands()
	flaggy.Parse()

	if paramsCmd.Used {
		fmt.Println("comando params usado")
	}
}
