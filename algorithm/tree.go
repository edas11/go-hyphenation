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
	currentCharIndx := matchIndx
	currentLevelTree := tree
	var hasMore bool
	for {
		word.processAllCurrenlyMatchedPatterns(currentLevelTree.patternsOfThisLevel, currentCharIndx - matchIndx, matchIndx)
		
		cantGoFurther := currentCharIndx >= word.wordLength || currentLevelTree == nil
		if cantGoFurther {
			break;
		}
		currentLevelTree, hasMore = currentLevelTree.nextTreeBranches[wordToHyphenate[currentCharIndx]]
		if !hasMore {
			break
		}
		currentCharIndx++
	}
}

func (word word) processAllCurrenlyMatchedPatterns(matchedPatterns []string, patternsLength int, matchIndx int) {
	for _, pattern := range matchedPatterns {
		if (strings.HasPrefix(pattern, ".") && matchIndx != 0) {
			continue
		}
		if (strings.HasSuffix(pattern, ".") && matchIndx != word.wordLength - patternsLength) {
			continue
		}
		word.updateWordHyphenationNumbers(pattern, matchIndx)
	}
}

func (tree *patternsTree) putPatternIntoTree (pattern string) {
	treeLevelForPattern := tree.findOrCreateTreeLevelForPattern(pattern)

	if treeLevelForPattern.patternsOfThisLevel == nil {
		treeLevelForPattern.patternsOfThisLevel = make([]string, 1)
	}
	treeLevelForPattern.patternsOfThisLevel = append(treeLevelForPattern.patternsOfThisLevel, pattern)
}

func (tree *patternsTree) findOrCreateTreeLevelForPattern (pattern string) *patternsTree {
	currentTreeLevel := tree
	var ok bool
	for _, patternChar := range pattern {
		if !unicode.IsLetter(patternChar) {
			continue
		}

		_, ok = currentTreeLevel.nextTreeBranches[patternChar]
		if !ok {
			currentTreeLevel.nextTreeBranches[patternChar] = &patternsTree{make(map[rune]*patternsTree), make([]string, 1)}
		}

		currentTreeLevel = currentTreeLevel.nextTreeBranches[patternChar]
	}
	return currentTreeLevel
}