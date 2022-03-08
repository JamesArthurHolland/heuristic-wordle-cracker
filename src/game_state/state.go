package game_state

import (
	"rule-based-wordle-cracker/src/lookup_table"
	"rule-based-wordle-cracker/src/move_suggestor"
	"rule-based-wordle-cracker/src/utils"
)
import "github.com/juliangruber/go-intersect"

type GameState struct {
	lookupTable   *lookup_table.LookupTable
	possibleWords WordList
}

type Cell struct {
	Letter   string
	Feedback CellFeedback
}

type Row []*Cell

type CellFeedback int

const (
	NotInWord CellFeedback = iota
	InWordNotInPosition
	InWordInPosition
)

type WordList []string

func NewGameState(lookupTable *lookup_table.LookupTable) *GameState {
	return &GameState{
		lookupTable:   lookupTable,
		possibleWords: lookupTable.GetAllWords(),
	}
}

func (gs *GameState) SuggestNextWord() string {
	return move_suggestor.SuggestMove(gs.possibleWords)
}

func (gs *GameState) ProcessRow(row Row) {
	for i, cell := range row {
		switch cell.Feedback {
		case NotInWord:
			gs.possibleWords = utils.RemoveWordsContainingLetter(gs.possibleWords, cell.Letter)
		case InWordNotInPosition:
			sliceInterface := intersect.Simple(gs.lookupTable.ByLetter(cell.Letter), gs.possibleWords)
			gs.possibleWords = utils.InterfaceSliceStringsToStringSlice(sliceInterface)
			gs.possibleWords = utils.RemoveWordsContainingLetterAtPosition(gs.possibleWords, cell.Letter, i)
		case InWordInPosition:
			sliceInterface := intersect.Simple(gs.lookupTable.ByLetterAtPosition(cell.Letter, i), gs.possibleWords)
			gs.possibleWords = utils.InterfaceSliceStringsToStringSlice(sliceInterface)
		}
	}
}
