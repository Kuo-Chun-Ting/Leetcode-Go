package a2_two_stacks

type MinStack struct {
	container    []int
	minContainer []int
}

func Constructor() MinStack {
	container := []int{}
	minContainer := []int{}
	return MinStack{container: container, minContainer: minContainer}
}

func (this *MinStack) Push(val int) {
	this.container = append(this.container, val)

	if len(this.minContainer) == 0 {
		this.minContainer = append(this.minContainer, val)
	} else {
		currMin := this.minContainer[len(this.minContainer)-1]
		if val <= currMin {
			this.minContainer = append(this.minContainer, val)
		}
	}
}

func (this *MinStack) Pop() {
	topVal := this.Top()
	this.container = this.container[:len(this.container)-1]

	if len(this.minContainer) > 0 {
		currMin := this.minContainer[len(this.minContainer)-1]
		if currMin == topVal {
			this.minContainer = this.minContainer[:len(this.minContainer)-1]
		}
	}
}

func (this *MinStack) Top() int {
	return this.container[len(this.container)-1]
}

func (this *MinStack) GetMin() int {
	return this.minContainer[len(this.minContainer)-1]
}
