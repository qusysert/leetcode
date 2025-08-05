package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := NewTrie()

	for i := range products {
		trie.Insert(products[i])
	}

	prefix := ""
	var res [][]string

	for char := range searchWord {
		prefix += string(searchWord[char])
		cur := trie.Search(prefix).Words
		sort.Strings(cur)
		if len(cur) > 3 {
			res = append(res, cur[:3])
		} else {
			res = append(res, cur)
		}

	}

	return res
}

type Trie struct {
	head *Node
}

type Node struct {
	children map[rune]*Node
	isEnd    bool
	Words    []string
}

func NewTrie() Trie {
	m := make(map[rune]*Node)
	return Trie{&Node{children: m}}
}

func (this *Trie) Insert(word string) {
	cur := this.head
	for _, c := range word {
		cur.Words = append(cur.Words, word)
		_, ok := cur.children[c]
		if !ok {
			m := make(map[rune]*Node)
			cur.children[c] = &Node{children: m, isEnd: false}
		}
		cur = cur.children[c]
	}
	cur.Words = append(cur.Words, word)
	cur.isEnd = true
}

func (this *Trie) Search(word string) *Node {
	cur := this.head

	for _, c := range word {
		new, ok := cur.children[c]
		if !ok {
			return &Node{}
		}
		cur = new
	}
	return cur
}
