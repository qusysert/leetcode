package main

type MinStack struct {
	Slice    []int
	MinSlice []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	if len(this.MinSlice) == 0 || val < this.MinSlice[len(this.MinSlice)-1] {
		this.MinSlice = append(this.MinSlice, val)
	} else {
		this.MinSlice = append(this.MinSlice, this.MinSlice[len(this.MinSlice)-1])
	}
	this.Slice = append(this.Slice, val)
}

func (this *MinStack) Pop() {
	if len(this.Slice) > 0 {
		this.Slice = this.Slice[:len(this.Slice)-1]
		this.MinSlice = this.MinSlice[:len(this.MinSlice)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.Slice) > 0 {
		return this.Slice[len(this.Slice)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.MinSlice) > 0 {
		return this.MinSlice[len(this.MinSlice)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
