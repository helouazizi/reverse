package main

import (
	"ascii-art/functions"
	flags "ascii-art/parser"
	"bufio"
	"fmt"
	"os"
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
		reverse, _ := Read(cf.Reverse)
		tex := functions.ReverseAsciiArt(reverse, ascciMap)
		fmt.Println(tex)
		// for _, v := range reverse{
		// 	fmt.Println(v,len(v))
		// }
		return
	}
	functions.TraitmentData(databyte, cf.StringArg, cf.OutputFile, cf.Align, cf.Color, width)
}

func Read(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// banner := make(map[string]rune)

	var lines []string
	// /	charRune := rune(32) // ASCII starts at 32 for printable characters

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines between blocks
		// if len(lines) == 0 && line == "" {
		// 	continue
		// }

		lines = append(lines, line)

		// When a full character block is read
		// if len(lines) == asciiHeight {
		// 	// Normalize: make sure all lines are the same length
		// 	// maxLen := 0
		// 	// for _, l := range lines {
		// 	// 	if len(l) > maxLen {
		// 	// 		maxLen = len(l)
		// 	// 	}
		// 	// }
		// 	// for i := range lines {
		// 	// 	lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i]))
		// 	// }

		// 	banner[strings.Join(lines, "\n")] = charRune
		// 	charRune++
		// 	lines = []string{}
		// }
	}

	return lines, nil
}
