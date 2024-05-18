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

func main() {
	if len(os.Args) == 3 {
		arguments := os.Args[1:]

		sampleTxt, err := os.Open(arguments[0])
		if err != nil {
			fmt.Println("Error opening input file!", err)
			return
		}
		defer sampleTxt.Close()

		// Variable fileDetails contains statistics/properties associated the input file
		fileDetails, err := os.Stat(arguments[0])
		if err != nil {
			fmt.Println("Error opening input file!", err)
			return
		}

		// Check if input file is empty and display message
		if fileDetails.Size() == 0 {
			fmt.Println("Input file is empty!")
			return
		}

		// Check for the correct extension types for both input and output file
		if filepath.Ext(arguments[0]) != ".txt" || filepath.Ext(arguments[1]) != ".txt" {
			fmt.Println("Invalid type of extension used on file!")
			return
		}

		resultTxt, err := os.Create(arguments[1])
		if err != nil {
			fmt.Println("Error creating file!", err)
			return
		}
		defer resultTxt.Close()

		scanner := bufio.NewScanner(sampleTxt)

		for scanner.Scan() {
			readLines := scanner.Text()

			sliceOfWords := strings.Split(string(readLines), " ")

			// Pass the sliceOfWords to various functions
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
			}
		}
	} else if len(os.Args) == 1 {
		sampleTxt, err := os.Open("sample.txt")
		if err != nil {
			fmt.Println("Error opening input file!", err)
			return
		}
		defer sampleTxt.Close()

		// Variable fileDetails contains statistics/properties of a file
		fileDetails, err := os.Stat("sample.txt")
		if err != nil {
			fmt.Println("Error opening input file!", err)
			return
		}

		// Check if input file is empty and display message
		if fileDetails.Size() == 0 {
			fmt.Println("Input file is empty!")
			return
		}

		// Check for the correct extension types for both input and output file
		if filepath.Ext("sample.txt") != ".txt" || filepath.Ext("result.txt") != ".txt" {
			fmt.Println("Invalid type of extension used on file!")
			return
		}

		resultTxt, err := os.Create("result.txt")
		if err != nil {
			fmt.Println("Error creating file!", err)
			return
		}
		defer resultTxt.Close()

		scanner := bufio.NewScanner(sampleTxt)

		for scanner.Scan() {
			readLines := scanner.Text()

			sliceOfWords := strings.Split(string(readLines), " ")

			// Pass the sliceOfWords to various functions
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
			}
		}
	} else {
		fmt.Println("Incorrect number of arguments passed! Use 'go run .' or 'go run . sample.txt result.txt'")
	}
}
