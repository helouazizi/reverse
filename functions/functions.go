package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")
	return result
}

// this is the the traitment functions
func TraitmentData(text []byte, arg, resultFile, align string, width int) {
	// cheeck if the char is in range or not if
	for _, char := range arg {
		if char < 32 || char > 126 {
			fmt.Println("Ereur : one of this carcters is not in range")
			return
		}
	}
	result := ""

	arrData := strings.Split(string(text), "\n")

	words := strings.Split(arg, `\n`)

	/////////////////////////////////
	// this part just for newlines
	count := 0
	for _, test := range words {
		if test == "" {
			count++
		}
	}
	// in case the data is all new line
	if count == (len(arg)/2)+1 {
		for range len(arg) / 2 {
			result += "\n"
		}
	} else {
		if resultFile != "" {
			result = Final_result(arrData, words, "left", width)
		} else {
			result = Final_result(arrData, words, align, width)
		}
	}
	if resultFile != "" {
		os.WriteFile(resultFile, []byte(result), 0o777)

	} else {
		fmt.Printf("%s", result)
	}
}

// traitment the data if it have charachters
func Final_result(arrData, words []string, align string, width int) string {
	result := ""
	textWidth := 0

	for k := range words {
		if words[k] == "" {
			result += "\n"
			continue
		}
		// this specialy for justyfy option
		wordsLength := len(strings.Fields(words[k]))
		fmt.Println(wordsLength, "tets")

		// lets claculate the text with to substract it from the terminal width
		if k == 0 {
			for j := 0; j < len(words[k]); j++ {
				Ascii := (int(words[k][j] - 32))
				start := Ascii*8 + Ascii + 1
				textWidth += len(arrData[start])
			}
		}

		for i := range 8 {
			// lets play some game with spaces depend on the alignment flag
			if align == "right" {
				spaceToAdd := width - textWidth
				result += fmt.Sprintf("%*s", spaceToAdd, "")
			}
			if align == "center" {
				spaceToAdd := width/2 - textWidth/2
				result += fmt.Sprintf("%*s", spaceToAdd, "")
			}
			for j := 0; j < len(words[k]); j++ {
				Ascii := (int(words[k][j] - 32))
				if Ascii == 0 && align == "justify" {
					var spaceToAdd int
					if wordsLength > 1 {
						spaceToAdd = (width - textWidth) / wordsLength
						spaceToAdd += spaceToAdd / (wordsLength - 1)
					}
					result += fmt.Sprintf("%*s", spaceToAdd, "")
				}
				start := Ascii*8 + Ascii + 1 + i
				result += arrData[start]
			}
			result += "\n"
		}
	}

	return result
}

func SafeFile(path string) bool {

	// Define restricted directories
	restrictedDirs := []string{
		"./banners/",
		"./functions/",
		"./parser/",
	}

	// Check if the absolute path is within any restricted directories
	for _, dir := range restrictedDirs {
		if strings.HasPrefix(path, dir) {
			return false // Path is inside a restricted directory
		}
	}
	if path == "main.go" || path == "go.mod" {
		return false
	}
	return true
}

func GetTerminalWidth() int {
	// First, try to fetch the terminal width from the $COLUMNS environment variable
	widthStr := os.Getenv("COLUMNS")
	if widthStr != "" {
		// If the environment variable exists, try to convert it to an integer
		width, err := strconv.Atoi(widthStr)
		if err == nil {
			return width
		}
	}

	// If $COLUMNS is not set or invalid, try `tput cols`
	cmd := exec.Command("tput", "cols")
	output, err := cmd.Output()
	if err == nil {
		// If `tput` works, convert the output to an integer
		width, err := strconv.Atoi(strings.TrimSpace(string(output)))
		if err == nil {
			return width
		}
	}
	return 80
}
