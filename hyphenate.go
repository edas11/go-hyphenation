package main

import (
	"flag"
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

	isTreeAlgorithm := flag.Bool("tree", true, "Whether to use tree or loop algorithm")
	wordsFileName := flag.String("file", "", "Name of file that contains words to hyphenate")
	flag.Parse()

	var algorithmRunner algorithm.HyphenationAlgorithm
	if *isTreeAlgorithm {
		algorithmRunner = algorithm.NewTreeAlgorithm(hyphenationPatterns)
	} else {
		algorithmRunner = algorithm.NewLoopAlgorithm(hyphenationPatterns)
	}

	if *wordsFileName == "" {
		wordsToHyphenate := flag.Args()
		hyphenatedWords = make([]string, len(wordsToHyphenate))
		for i, singleWord := range wordsToHyphenate {
			hyphenatedWords[i] = algorithmRunner.HyphenateWord(singleWord)
		}
	} else {
		wordsFileContent, err := ioutil.ReadFile(*wordsFileName)
		if err != nil {
			fmt.Println("Couldnt read " + *wordsFileName)
			return
		}
		words := strings.Split(string(wordsFileContent), "\n")
		hyphenatedWords = make([]string, len(words))
		for i, singleWord := range words {
			hyphenatedWords[i] = algorithmRunner.HyphenateWord(singleWord)
		}
	}

	fmt.Println(strings.Join(hyphenatedWords, "\n"))
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}