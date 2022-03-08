package utils

func SliceHasString(slice []string, givenString string) bool {
	for _, existingString := range slice {
		if existingString == givenString {
			return true
		}
	}
	return false
}

func InterfaceSliceStringsToStringSlice(sliceInterface []interface{}) []string {
	var words []string
	for _, item := range sliceInterface {
		words = append(words, item.(string))
	}
	return words
}

func RemoveWordsContainingLetter(words []string, letter string) []string {
	var newWords []string
	for _, word := range words {
		chars := []rune(word)
		charFound := false
		for i := 0; i < len(chars); i++ {
			char := string(chars[i])
			if char == letter {
				charFound = true
			}
		}
		if charFound == false {
			newWords = append(newWords, word)
		}
	}
	return newWords
}

func RemoveWordsContainingLetterAtPosition(words []string, letter string, position int) []string {
	var newWords []string
	for _, word := range words {
		chars := []rune(word)

		if string(chars[position]) != letter {
			newWords = append(newWords, word)
		}
	}
	return newWords
}

func CalculateLetterFrequencies(words []string) map[string]int {
	frequencies := make(map[string]int)

	for _, word := range words {
		chars := []rune(word)
		for i := 0; i < len(chars); i++ {
			char := string(chars[i])
			frequencies[char]++
		}
	}

	return frequencies
}

func CalculateWordLetterFrequencySum(word string, frequencies map[string]int) int {
	chars := []rune(word)
	frequencySum := 0
	letterCache := make(map[string]struct{}) // TODO document why you must not include repeated letters in the total.

	for i := 0; i < len(chars); i++ {
		char := string(chars[i])
		if _, exists := letterCache[char]; !exists {
			frequencySum += frequencies[char]
			letterCache[char] = struct{}{}
		}
	}
	return frequencySum
}

func CalculateBestFrequencyWord(words []string, frequencies map[string]int) string {
	bestWord := ""
	bestFrequency := 0
	for _, word := range words {
		wordFrequency := CalculateWordLetterFrequencySum(word, frequencies)
		if wordFrequency > bestFrequency {
			bestWord = word
			bestFrequency = wordFrequency
		}
	}
	return bestWord
}
