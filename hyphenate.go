package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
	"time"

	"github.com/edas11/hyphenation/algorithm"
)

type cli struct {
	isTreeAlgorithm bool
	isConcurrent    bool
	wordsFileName   string
	args            []string
}

var hyphenatedWords []string
var cliData cli
var algorithmRunner algorithm.HyphenationAlgorithm
var chanForWords chan []string

func main() {
	hyphenationPatterns := loadPatterns()
	cliData = parseCli()
	algorithmRunner = getAlgorithm(hyphenationPatterns)
	start := time.Now()
	hyphenatedWords := runAlgorithm()
	elapsed := time.Since(start)
	fmt.Println(strings.Join(hyphenatedWords, "\n"))
	fmt.Println(elapsed)
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
	isConcurrent := flag.Bool("concurrent", true, "Whether to run algorithm concurrently")
	wordsFileName := flag.String("file", "", "Name of file that contains words to hyphenate")
	flag.Parse()
	return cli{*isTreeAlgorithm, *isConcurrent, *wordsFileName, flag.Args()}
}

func getAlgorithm(hyphenationPatterns []string) algorithm.HyphenationAlgorithm {
	if cliData.isTreeAlgorithm {
		return algorithm.NewTreeAlgorithm(hyphenationPatterns)
	}
	return algorithm.NewLoopAlgorithm(hyphenationPatterns)
}

func runAlgorithm() []string {
	var words []string
	if cliData.wordsFileName == "" {
		words = cliData.args
	} else {
		wordsFileContent, err := ioutil.ReadFile(cliData.wordsFileName)
		if err != nil {
			panic("Couldnt read " + cliData.wordsFileName)
		}
		words = strings.Split(string(wordsFileContent), "\n")
	}

	if cliData.isConcurrent {
		return runConcurrentlyAlgorithmOnWords(words)
	}
	return runAlgorithmOnWords(words)
}

func runConcurrentlyAlgorithmOnWords(words []string) []string {
	hyphenatedWords = make([]string, len(words))
	numCPU := runtime.NumCPU()
	numOfWordsForOneCPU := len(words) / numCPU
	chanForWords = make(chan []string)
	for i := 0; i < numCPU; i++ {
		if i == numCPU-1 {
			go runAlgorithmOnWords(words[i*numOfWordsForOneCPU:])
		} else {
			go runAlgorithmOnWords(words[i*numOfWordsForOneCPU : (i+1)*numOfWordsForOneCPU])
		}
	}
	emptyFrom := 0
	for i := 0; i < numCPU; i++ {
		blockOfHyphenatedWords := <-chanForWords
		copy(hyphenatedWords[emptyFrom:emptyFrom+len(blockOfHyphenatedWords)], blockOfHyphenatedWords)
		emptyFrom += len(blockOfHyphenatedWords)
	}
	return hyphenatedWords
}

func runAlgorithmOnWords(words []string) []string {
	hyphenatedWords := make([]string, len(words))
	for i, singleWord := range words {
		hyphenatedWords[i] = algorithmRunner.HyphenateWord(singleWord)
	}
	if chanForWords == nil {
		return hyphenatedWords
	}
	chanForWords <- hyphenatedWords
	return nil
}
