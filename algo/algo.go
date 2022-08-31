package algo

import "fmt"

const (
	charSize = 26
)

type trieNode struct {
	children [charSize]*trieNode
	isEnd    bool
	pos      int
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word string, index int) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &trieNode{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
	curr.pos = index
}

func (t *trie) find(word string, edits int) bool {
	curr := t.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	fmt.Println(curr.pos)

	// run a bfs
	return curr.isEnd
}

// bfs := func(word string, node *trieNode, edits int, index int) []string {
// 	if index == len(word) || edits < 0 {
// 		return
// 	}
// 	idx := word[i] - 'a'
// 	if node.children[idx] == nil {
// 		edits -= 1
// 		// run the bfs here and exclude the real child
// 		nextNum := word[i + 1] - 'a'
// 		for i := 0; i < len(node.children); i++ {
// 			if node.children[i] != nil && node.children[i] != node.children[nextNum] {
// 				bfs(word, node.children[i], edits, index + 1)
// 			}
// 		}
// 	}

// 	bfs(word, node.children[idx], edits, index + 1)

// 	if node.isEnd {
// 		return node.pos
// 	}
// }(word, curr, edits, index)

func Blah() []string {
	arr := []string{}
	trie := initTrie()
	words := []string{"ape", "app", "apple", "add", "aud"}
	for i := 0; i < len(words); i++ {
		trie.insert(words[i], i)
	}
	// match := []string{"ape", "app", "aid", "aud", "aplle"}
	match := []string{"app", "ape", "apple", "add", "aplle", "aud"}
	// match := []string{"app", "ape", "aid", "apple", "add"}
	for i := 0; i < len(words); i++ {
		matched := trie.find(match[i], 2)
		if matched {
			arr = append(arr, match[i])
		}
	}
	return arr
}
