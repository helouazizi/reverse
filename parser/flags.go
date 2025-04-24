package flags

import (
	"ascii-art/functions"
	"flag"
	"fmt"
	"os"
)

type Config struct {
	OutputFile string
	Align      string
	StringArg  string
	Banner     string
	Color      string
	Reverse    string
}

func Parse() *Config {
	// Set defaults
	message := "Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard"

	//// hhhhhhhhhhhhhhhhhhh
	for i, arg := range os.Args {
		if arg == "--align" && i+1 < len(os.Args) {
			fmt.Println(message)
			os.Exit(0)
		}
		if arg == "--output" && i+1 < len(os.Args) {
			fmt.Println(message)
			os.Exit(0)
		}
		if arg == "--color" && i+1 < len(os.Args) {
			fmt.Println(message)
			os.Exit(0)
		}
		if arg == "--reverse" && i+1 < len(os.Args) {
			fmt.Println(message)
			os.Exit(0)
		}
	}
	var align string
	var outputFile string
	var color string
	var reverse string
	// Define flags
	flag.StringVar(&outputFile, "output", "", "Output file name <exemple.txt>")
	flag.StringVar(&align, "align", "left", "Text alignment: left, center,right or justify")
	flag.StringVar(&color, "color", "", "color text <red>")
	flag.StringVar(&reverse, "reverse", "", "reverse file <file.txt>")
	flag.Usage = func() {
		// fmt.Println(message)
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	// check args
	args := flag.Args()
	if len(args) < 1 || len(args) > 2 {
		fmt.Println(message)
		os.Exit(0)
	}

	inputString := args[0]
	banner := "standard"
	if len(args) == 2 {
		banner = args[1]
	}

	//check banner
	if banner != "" && banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		fmt.Println(message)
		os.Exit(0)
	}

	//check alignement options
	if align != "" && align != "left" && align != "right" && align != "center" && align != "justify" {
		fmt.Println(message)
		os.Exit(0)
	}

	//check output file

	if !functions.SafeFile(outputFile) {
		fmt.Println(message)
		os.Exit(0)
	}
	return &Config{
		OutputFile: outputFile,
		Align:      align,
		StringArg:  inputString,
		Banner:     banner,
		Color:      color,
		Reverse:    reverse,
	}
}
