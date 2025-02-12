package test

import (
	"log"
	"rule-based-wordle-cracker/src/input"
	"testing"
)

func TestInput(t *testing.T) {
	suggestionWord := "hoard"
	givenInput := "02100"
	row, err := input.HandleInput(suggestionWord, givenInput)

	log.Println(row)
	log.Println(err)
}
