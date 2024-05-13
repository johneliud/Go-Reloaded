package main

import (
	"fmt"
	"os"
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

func main() {
	var sampleTxt []byte
	var sliceOfWords []string
	var err error
	var joinedWords string
	var resultTxt *os.File

	if len(os.Args) == 3 {
		arguments := os.Args[1:]

		sampleTxt, err = os.ReadFile(arguments[0])
		if err != nil {
			fmt.Println("Error reading sample.txt file!", err)
			return
		}

		if len(sampleTxt) == 0 {
			fmt.Println("File to read from is empty!")
			return
		}

		sliceOfWords = strings.Split(string(sampleTxt), " ")

		// Pass the sliceOfWords to various appropriate functions
		sliceOfWords = hexadecimal.HexadecimalToDecimal(sliceOfWords)
		sliceOfWords = binary.BinaryToDecimal(sliceOfWords)
		sliceOfWords = vowel.CorrectVowels(sliceOfWords)
		sliceOfWords = uppercase.ConvertToUpperCase(sliceOfWords)
		sliceOfWords = lowercase.ConvertToLowerCase(sliceOfWords)
		sliceOfWords = capitalize.Capitalize(sliceOfWords)
		sliceOfWords = punctuation.CheckPunctuation(sliceOfWords)
		sliceOfWords = apostrophe.CheckApostrophe(sliceOfWords)

		joinedWords = strings.Join(sliceOfWords, " ") + "\n"

		resultTxt, err = os.Create(arguments[1])
		if err != nil {
			fmt.Println("Error creating result.txt file!", err)
			return
		}

		_, err = resultTxt.Write([]byte(joinedWords))
		if err != nil {
			fmt.Println("Error writing to result.txt file!", err)
			return
		}
	} else if len(os.Args) == 1 {
		sampleTxt, err = os.ReadFile("sample.txt")
		if err != nil {
			fmt.Println("Error reading sample.txt file!", err)
			return
		}

		if len(sampleTxt) == 0 {
			fmt.Println("File is read from is empty!")
			return
		}

		sliceOfWords = strings.Split(string(sampleTxt), " ")

		// Pass the sliceOfWords to various appropriate functions
		sliceOfWords = hexadecimal.HexadecimalToDecimal(sliceOfWords)
		sliceOfWords = binary.BinaryToDecimal(sliceOfWords)
		sliceOfWords = vowel.CorrectVowels(sliceOfWords)
		sliceOfWords = uppercase.ConvertToUpperCase(sliceOfWords)
		sliceOfWords = lowercase.ConvertToLowerCase(sliceOfWords)
		sliceOfWords = capitalize.Capitalize(sliceOfWords)
		sliceOfWords = punctuation.CheckPunctuation(sliceOfWords)
		sliceOfWords = apostrophe.CheckApostrophe(sliceOfWords)

		joinedWords = strings.Join(sliceOfWords, " ") + "\n"

		resultTxt, err = os.Create("result.txt")
		if err != nil {
			fmt.Println("Error creating result.txt file!", err)
			return
		}

		_, err = resultTxt.Write([]byte(joinedWords))
		if err != nil {
			fmt.Println("Error writing to result.txt file!", err)
			return
		}
	} else {
		fmt.Println("Incorrect number of arguments passed! Usage: 'go run .' or 'go run . sample.txt result.txt'")
	}
}
