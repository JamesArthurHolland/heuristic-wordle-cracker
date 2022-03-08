package lookup_table

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"rule-based-wordle-cracker/src/utils"
)

type LookupTable struct {
	allFiveLetterWords []string                    `json:"all_five_letter_words"`
	byLetter           map[string][]string         `json:"by_letter"`
	byLetterPosition   map[int]map[string][]string `json:"by_position"`
}

func NewLookupTable(allFiveLetterWords []string, byLetter map[string][]string, byLetterPosition map[int]map[string][]string) *LookupTable {
	return &LookupTable{
		allFiveLetterWords: allFiveLetterWords,
		byLetter:           byLetter,
		byLetterPosition:   byLetterPosition,
	}
}

func (lt *LookupTable) ByLetter(letter string) []string {
	return lt.byLetter[letter]
}

func (lt *LookupTable) ByLetterAtPosition(letter string, position int) []string {
	return lt.byLetterPosition[position][letter]
}

func (lt *LookupTable) GetAllWords() []string {
	newSlice := append([]string(nil), lt.allFiveLetterWords...)
	return newSlice
}

func (lt *LookupTable) LookupTableToFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	jsonBytes, err := json.Marshal(lt)
	if err != nil {
		log.Fatalln(err)
	}

	file.Write(jsonBytes)
}

func LookupTableFromFile(filename string) (*LookupTable, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var allWords []string
	byLetter := make(map[string][]string)
	byLetterPosition := make(map[int]map[string][]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		allWords = append(allWords, word)
		chars := []rune(word)

		for i := 0; i < len(chars); i++ {
			char := string(chars[i])

			if !utils.SliceHasString(byLetter[char], word) {
				byLetter[char] = append(byLetter[char], word)
			}
			if _, exists := byLetterPosition[i]; !exists {
				byLetterPosition[i] = make(map[string][]string)
			}
			if !utils.SliceHasString(byLetterPosition[i][char], word) {
				byLetterPosition[i][char] = append(byLetterPosition[i][char], word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return NewLookupTable(allWords, byLetter, byLetterPosition), nil
}
