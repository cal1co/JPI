package algo

import (
	"fmt"
	"time"
)

const (
	charSize = 26
)

type trieNode struct {
	children [charSize]*trieNode
	isEnd    bool
	pos      int
	length   int
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
	curr.length = len(word)
}

// func (t *trie) find(word string, edits int) bool {
// 	curr := t.root
// 	for i := 0; i < len(word); i++ {
// 		idx := word[i] - 'a'
// 		if curr.children[idx] == nil {
// 			return false
// 		}
// 		curr = curr.children[idx]
// 	}
// 	fmt.Println(curr.pos)

// 	return curr.isEnd
// }

func (t *trie) find(word string, node *trieNode, edits int, index int, changes int) []map[int]float64 {
	curr := t.root
	// run a bfs
	var slice []map[int]float64
	var bfs func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int)
	bfs = func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int) {
		// fmt.Println("BFS CALLED", word, index)
		if node.isEnd {
			score := float64((node.length - changes)) / float64(len(word))
			if score >= 0.6 {
				slice = append(slice, map[int]float64{node.pos: score})
			}
		}
		if index == len(word) || edits < 0 {
			return nil
		}
		idx := word[index] - 'a'
		if node.children[idx] == nil {
			edits -= 1
			// run the bfs here and exclude the real child
			// nextNum := word[i + 1] - 'a'
			for i := 0; i < len(node.children); i++ {
				if node.children[i] != nil {
					// fmt.Println(node.children[i])
					bfs(word, node.children[i], edits, index+1, changes+1)
				}
			}
		} else {
			bfs(word, node.children[idx], edits, index+1, changes)
		}

		// // a - runs 69, p - '', p - '', e - runs 58, l - runs 58
		// // appel will return the index of apple as well as a score of 0.6
		// // the score is calculated by dividing 3/5 (len(term) - edits made)/len(res)
		return nil
	}
	bfs(word, curr, edits, 0, 0)
	fmt.Println(word, slice)
	return slice
}

func Blah() [][]map[int]float64 {
	start := time.Now()
	var arr [][]map[int]float64
	trie := initTrie()
	words := []string{"ape", "app", "apple", "add", "aud"}
	for i := 0; i < len(words); i++ {
		trie.insert(words[i], i)
	}
	// match := []string{"ape", "app", "aid", "aud", "aplle"}
	match := []string{"app", "ape", "apple", "add", "aplle", "aud", "aid"}
	for i := 0; i < len(words); i++ {
		matched := trie.find(match[i], trie.root, 1, 0, 0)
		if len(matched) > 0 {
			arr = append(arr, matched)
		}
	}
	fmt.Println(time.Since(start))
	return arr
}
