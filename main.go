package main

import (
	"strings"
	"ascii-art/functions"
	flags "ascii-art/parser"
)



func main() {
	cf := flags.Parse()
	data := functions.ReadFile("./banners/" + cf.Banner + ".txt")
	test := strings.Join(data, "\n")
	databyte := []byte(test)

	functions.TraitmentData(databyte, cf.StringArg, cf.OutputFile)
}
