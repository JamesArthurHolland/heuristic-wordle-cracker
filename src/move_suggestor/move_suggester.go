package move_suggestor

import "rule-based-wordle-cracker/src/utils"

func SuggestMove(possibleWords []string) string {
	letterFrequencies := utils.CalculateLetterFrequencies(possibleWords)
	return utils.CalculateBestFrequencyWord(possibleWords, letterFrequencies)
}
