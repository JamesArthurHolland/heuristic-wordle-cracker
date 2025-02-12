package game_engine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"rule-based-wordle-cracker/src/game_state"
	"rule-based-wordle-cracker/src/input"
)

func Run(state *game_state.GameState) {
	for {
		suggestedWord := state.SuggestNextWord()
		fmt.Printf("Suggestion: %s \n", suggestedWord)
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter input: ")
		givenInput, _ := reader.ReadString('\n')
		row, err := input.HandleInput(suggestedWord, givenInput)
		if err != nil {
			log.Println(err)
			continue
		}
		state.ProcessRow(row)

	}
}
