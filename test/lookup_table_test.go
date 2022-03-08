package test

import (
	"log"
	"rule-based-wordle-cracker/src/lookup_table"
	"testing"
)

func TestByLetterPosition(t *testing.T) {
	lookupTable, err := lookup_table.LookupTableFromFile("../five_letter_word_test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	if len(lookupTable.ByLetterAtPosition("a", 0)) != 1 {
		t.Error("Test list should have only 1 word with a at the start, `atlas`")
	}
}
