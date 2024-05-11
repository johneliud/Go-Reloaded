## Project Description

Go-Reloaded is a simple program that acts as a text completion/editing/auto-correction tool and was built using Go programming langauge.

## Features

- Reads data from a file(sample.txt) and writes the corrected/edited version of the data to a new file(result.txt)

## Prerequisites
Below is a list of possible modifications that the program should execute:

- Every instance of (hex) should replace the word before with the decimal version of the word (in this case the word will always be a hexadecimal number). (Ex: "1E (hex) files were added" -> "30 files were added")

- Every instance of (bin) should replace the word before with the decimal version of the word (in this case the word will always be a binary number). (Ex: "It has been 10 (bin) years" -> "It has been 2 years")

- Every instance of (up) converts the word before with the Uppercase version of it. (Ex: "Ready, set, go (up) !" -> "Ready, set, GO !")

- Every instance of (low) converts the word before with the Lowercase version of it. (Ex: "I should stop SHOUTING (low)" -> "I should stop shouting")

- Every instance of (cap) converts the word before with the capitalized version of it. (Ex: "Welcome to the Brooklyn bridge (cap)" -> "Welcome to the Brooklyn Bridge")
1. For (low), (up), (cap) if a number appears next to it, like so: (low, <number>) it turns the previously specified number of words in lowercase, uppercase or capitalized accordingly. (Ex: "This is so exciting (up, 2)" -> "This is SO EXCITING")

- Every instance of the punctuations ., ,, !, ?, : and ; should be close to the previous word and with space apart from the next one. (Ex: "I was sitting over there ,and then BAMM !!" -> "I was sitting over there, and then BAMM!!").
1. Except if there are groups of punctuation like: ... or !?. In this case the program should format the text as in the following example: "I was thinking ... You were right" -> "I was thinking... You were right".

2. The punctuation mark ' will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces. (Ex: "I am exactly how they describe me: ' awesome '" -> "I am exactly how they describe me: 'awesome'")

3. If there are more than one word between the two ' ' marks, the program should place the marks next to the corresponding words (Ex: "As Elton John said: ' I am the most well-known homosexual in the world '" -> "As Elton John said: 'I am the most well-known homosexual in the world'")

- Every instance of a should be turned into an if the next word begins with a vowel (a, e, i, o, u) or a h. (Ex: "There it was. A amazing rock!" -> "There it was. An amazing rock!").


## Installation

1. Clone the repository:

   ```
   $ git clone https://learn.zone01kisumu.ke/git/johnodhiambo0/go-reloaded.git
   ```

2. Navigate into the go-reloaded directory:
   ```
   $ cd go-reloaded
   ```

## Usage

- Run the program with the folowing commands
    ```
    $ go run . sample.txt result.txt
    ```

    ```
    $ go run .
    ```

## Examples

- Write the text "1E (hex) files were added" to the sample.txt
   ```bash
   $ echo "1E (hex) files were added" > sample.txt
   ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat sample.txt
   ```

   Output:
   ```bash
   1E (hex) files were added
   ```
- Use the run command to copy contents of sample.txt to result.txt
    ```
    $ go run . sample.txt result.txt
    ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat result.txt
   ```

    Output:
   ```bash
   30 files were added
   ```

- Write the text "Ready, set, go (up) !" to the sample.txt
   ```bash
   $ echo "Ready, set, go (up) !" > sample.txt
   ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat sample.txt
   ```

   Output:
   ```bash
   Ready, set, go (up) !
   ```
- Use the run command to copy contents of sample.txt to result.txt
    ```
    $ go run . sample.txt result.txt
    ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat result.txt
   ```

    Output:
   ```bash
   Ready, set, GO!
   ```

- Write the text "This is so exciting (up, 2)" to the sample.txt
   ```bash
   $ echo "This is so exciting (up, 2)" > sample.txt
   ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat sample.txt
   ```

   Output:
   ```bash
   This is so exciting (up, 2)
   ```
- Use the run command to copy contents of sample.txt to result.txt
    ```
    $ go run . sample.txt result.txt
    ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat result.txt
   ```

    Output:
   ```bash
   This is SO EXCITING
   ```

- Write the text "I was thinking ... You were right" to the sample.txt
   ```bash
   $ echo "I was thinking ... You were right" > sample.txt
   ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat sample.txt
   ```

   Output:
   ```bash
   I was thinking ... You were right
   ```
- Use the run command to copy contents of sample.txt to result.txt
    ```
    $ go run . sample.txt result.txt
    ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat result.txt
   ```

    Output:
   ```bash
   I was thinking... You were right
   ```

- Write the text "There it was. A amazing rock!" to the sample.txt
   ```bash
   $ echo "There it was. A amazing rock!" > sample.txt
   ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat sample.txt
   ```

   Output:
   ```bash
   There it was. A amazing rock!
   ```
- Use the run command to copy contents of sample.txt to result.txt
    ```
    $ go run . sample.txt result.txt
    ```
- Use cat command to view contents in sample.txt
   ```bash
   $ cat result.txt
   ```

    Output:
   ```bash
   There it was. An amazing rock!
   ```