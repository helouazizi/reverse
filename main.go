package main

import (
	"ascii-art/functions"
	flags "ascii-art/parser"
	"fmt"
	"strings"
)

func main() {
	cf := flags.Parse()
	data := functions.ReadFile("./banners/" + cf.Banner + ".txt")
	test := strings.Join(data, "\n")
	databyte := []byte(test)
	width := functions.GetTerminalWidth()
	if cf.Reverse != "" {
		ascciMap, _ := functions.LoadBanner("./banners/standard.txt")
		reverse := functions.ReadFile(cf.Reverse)
		tex := functions.ReverseAsciiArt(reverse,ascciMap)
		fmt.Println(tex)
		return
	}

	functions.TraitmentData(databyte, cf.StringArg, cf.OutputFile, cf.Align, cf.Color, width)
}
