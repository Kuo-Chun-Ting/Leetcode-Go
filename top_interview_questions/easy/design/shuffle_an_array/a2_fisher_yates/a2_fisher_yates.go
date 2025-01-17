package a2_fisher_yates

import "math/rand"

type Solution struct {
	initNums []int
	currNums []int
}

func Constructor(nums []int) Solution {
	initNums := make([]int, len(nums))
	copy(initNums, nums)

	return Solution{initNums: initNums, currNums: nums}
}

func (this *Solution) Reset() []int {
	copy(this.currNums, this.initNums)
	return this.currNums
}

func (this *Solution) Shuffle() []int {
	for i := 0; i < len(this.currNums); i++ {
		random := this.randomInRange(i, len(this.currNums)-1)
		this.currNums[i], this.currNums[random] = this.currNums[random], this.currNums[i]
	}

	return this.currNums
}

func (this *Solution) randomInRange(min, max int) int {
	if min > max {
		panic("invalid range: min is greater than max")
	}
	return rand.Intn(max-min+1) + min
}
