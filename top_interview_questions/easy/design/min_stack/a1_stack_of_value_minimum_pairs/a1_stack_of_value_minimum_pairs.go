package a1_stack_of_value_minimum_pairs

type StackItem struct {
	Val int
	Min int
}

type MinStack struct {
	container []StackItem
}

func Constructor() MinStack {
	container := []StackItem{}
	return MinStack{container: container}
}

func (this *MinStack) Push(val int) {
	newItem := StackItem{Val: val, Min: val}

	if len(this.container) > 0 {
		topItem := this.container[len(this.container)-1]
		if topItem.Min < newItem.Min {
			newItem.Min = topItem.Min
		}
	}

	this.container = append(this.container, newItem)
}

func (this *MinStack) Pop() {
	this.container = this.container[:len(this.container)-1]
}

func (this *MinStack) Top() int {
	return this.container[len(this.container)-1].Val
}

func (this *MinStack) GetMin() int {
	return this.container[len(this.container)-1].Min
}
