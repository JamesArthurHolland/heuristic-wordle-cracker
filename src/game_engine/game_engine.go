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
	log.Printf("Suggestion: %s", state.SuggestNextWord())
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter input: ")
		givenInput, _ := reader.ReadString('\n')
		row, err := input.HandleInput(givenInput)
		if err != nil {
			log.Println(err)
			continue
		}
		state.ProcessRow(row)
		log.Printf("Suggestion: %s", state.SuggestNextWord())
	}
}
