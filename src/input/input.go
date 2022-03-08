package input

import (
	"errors"
	"log"
	"regexp"
	"rule-based-wordle-cracker/src/game_state"
	"strconv"
)

func HandleInput(givenInput string) (game_state.Row, error) {
	regex, err := regexp.Compile("([a-z][0-2])")
	if err != nil {
		log.Fatal(err)
	}

	res := regex.FindAllStringSubmatch(givenInput, -1)
	if len(res) != 5 {
		return nil, errors.New("Input didn't match")
	}
	var cells []*game_state.Cell
	for i, _ := range res {
		cellString := res[i][1]
		chars := []rune(cellString)
		parsedFeedback, err := strconv.ParseInt(string(chars[1]), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		cell := &game_state.Cell{
			Letter:   string(chars[0]),
			Feedback: game_state.CellFeedback(parsedFeedback),
		}
		cells = append(cells, cell)
	}

	return cells, nil
}
