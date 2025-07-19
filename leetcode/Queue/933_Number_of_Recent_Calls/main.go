package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(642))
	fmt.Println(obj.Ping(1849))
	fmt.Println(obj.Ping(4921))
	fmt.Println(obj.Ping(5936))
	fmt.Println(obj.Ping(5957))
}

type RecentCounter struct {
	history []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	cnt := 0
	this.history = append(this.history, t)
	for i := range len(this.history) {
		if this.history[i] < t-3000 {
			cnt++
			continue
		}
		this.history = this.history[cnt+1:]
	}
	return len(this.history)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
