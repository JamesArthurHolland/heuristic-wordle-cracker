package test

import (
	"log"
	"rule-based-wordle-cracker/src/input"
	"testing"
)

func TestInput(t *testing.T) {
	givenInput := "h0o2a1r0d0"
	row, err := input.HandleInput(givenInput)

	log.Println(row)
	log.Println(err)
}
