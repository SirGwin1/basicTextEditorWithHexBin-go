package main

import (
	"BASICTEXTEDITORWITHHEXBIN-GO/functionality"
	"fmt"
	"io"
	"os"
)

func main() {

	commandlineArg := os.Args[1:]

	if len(commandlineArg) != 2 {
		fmt.Println("you are required to enter an input and output file to run the program")
		os.Exit(1)
	}

	inputfile := commandlineArg[0]
	outputfile := commandlineArg[1]

	file1, err := os.Open(inputfile)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file1.Close()

	readContent, err := io.ReadAll(file1)
	if err != nil {
		fmt.Println(err.Error())
	}

	acceptedContent := string(readContent)

	// main transformation
	step1 := functionality.TextNormalizer(acceptedContent)
	step2 := functionality.BaseConversion(step1)
	step3 := functionality.CaseModifier(step2)
	step4 := functionality.Punctuation(step3)
	step5 := functionality.Vowelcheck(step4)

	file2, err := os.OpenFile(outputfile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println(err.Error())
	}
	defer file2.Close()

	_, err = file2.WriteString(step5)
	if err != nil {
		fmt.Println(err.Error())
	}
}
