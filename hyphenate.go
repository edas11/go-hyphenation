package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"github.com/edas11/hyphenation/algorithm"
)

var hyphenatedWords []string

func main() {
	start := time.Now()

	patternsFileContent, err := ioutil.ReadFile("tex-hyphenation-patterns.txt")
	if err != nil {
		return
	}
	hyphenationPatterns := strings.Split(string(patternsFileContent), "\n")

	loopAlgorithm := algorithm.NewLoopAlgorithm(hyphenationPatterns)
	cliArguments := os.Args
	if len(cliArguments) >= 2 {
		hyphenatedWords = make([]string, len(cliArguments) - 1)
		for i, singleWord := range os.Args[1:] {
			hyphenatedWords[i] = loopAlgorithm.HyphenateWord(singleWord)
		}
	} else if len(cliArguments) == 1 {
		wordsFileContent, err := ioutil.ReadFile("words.txt")
		if err != nil {
			return
		}
		words := strings.Split(string(wordsFileContent), "\n")
		hyphenatedWords = make([]string, len(words))
		for i, singleWord := range words {
			hyphenatedWords[i] = loopAlgorithm.HyphenateWord(singleWord)
		}
	}
	fmt.Println(strings.Join(hyphenatedWords, "\n"))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}