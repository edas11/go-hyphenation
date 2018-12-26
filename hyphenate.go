package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"github.com/edas11/hyphenation/algorithm"
)

type cli struct {
	isTreeAlgorithm bool
	wordsFileName string
	args []string
}

var hyphenatedWords []string
var cliData cli

func main() {
	start := time.Now()
	hyphenationPatterns := loadPatterns()
	cliData = parseCli()
	algorithmRunner := getAlgorithm(hyphenationPatterns)
	hyphenatedWords := runAlgorithm(algorithmRunner)
	fmt.Println(strings.Join(hyphenatedWords, "\n"))
	fmt.Println(time.Since(start))
}

func loadPatterns() []string {
	patternsFileContent, err := ioutil.ReadFile("tex-hyphenation-patterns.txt")
	if err != nil {
		panic("Cant load patterns")
	}
	return strings.Split(string(patternsFileContent), "\n")
}

func parseCli() cli {
	isTreeAlgorithm := flag.Bool("tree", true, "Whether to use tree or loop algorithm")
	wordsFileName := flag.String("file", "", "Name of file that contains words to hyphenate")
	flag.Parse()
	return cli{*isTreeAlgorithm, *wordsFileName, flag.Args()}
}

func getAlgorithm(hyphenationPatterns []string) algorithm.HyphenationAlgorithm {
	if cliData.isTreeAlgorithm {
		return algorithm.NewTreeAlgorithm(hyphenationPatterns)
	}
	return algorithm.NewLoopAlgorithm(hyphenationPatterns)
}

func runAlgorithm(algorithmRunner algorithm.HyphenationAlgorithm) []string {
	if cliData.wordsFileName == "" {
		return runAlgorithmOnWordsFromCli(algorithmRunner)
	}
	return runAlgorithmOnWordsFromFile(algorithmRunner)
}

func runAlgorithmOnWordsFromCli(algorithmRunner algorithm.HyphenationAlgorithm) []string {
	wordsToHyphenate := cliData.args
	hyphenatedWords = make([]string, len(wordsToHyphenate))
	for i, singleWord := range wordsToHyphenate {
		hyphenatedWords[i] = algorithmRunner.HyphenateWord(singleWord)
	}
	return hyphenatedWords
}

func runAlgorithmOnWordsFromFile(algorithmRunner algorithm.HyphenationAlgorithm) []string {
	wordsFileContent, err := ioutil.ReadFile(cliData.wordsFileName)
	if err != nil {
		panic("Couldnt read " + cliData.wordsFileName)
	}
	words := strings.Split(string(wordsFileContent), "\n")
	hyphenatedWords = make([]string, len(words))
	for i, singleWord := range words {
		hyphenatedWords[i] = algorithmRunner.HyphenateWord(singleWord)
	}
	return hyphenatedWords
}