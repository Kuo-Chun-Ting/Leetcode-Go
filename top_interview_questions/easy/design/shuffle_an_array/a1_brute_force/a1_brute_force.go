package a1_brute_force

import (
	"math/rand"
)

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
	hat := make([]int, len(this.initNums))
	copy(hat, this.initNums)

	for i := 0; i < len(this.currNums); i++ {
		random := rand.Intn(len(hat))
		this.currNums[i] = hat[random]
		hat = this.remove(random, hat)
	}

	return this.currNums
}

func (this *Solution) remove(index int, nums []int) []int {
	return append(nums[:index], nums[index+1:]...)
}
