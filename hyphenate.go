package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
	"time"
)

var wordToHyphenate string
var wordHyphenationNumbers []rune
var hyphenatedWords []string

func main() {
	start := time.Now()

	patternsFileContent, err := ioutil.ReadFile("tex-hyphenation-patterns.txt")
	if err != nil {
		return
	}
	hyphenationPatterns := strings.Split(string(patternsFileContent), "\n")

	cliArguments := os.Args
	if len(cliArguments) >= 2 {
		hyphenatedWords = make([]string, len(cliArguments) - 1)
		for i, singleWord := range os.Args[1:] {
			wordToHyphenate = singleWord
			hyphenatedWords[i] = hyphenateWord(hyphenationPatterns)
		}
	} else if len(cliArguments) == 1 {
		wordsFileContent, err := ioutil.ReadFile("words.txt")
		if err != nil {
			return
		}
		words := strings.Split(string(wordsFileContent), "\n")
		hyphenatedWords = make([]string, len(words))
		for i, singleWord := range words {
			wordToHyphenate = singleWord
			hyphenatedWords[i] = hyphenateWord(hyphenationPatterns)
		}
	}
	fmt.Println(strings.Join(hyphenatedWords, "\n"))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func hyphenateWord(hyphenationPatterns []string) string {
	wordHyphenationNumbers = make([]rune, len(wordToHyphenate) + 1)
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
			matchIndex := indexFrom(wordToHyphenate, reducedPattern, 0)
			for matchIndex != -1 {
				updateWordHyphenationNumbers(pattern, matchIndex)
				matchIndex = indexFrom(wordToHyphenate, reducedPattern, matchIndex + 1)
			}
		}
	}
	return generateHyphenatedWord()
}

func indexFrom(str, search string, from int) int{
	if from > len(str) - 1 {
		return -1
	}
	matchInSubstring := strings.Index(str[from:], search)
	if matchInSubstring == -1 {
		return -1
	}
	return matchInSubstring + from
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
		} else if unicode.IsDigit(rune) {
			if rune > wordHyphenationNumbers[matchIndx + currentWordGapIndx] {
				wordHyphenationNumbers[matchIndx + currentWordGapIndx] = rune
			}
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
