## Project Description

Go-Reloaded is a simple program that acts as a text completion/editing/auto-correction tool and was built using Go programming language.

## Features

- Reads data from a file (sample.txt) and writes the corrected/edited version of the data to another file (result.txt)

Below is a list of possible modifications that the program should execute:

- Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

- Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

- Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

- Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

- Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")

1. For (low), (up), (cap) if a number appears next to it, like so: (low, _number_) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").

1. Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

2. The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

3. If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known musician in the world '" -> "As Elton John said: 'I am the most well-known musician in the world'")

- Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").

## Getting Started

To be able to interact with the program on your local machine, you will need to have Go language installed on your machine. Use the link below to download and install Golang.

https://go.dev/doc/install

## Installation

1. Clone the repository:

   ```
   $ git clone https://github.com/johneliud/Go-Reloaded.git
   ```

2. Navigate into the go-reloaded directory:
   ```
   $ cd Go-Reloaded
   ```

## Usage

- You can the program with the following commands

  ```
  $ go run . sample.txt result.txt
  ```

  or

  ```
  $ go run .
  ```

## Examples

- Example 1

1. Write the text "1E (hex) files were added" to the sample.txt
   ```bash
   $ echo "1E (hex) files were added" > sample.txt
   ```
2. Use cat command to view contents in sample.txt

   ```bash
   $ cat sample.txt
   ```

   Output:

   ```bash
   1E (hex) files were added
   ```

3. Use the run command to copy contents of sample.txt to result.txt
   ```
   $ go run . sample.txt result.txt
   ```
4. Use cat command to view contents in sample.txt

   ```bash
   $ cat result.txt
   ```

   Output:

   ```bash
   30 files were added
   ```

- Example 2

1. Write the text "Ready, set, go (up) !" to the sample.txt
   ```bash
   $ echo "Ready, set, go (up) !" > sample.txt
   ```
2. Use cat command to view contents in sample.txt

   ```bash
   $ cat sample.txt
   ```

   Output:

   ```bash
   Ready, set, go (up) !
   ```

3. Use the run command to copy contents of sample.txt to result.txt
   ```
   $ go run . sample.txt result.txt
   ```
4. Use cat command to view contents in sample.txt

   ```bash
   $ cat result.txt
   ```

   Output:

   ```bash
   Ready, set, GO!
   ```

- Example 3

1. Write the text "This is so exciting (up, 2)" to the sample.txt
   ```bash
   $ echo "This is so exciting (up, 2)" > sample.txt
   ```
2. Use cat command to view contents in sample.txt

   ```bash
   $ cat sample.txt
   ```

   Output:

   ```bash
   This is so exciting (up, 2)
   ```

3. Use the run command to copy contents of sample.txt to result.txt
   ```
   $ go run . sample.txt result.txt
   ```
4. Use cat command to view contents in sample.txt

   ```bash
   $ cat result.txt
   ```

   Output:

   ```bash
   This is SO EXCITING
   ```

- Example 4

1. Write the text "I was thinking ... You were right" to the sample.txt
   ```bash
   $ echo "I was thinking ... You were right" > sample.txt
   ```
2. Use cat command to view contents in sample.txt

   ```bash
   $ cat sample.txt
   ```

   Output:

   ```bash
   I was thinking ... You were right
   ```

3. Use the run command to copy contents of sample.txt to result.txt
   ```
   $ go run . sample.txt result.txt
   ```
4. Use cat command to view contents in sample.txt

   ```bash
   $ cat result.txt
   ```

   Output:

   ```bash
   I was thinking... You were right
   ```

- Example 5

1. Write the text "There it was. A amazing rock!" to the sample.txt
   ```bash
   $ echo "There it was. A amazing rock!" > sample.txt
   ```
2. Use cat command to view contents in sample.txt

   ```bash
   $ cat sample.txt
   ```

   Output:

   ```bash
   There it was. A amazing rock!
   ```

3. Use the run command to copy contents of sample.txt to result.txt
   ```
   $ go run . sample.txt result.txt
   ```
4. Use cat command to view contents in sample.txt

   ```bash
   $ cat result.txt
   ```

   Output:

   ```bash
   There it was. An amazing rock!
   ```

## Contact

Incase of further enquiries, please reach out via the email address johneliud4@gmail.com
