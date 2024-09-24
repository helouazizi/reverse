package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-output/functions"
)

func main() {
	// check the args befor extracting our args
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <string>")
		return
	}
	Args := os.Args
	resultFile := Args[1]
	str := Args[2]
	template := ""

	// check if there is no banner set the banner in standard
	if len(Args) > 3 {
		template = os.Args[3]
	} else {
		template = "standard"
	}

	// extract data from the  file as an array of caracters

	data := functions.ReadFile("./banners/" + template + ".txt")
	test := strings.Join(data, "\n")
	databyte := []byte(test)

	// extract the file result
	resultFile = functions.Extract_File_Name(resultFile)

	// send this args to trairment and print inside this function
	functions.TraitmentData(databyte, str, resultFile)

	////////////// so the project is done ////////////////////////
}
