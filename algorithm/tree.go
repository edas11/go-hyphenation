package algorithm

import(
	"unicode"
	"strings"
)

type TreeAlgorithm struct {
	tree patternsTree
}

type patternsTree struct {
	nextTreeBranches map[rune]*patternsTree
	patternsOfThisLevel []string
}

func NewTreeAlgorithm(hyphenationPatterns []string) TreeAlgorithm {
	tree := patternsTree{make(map[rune]*patternsTree), make([]string, 1)}
	for _, pattern := range hyphenationPatterns {
		tree.putPatternIntoTree(pattern)
	}
	return TreeAlgorithm{tree}
}

func (algorithm TreeAlgorithm) HyphenateWord(wordToHyphenate string) string {
	word := newWord(wordToHyphenate)
	for i := range word.wordToHyphenate {
		word.matchAllPatternsAt(i, &algorithm.tree)
	}
	return word.generateHyphenatedWord()
}

func (word word) matchAllPatternsAt(matchIndx int, tree *patternsTree) {
	wordToHyphenate := []rune(word.wordToHyphenate)
	wordLength := len(wordToHyphenate)
	currentCharIndx := matchIndx
	currentLevelTree := tree
	var ok bool
	for {
		for _, pattern := range currentLevelTree.patternsOfThisLevel {
			if (strings.HasPrefix(pattern, ".") && matchIndx != 0) {
				continue
			}
			patternLength := currentCharIndx - matchIndx
			if (strings.HasSuffix(pattern, ".") && matchIndx != wordLength - patternLength) {
				continue
			}
			word.updateWordHyphenationNumbers(pattern, matchIndx)
		}
		if currentCharIndx >= wordLength || currentLevelTree == nil {
			break;
		}
		currentLevelTree, ok = currentLevelTree.nextTreeBranches[wordToHyphenate[currentCharIndx]]
		if !ok {
			break
		}
		currentCharIndx++
	}
}

func (tree *patternsTree) putPatternIntoTree (pattern string) {
	currentTreeLevel := tree
	for _, patternChar := range pattern {
		if patternChar == '.' || unicode.IsDigit(patternChar) {
			continue
		}

		_, ok := currentTreeLevel.nextTreeBranches[patternChar]
		if !ok {
			currentTreeLevel.nextTreeBranches[patternChar] = &patternsTree{make(map[rune]*patternsTree), make([]string, 1)}
		}
		currentTreeLevel = currentTreeLevel.nextTreeBranches[patternChar]
	}

	if currentTreeLevel.patternsOfThisLevel == nil {
		currentTreeLevel.patternsOfThisLevel = make([]string, 1)
	}
	currentTreeLevel.patternsOfThisLevel = append(currentTreeLevel.patternsOfThisLevel, pattern)
}