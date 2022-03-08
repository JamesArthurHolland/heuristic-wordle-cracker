package main

import (
	"log"
	"rule-based-wordle-cracker/src/lookup_table"
)

func main() {
	lookupTable, err := lookup_table.LookupTableFromFile("five_letter_word_test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lookupTable.LookupTableToFile("lookup.json")
}
