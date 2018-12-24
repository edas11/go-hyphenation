package main

import (
	//"os"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

var wordToHyphenate string
var wordHyphenationNumbers []rune

func main() {
	//hyphenationArguments := os.Args
	wordToHyphenate = "recursion"
	wordHyphenationNumbers = make([]rune, len(wordToHyphenate) + 1)
	patternsFileContent, err := ioutil.ReadFile("tex-hyphenation-patterns.txt")
	if err != nil {
		return
	}
	hyphenationPatterns := strings.Split(string(patternsFileContent), "\n")
	for _, pattern := range hyphenationPatterns {
		reducedPattern := strings.Map(removeDigit, pattern)
		if strings.HasPrefix(pattern, ".") {
			if strings.HasPrefix(wordToHyphenate, reducedPattern) {
				updateWordHyphenationNumbers(pattern, 0)
			}
			continue
		} else if strings.HasSuffix(pattern, ".") {
			if strings.HasSuffix(wordToHyphenate, reducedPattern) {
				updateWordHyphenationNumbers(pattern, len(wordToHyphenate) - len(reducedPattern))
			}
			continue
		} else {
			matchIndex := strings.Index(wordToHyphenate, reducedPattern)
			for matchIndex != -1 {
				updateWordHyphenationNumbers(pattern, matchIndex)
				if len(wordToHyphenate) - 1 != matchIndex {
					matchIndex = strings.Index(wordToHyphenate[matchIndex + 1:], reducedPattern)
				}
			}
		}
	}
	fmt.Println(generateHyphenatedWord())
}

func generateHyphenatedWord() string {
	hyphenatedWord := wordToHyphenate
	numOfCuts := 0
	for indx := range wordHyphenationNumbers {
		if indx == 0 || indx == len(wordHyphenationNumbers) - 1 {
			continue
		}
		number := int(wordHyphenationNumbers[indx]) - '0'
		if number%2 == 1 {
			cutPoint := indx + numOfCuts
			hyphenatedWord = hyphenatedWord[:cutPoint] + "-" + hyphenatedWord[cutPoint:]
			numOfCuts += 1
		}
	}
	return hyphenatedWord
}

func updateWordHyphenationNumbers(pattern string, matchIndx int) {
	currentWordGapIndx := 0;
	for _, rune := range pattern {
		if rune == '.' {
			continue
		} else if unicode.IsDigit(rune) && rune > wordHyphenationNumbers[matchIndx + currentWordGapIndx] {
			wordHyphenationNumbers[matchIndx + currentWordGapIndx] = rune
		} else {
			currentWordGapIndx += 1
		}
	}
}

func removeDigit(character rune) rune {
	if unicode.IsDigit(character) || character == '.' {
		return rune(-1)
	}
	return character
}
