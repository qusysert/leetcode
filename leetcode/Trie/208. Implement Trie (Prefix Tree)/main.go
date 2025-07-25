package main

type Trie struct {
	head *Node
}

type Node struct {
	children map[rune]*Node
	isEnd    bool
}

func Constructor() Trie {
	m := make(map[rune]*Node)
	return Trie{&Node{children: m}}
}

func (this *Trie) Insert(word string) {
	cur := this.head
	for _, c := range word {
		_, ok := cur.children[c]
		if !ok {
			m := make(map[rune]*Node)
			cur.children[c] = &Node{children: m, isEnd: false}
		}
		cur = cur.children[c]
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.head

	for _, c := range word {
		new, ok := cur.children[c]
		if !ok {
			return false
		}
		cur = new
	}
	if cur.isEnd == false {
		return false
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.head

	for _, c := range prefix {
		newNode, ok := cur.children[c]
		if !ok {
			return false
		}
		cur = newNode
	}
	return true
}
