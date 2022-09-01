package algo

import "sort"

const (
	charSize = 26
)

type trieNode struct {
	children [charSize]*trieNode
	isEnd    bool
	pos      []int
	length   int
}

type Trie struct {
	Root *trieNode
}

type result struct {
	Position []int
	Score    float64
}

func InitTrie() *Trie {
	return &Trie{
		Root: &trieNode{},
	}
}

func Insert(word string, index int, t *Trie) {
	curr := t.Root
	for i := 0; i < len(word); i++ {
		if word[i] < 97 || word[i] > 123 {
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
	sort.Slice(curr.pos, func(i, j int) bool {
		return curr.pos[i] > curr.pos[j]
	})
	curr.length = len(word)
}

func Find(word string, node *Trie, edits int, index int, changes int) []result {
	var slice []result

	var bfs func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int)

	bfs = func(word string, node *trieNode, edits int, index int, changes int) func(map[int]int) {
		if index == len(word) && edits > 0 {
			index = len(word) - 1
		}
		if edits < 0 {
			return nil
		}
		if node.isEnd {
			var scr float64
			scr = float64((node.length - changes)) / float64(len(word))
			if node.length > len(word) {
				scr = float64((len(word) - changes)) / float64(len(word))
			}
			if scr >= 0.6 && scr <= 1.0 {
				slice = append(slice, result{Position: node.pos, Score: scr})
			}
		}
		if index == len(word) {
			return nil
		}

		idx := word[index] - 'a'
		for i := 0; i < len(node.children); i++ {
			if node.children[i] != nil {
				if node.children[i] != node.children[idx] {
					bfs(word, node.children[i], edits-1, index+1, changes+1)
				} else {
					bfs(word, node.children[idx], edits, index+1, changes)
				}
			}
		}
		return nil
	}
	bfs(word, node.Root, edits, 0, 0)

	return slice
}
