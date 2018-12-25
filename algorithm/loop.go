package algorithm

import (
	"github.com/edas11/hyphenation/util/text"
	"strings"
	"unicode"
)

type LoopAlgorithm struct {
	patterns        []string
	reducedPatterns []string
}

func NewLoopAlgorithm(hyphenationPatterns []string) LoopAlgorithm {
	reducedPatterns := make([]string, len(hyphenationPatterns))
	for i, pattern := range hyphenationPatterns {
		reducedPatterns[i] = strings.Map(text.RemoveDigit, pattern)
	}
	return LoopAlgorithm{hyphenationPatterns, reducedPatterns}
}

func (algorithm LoopAlgorithm) HyphenateWord(wordToHyphenate string) string {
	word := newWord(wordToHyphenate)
	for i, pattern := range algorithm.patterns {
		if strings.HasPrefix(pattern, ".") {
			word.matchBeginningPattern(pattern, algorithm.reducedPatterns[i])
		} else if strings.HasSuffix(pattern, ".") {
			word.matchEndPattern(pattern, algorithm.reducedPatterns[i])
		} else {
			word.matchGeneralPattern(pattern, algorithm.reducedPatterns[i])
		}
	}
	return word.generateHyphenatedWord()
}

func (word word) matchBeginningPattern(pattern, reducedPattern string) {
	if strings.HasPrefix(word.wordToHyphenate, reducedPattern) {
		word.updateWordHyphenationNumbers(pattern, 0)
	}
}

func (word word) matchEndPattern(pattern, reducedPattern string) {
	if strings.HasSuffix(word.wordToHyphenate, reducedPattern) {
		word.updateWordHyphenationNumbers(pattern, len(word.wordToHyphenate)-len(reducedPattern))
	}
}

func (word word) matchGeneralPattern(pattern, reducedPattern string) {
	matchIndex := text.IndexFrom(word.wordToHyphenate, reducedPattern, 0)
	for matchIndex != -1 {
		word.updateWordHyphenationNumbers(pattern, matchIndex)
		matchIndex = text.IndexFrom(word.wordToHyphenate, reducedPattern, matchIndex+1)
	}
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
