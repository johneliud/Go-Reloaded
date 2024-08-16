package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-reloaded/apostrophe"
	"go-reloaded/binary"
	"go-reloaded/capitalize"
	"go-reloaded/hexadecimal"
	"go-reloaded/lowercase"
	"go-reloaded/punctuation"
	"go-reloaded/uppercase"
	"go-reloaded/vowel"
)

func processFile(inputFile, outputFile string) {
	sampleTxt, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file!", err)
		return
	}
	defer sampleTxt.Close()

	fileDetails, err := os.Stat(inputFile)
	if err != nil {
		fmt.Println("Error opening input file!", err)
		return
	}

	if fileDetails.Size() == 0 {
		fmt.Println("Input file is empty!")
		return
	}

	if filepath.Ext(inputFile) != ".txt" || filepath.Ext(outputFile) != ".txt" {
		fmt.Println("Invalid type of extension used on file!")
		return
	}

	resultTxt, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file!", err)
		return
	}
	defer resultTxt.Close()

	scanner := bufio.NewScanner(sampleTxt)

	for scanner.Scan() {
		readLines := scanner.Text()

		sliceOfWords := strings.Split(readLines, " ")

		sliceOfWords = hexadecimal.HexadecimalToDecimal(sliceOfWords)
		sliceOfWords = binary.BinaryToDecimal(sliceOfWords)
		sliceOfWords = vowel.CorrectVowels(sliceOfWords)
		sliceOfWords = uppercase.ConvertToUpperCase(sliceOfWords)
		sliceOfWords = lowercase.ConvertToLowerCase(sliceOfWords)
		sliceOfWords = capitalize.Capitalize(sliceOfWords)
		sliceOfWords = punctuation.CheckPunctuation(sliceOfWords)
		sliceOfWords = apostrophe.CheckApostrophe(sliceOfWords)

		joinedWords := strings.Join(sliceOfWords, " ")

		_, err := resultTxt.Write([]byte(joinedWords + "\n"))
		if err != nil {
			fmt.Println("Error writing to output file!", err)
			return
		} else {
			fmt.Println("File successfully written!")
		}
	}
}

func main() {
	if len(os.Args) == 3 {
		processFile(os.Args[1], os.Args[2])
	} else if len(os.Args) == 1 {
		processFile("sample.txt", "result.txt")
	} else {
		fmt.Println("Incorrect number of arguments passed! Use 'go run .' or 'go run . sample.txt result.txt'")
	}
}
