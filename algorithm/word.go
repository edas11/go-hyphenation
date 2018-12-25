package algorithm

type word struct {
	wordToHyphenate    string
	hyphenationNumbers []rune
}

func newWord(wordToHyphenate string) word {
	return word{wordToHyphenate, make([]rune, len(wordToHyphenate)+1)}
}

func (word word) generateHyphenatedWord() string {
	hyphenatedWord := word.wordToHyphenate
	numOfCuts := 0
	for indx := range word.hyphenationNumbers {
		if word.pointsTofirstOrLastNumber(indx) {
			continue
		}
		if isOdd(word.digitAtGap(indx)) {
			cutPoint := indx + numOfCuts
			hyphenatedWord = hyphenatedWord[:cutPoint] + "-" + hyphenatedWord[cutPoint:]
			numOfCuts++
		}
	}
	return hyphenatedWord
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
