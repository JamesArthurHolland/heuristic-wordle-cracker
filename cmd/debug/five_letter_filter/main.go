package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func readWordsDictionaryPrintFiveLetterWords(w *bufio.Writer) {
	file, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Regular expression that matches a string that "starts, has 5 letters, then ends".
	regex, err := regexp.Compile("^[a-z]{5}$")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Remove any accidental spaces
		line = strings.TrimSpace(line)

		// Make everything lowercase
		line = strings.ToLower(line)

		if regex.MatchString(line) {
			w.WriteString(line + "\n")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	file, err := os.OpenFile("five_letter_words.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	readWordsDictionaryPrintFiveLetterWords(datawriter)

	datawriter.Flush()
	file.Close()
}
