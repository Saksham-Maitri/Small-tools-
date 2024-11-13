package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func fetchinfo(text string) (int, int, int) {
	lines := strings.Split(text, "\n")
	wordsCount := 0
	linesCount := len(lines)
	charCount := 0

	for _, line := range lines {
		words := strings.Fields(line)
		wordsCount += len(words)

		// Count characters excluding spaces
		for _, char := range line {
			if char != ' ' {
				charCount++
			}
		}
	}
	return linesCount, wordsCount, charCount
}

func main() {
	fmt.Println("Hello, User!")

	// Define flags
	filepath := flag.String("file", "", "File to read")
	textline := flag.String("text", "", "Text to read")
	flag.Parse()

	var text string

	// Handle text input
	if *textline != "" {
		text = *textline
	} else if *filepath != "" {
		file, err := os.Open(*filepath)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error encountered while reading file: %s\n", err)
			return
		}
	} else {
		fmt.Println("No text or file provided. Please provide it using the -text or -file flag.")
		return
	}

	lines, words, chars := fetchinfo(text)
	fmt.Printf("Lines: %d\nWords: %d\nCharacters (excluding spaces): %d\n", lines, words, chars)
}
