package main

import (
	"go.uber.org/dig"
	"log"
	"rule-based-wordle-cracker/src/game_engine"
	"rule-based-wordle-cracker/src/game_state"
	"rule-based-wordle-cracker/src/lookup_table"
)

func getLookupTable() *lookup_table.LookupTable {
	lookupTable, err := lookup_table.LookupTableFromFile("five_letter_words.txt")
	if err != nil {
		log.Fatalln(err)
	}
	return lookupTable
}

func buildContainer() *dig.Container {
	container := dig.New()

	container.Provide(getLookupTable)
	container.Provide(game_state.NewGameState)

	return container
}

func main() {
	if err := buildContainer().Invoke(game_engine.Run); err != nil {
		log.Fatalln(err)
	}
}
