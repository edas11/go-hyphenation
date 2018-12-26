package algorithm

import "unicode"

type word struct {
	wordToHyphenate string
	wordLength int
	hyphenationNumbers []rune // Represent gaps in word
}

func newWord(wordToHyphenate string) word {
	return word{wordToHyphenate, len(wordToHyphenate), make([]rune, len(wordToHyphenate)+1)}
}

func (word word) updateWordHyphenationNumbers(pattern string, matchIndx int) {
	currentWordGapIndx := 0
	for _, patternChar := range pattern {
		if patternChar == '.' {
			continue
		} else if unicode.IsDigit(patternChar) {
			word.replaceDigitIfLarger(matchIndx+currentWordGapIndx, patternChar)
		} else {
			currentWordGapIndx++
		}
	}
}

func (word word) replaceDigitIfLarger(wordNumberIndx int, patternDigit rune) {
	if patternDigit > word.hyphenationNumbers[wordNumberIndx] {
		word.hyphenationNumbers[wordNumberIndx] = patternDigit
	}
}

func (word word) generateHyphenatedWord() string {
	hyphenatedWord := []rune(word.wordToHyphenate)
	numOfCuts := 0
	for indx := range word.hyphenationNumbers {
		if word.pointsTofirstOrLastNumber(indx) {
			continue
		}
		if isOdd(word.digitAtGap(indx)) {
			cutPoint := indx + numOfCuts
			hyphenatedWord = insertCharAt(hyphenatedWord, cutPoint, '-')
			numOfCuts++
		}
	}
	return string(hyphenatedWord)
}

func (word word) pointsTofirstOrLastNumber(indx int) bool {
	return indx == 0 || indx == len(word.hyphenationNumbers)-1
}

func (word word) digitAtGap(indx int) int {
	return int(word.hyphenationNumbers[indx]) - '0'
}

func isOdd(number int) bool {
	return number%2 == 1
}

func insertCharAt(text []rune, at int, char rune) []rune {
	newText := append(text, ' ')
	copy(newText[at+1:], newText[at:])
	newText[at] = char
	return newText
}
