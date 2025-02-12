package input

import (
	"errors"
	"log"
	"regexp"
	"rule-based-wordle-cracker/src/game_state"
	"strconv"
)

func getInputErrorMessage() error {
	return errors.New("Input didn't match, please input 0 if the letter is not in the word, 1 if the letter " +
		"is in the word but not in the correct position, and 2 if the letter is in the word and in the correct position")
}

func HandleInput(suggestionWord, givenInput string) (game_state.Row, error) {
	suggestionRunes := []rune(suggestionWord)
	regex, err := regexp.Compile("[0-2]")
	if err != nil {
		log.Fatal(err)
	}

	res := regex.FindAllStringSubmatch(givenInput, -1)
	if len(res) != 5 {
		return nil, errors.New("Input didn't match, please input 0 if the letter is not in the word, 1 if the letter " +
			"is in the word but not in the correct position, and 2 if the letter is in the word and in the correct position")
	}
	var cells []*game_state.Cell
	for i, _ := range res {
		inputNumberString := res[i][0]
		parsedFeedback, err := strconv.ParseInt(inputNumberString, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		cell := &game_state.Cell{
			Letter:   string(suggestionRunes[i]),
			Feedback: game_state.CellFeedback(parsedFeedback),
		}
		cells = append(cells, cell)
	}

	return cells, nil
}
