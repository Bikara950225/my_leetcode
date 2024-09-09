package main

func circularArrayLoop(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	mp := map[int]bool{} // 记录走过的下标

	l := len(nums)
	for i := 0; i < l; i++ {
		ii := i
		loopLen := 0
		leftStep, rightStep := false, false
		for {
			isXX := leftStep && rightStep
			if mp[ii] && ii == i && loopLen > 1 && !isXX {
				return true
			} else if mp[ii] {
				mp = map[int]bool{}
				break
			}
			mp[ii] = true
			loopLen++

			if nums[ii] > 0 {
				leftStep = true
			} else {
				rightStep = true
			}
			ii += nums[ii]
			if ii >= l {
				ii %= l
			} else if ii < 0 {
				ii = l + (ii % l)
			}
		}
	}
	return false
}

func main() {

}
