package main

import (
	"ascii-art/functions"
	flags "ascii-art/parser"
	"strings"
)

func main() {
	cf := flags.Parse()
	data := functions.ReadFile("./banners/" + cf.Banner + ".txt")
	test := strings.Join(data, "\n")
	databyte := []byte(test)
	width := functions.GetTerminalWidth()

	functions.TraitmentData(databyte, cf.StringArg, cf.OutputFile, cf.Align, width)
}
