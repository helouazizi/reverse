package functions

import (
	"bufio"
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
func TraitmentData(text []byte, arg, resultFile, align, color string, width int) {
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
		PrintColored(result, color)
	}
}

// traitment the data if it have charachters
func Final_result(arrData, words []string, align string, width int) string {
	result := ""
	textWidth := 0
	// fmt.Println(textWidth,"textwidth")

	for k := range words {
		if words[k] == "" {
			fmt.Println("some here")
			result += "\n"
			continue
		}
		// this specialy for justyfy option
		wordsLength := len(strings.Split(words[k], " "))
		// fmt.Println(wordsLength, "wordlength")

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
				// fmt.Println(Ascii,string(words[k][j]),"debud")
				start := Ascii*8 + Ascii + 1 + i
				result += arrData[start]
				if Ascii == 0 && align == "justify" {
					// in this part we need to detect if the length odd or even
					var spaceToAdd int
					if wordsLength > 1 {
						spaceToAdd = (width - textWidth) / wordsLength
						spaceToAdd += spaceToAdd / (wordsLength - 1)
					}
					result += fmt.Sprintf("%*s", spaceToAdd, "")
					// continue
				}
			}
			result += "\n"
		}
	}
	// result = strings.TrimSuffix(result," ")
	return result
}

func PrintColored(text string, color string) {
	colorCodes := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
	}
	reset := "\033[0m"
	code, exists := colorCodes[color]
	if !exists {
		fmt.Printf("%s", text) // fallback to normal
		return
	}

	fmt.Printf("%s%s%s\n", code, text, reset)
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

// /////////////////////////////////////////////////////////////////////////////////////////
const asciiHeight = 8

func LoadBanner(path string) (map[string]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	banner := make(map[string]rune)

	var lines []string
	charRune := rune(32) // ASCII starts at 32 for printable characters

	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines between blocks
		if len(lines) == 0 && line == "" {
			continue
		}

		lines = append(lines, line)

		// When a full character block is read
		if len(lines) == asciiHeight {
			// Normalize: make sure all lines are the same length
			// maxLen := 0
			// for _, l := range lines {
			// 	if len(l) > maxLen {
			// 		maxLen = len(l)
			// 	}
			// }
			// for i := range lines {
			// 	lines[i] = lines[i] + strings.Repeat(" ", maxLen-len(lines[i]))
			// }

			banner[strings.Join(lines, "\n")] = charRune
			charRune++
			lines = []string{}
		}
	}

	return banner, nil
}

func ReverseAsciiArt(asciiLines []string, charMap map[string]rune) string {
	if len(asciiLines) != asciiHeight {
		return "[ERROR: asciiLines must be 8 lines tall]"
	}
	result := ""
	pos := 0
	lineLength := len(asciiLines[0])

	for pos < lineLength {
		matched := false
		// Try to match the longest possible character block from this position
		for w := 1; w <= lineLength-pos; w++ {
			chunk := make([]string, asciiHeight)
			for i := 0; i < asciiHeight; i++ {
				chunk[i] = asciiLines[i][pos : pos+w]
			}
			key := strings.Join(chunk, "\n")

			if ch, ok := charMap[key]; ok {
				result += string(ch)
				pos += w
				matched = true
				break
			}
		}

		if !matched {
			// No match found — skip 1 column and append ?
			result += "?"
			pos++
		}
	}

	return result
}

