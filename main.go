package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ConvertToFont() [95][8]string {

	file, err := os.Open("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	// create an 2d array that will hold the number of characters
	// and the number of rows there are in each character
	storeCharacters := [95][8]string{}
	// since there are white spacees start at 0 and ends at 94
	// we are starting at -1 to start counting characters
	charCount := -1
	rowCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		// when a white space is found this indicates a line break so increase the character count
		if line == "" {
			// fmt.Printf("count %d - value: %v\n", count, scanner.Text())
			charCount++
			// dont increase the row count when there is a white space
			rowCount = 0
		} else {
			// store the value of each character into a row
			storeCharacters[charCount][rowCount] = scanner.Text()
			// increase row count
			rowCount++
		}
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
	return storeCharacters
}

func IndividualCharacter(c rune, s [95][8]string) []string {
	// store the rows of the character into a slice of string
	char := []string{}
	// minus the ascii value from 32 to give the corrresponding index of the character in the 2d array
	c = c - 32
	for i := 0; i < 8; i++ {
		char = append(char, s[c][i])
	}
	return char
}

func main() {
	test := "Hello"
	runeStr := []rune(test)
	font := ConvertToFont()

	result := [][]string{}
	for x := 0; x < len(runeStr); x++ {
		result = append(result, IndividualCharacter(runeStr[x], font))
	}
	// }

	// fmt.Println(result)
	for i := range result {
		for _, v := range result[i] {
			fmt.Println(v)
		}
	}
}
