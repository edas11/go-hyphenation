package algorithm

type HyphenationAlgorithm interface {
	HyphenateWord(wordToHyphenate string) string
}