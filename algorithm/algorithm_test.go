package algorithm

import (
	"testing"
	"io/ioutil"
	"strings"
)

var testData map[string]string = map[string]string{
	"mistranslate": "mis-trans-late",
	"alphabetical": "al-pha-bet-i-cal",
	"bewildering": "be-wil-der-ing",
	"buttons": "but-ton-s",
	"ceremony": "cer-e-mo-ny",
	"hovercraft": "hov-er-craft",
	"lexicographically": "lex-i-co-graph-i-cal-ly",
	"programmer": "pro-gram-mer",
	"recursion": "re-cur-sion",
}

func TestLoop(t *testing.T) {
	loopAlgorithm := NewLoopAlgorithm(loadPatterns())
	runTestsForAlgorithm(t, loopAlgorithm)
}

func TestTree(t *testing.T) {
	treeAlgorithm := NewTreeAlgorithm(loadPatterns())
	runTestsForAlgorithm(t, treeAlgorithm)
}

func loadPatterns() []string {
	patternsFileContent, err := ioutil.ReadFile("../tex-hyphenation-patterns.txt")
	if err != nil {
		panic("Cant load patterns")
	}
	return strings.Split(string(patternsFileContent), "\n")
}

func runTestsForAlgorithm(t *testing.T, algorithm HyphenationAlgorithm) {
	for input, expected := range testData {
		result := algorithm.HyphenateWord(input)
		if expected != result {
			t.Error("Expected " + expected + " got " + result)
		}
	}
}