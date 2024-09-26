package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-output/functions"
)

func main() {
	// check the args befor extracting our args
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		return
	}
	Args := os.Args
	resultFile := ""
	str := ""
	template := "standard"
	status := false
	

if len(Args) == 4 {
	resultFile = Args[1]
	str=Args[2]
	template = Args[3]
	status = true
}else if len(Args) == 3 {
	if functions.TrueOutput(Args[1]){
		resultFile = Args[1]
		str=Args[2]
		status = true

	} else {
		str = Args[1]
		template = Args[2]

	}

}else {
	str = Args[1]
}
	// extract the file result
	resultFile = functions.Extract_Result_File_Name(resultFile)

	// extract data from the  file as an array of caracters

	data := functions.ReadFile("./banners/" + template + ".txt")
	test := strings.Join(data, "\n")
	databyte := []byte(test)

	// send this args to trairment and print inside this function
	functions.TraitmentData(databyte, str, resultFile, status)
}
