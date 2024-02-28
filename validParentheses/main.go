package main

func main() {

}

func isValid(s string) bool {
	dummy := &Node{}
	stackHead := dummy
	parenMap := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, char := range s {
		newest := &Node{Char: char, Prev: stackHead}
		stackHead = newest
		if isClosingParen(newest.Char) {
			if parenMap[newest.Prev.Char] != newest.Char {
				return false
			} else {
				stackHead = stackHead.Prev.Prev
			}
		}
	}
	if stackHead != dummy {
		return false
	}
	return true
}

func isClosingParen(char rune) bool {
	return char == ')' || char == '}' || char == ']'
}

type Node struct {
	Char rune
	Prev *Node
}
