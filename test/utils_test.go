package test

import (
	"rule-based-wordle-cracker/src/utils"
	"testing"
)

func TestWordRemovalByLetter(t *testing.T) {
	words := []string{"hoard", "nomad", "altas", "conch"}
	remaining := utils.RemoveWordsContainingLetter(words, "a")
	if len(remaining) != 1 {
		t.Error("Should only be 1 word remaining. \"conch\"")
	}
}

func TestWordRemovalByLetterPosition(t *testing.T) {
	words := []string{"hoard", "nomad", "altas", "conch"}
	remaining := utils.RemoveWordsContainingLetterAtPosition(words, "o", 1)
	if len(remaining) != 1 {
		t.Error("Should only be 1 word remaining. \"atlas\"")
	}
}

func TestLetterFrequencies(t *testing.T) {
	words := []string{"hoard", "nomad", "altas", "conch"}
	frequencies := utils.CalculateLetterFrequencies(words)

	if frequencies["o"] != 3 {
		t.Error("Should be 3 \"o\"s")
	}
	if frequencies["r"] != 1 {
		t.Error("Should be 1 \"r\"")
	}
}
