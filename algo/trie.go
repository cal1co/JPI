package algo

import (
	"fmt"
)

const (
	charSize = 26
)

type trieNode struct {
	children [charSize]*trieNode
	isEnd    bool
	pos      []int
	length   int
}

type trie struct {
	Root *trieNode
}

type result struct {
	Position []int
	score    float64
}

func InitTrie() *trie {
	return &trie{
		Root: &trieNode{},
	}
}

func Insert(word string, index int, t *trie) {
	// if word == "" {
	// 	return
	// }
	curr := t.Root
	for i := 0; i < len(word); i++ {
		if word[i] < 97 || word[i] > 123 {
			// break
			return
		}
		idx := word[i] - 'a'
		if curr.children[idx] == nil {
			curr.children[idx] = &trieNode{}
		}
		curr = curr.children[idx]
	}
	curr.isEnd = true
	curr.pos = append(curr.pos, index)
	curr.length = len(word)
}

func Find(word string, node *trie, edits int, index int, changes int) []result {
	// curr := t.Root
	var slice []result
	var bfs func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int)
	bfs = func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int) {
		if node.isEnd {
			scr := float64((node.length - changes)) / float64(len(word))
			if scr >= 0.6 && scr <= 1.0 {
				fmt.Println(node.pos)
				slice = append(slice, result{Position: node.pos, score: scr})
			}
		}
		if index == len(word) || edits < 0 {
			return nil
		}
		idx := word[index] - 'a'
		if node.children[idx] == nil {
			edits -= 1
			for i := 0; i < len(node.children); i++ {
				if node.children[i] != nil {
					bfs(word, node.children[i], edits, index+1, changes+1)
				}
			}
		} else {
			bfs(word, node.children[idx], edits, index+1, changes)
		}
		return nil
	}
	bfs(word, node.Root, edits, 0, 0)
	return slice
}

// func Blah() [][]map[int]float64 {
// 	start := time.Now()
// 	var arr [][]map[int]float64
// 	trie := InitTrie()
// 	words := []string{"ape", "app", "apple", "add", "aud"}
// 	for i := 0; i < len(words); i++ {
// 		trie.insert(words[i], i)
// 	}
// 	// match := []string{"ape", "app", "aid", "aud", "aplle"}
// 	match := []string{"app", "ape", "apple", "add", "aplle", "aud", "aid"}
// 	for i := 0; i < len(words); i++ {
// 		matched := trie.find(match[i], trie.root, 2, 0, 0)
// 		if len(matched) > 0 {
// 			arr = append(arr, matched)
// 		}
// 	}
// 	fmt.Println(time.Since(start))
// 	return arr
// }
