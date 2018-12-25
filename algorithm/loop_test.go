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

func TestHyphentation(t *testing.T) {
	patternsFileContent, err := ioutil.ReadFile("../tex-hyphenation-patterns.txt")
	if err != nil {
		return
	}
	hyphenationPatterns := strings.Split(string(patternsFileContent), "\n")

	loopAlgorithm := NewLoopAlgorithm(hyphenationPatterns)
	for input, expected := range testData {
		result := loopAlgorithm.HyphenateWord(input)
		if expected != result {
			t.Error("Expected " + expected + " got " + result)
		}
	}
}