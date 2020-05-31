import "fmt"

func main() {
	nums := []int{1,3,-1,-3,5,3,6,7}
	fmt.Println(maxSlidingWindow(nums, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}

	var ret []int
	for i := 0; i < l - k + 1; i++ {
		max := nums[i]
		j := i + 1
		for  j <= i + k -1 {
			max = checkTwoNum(max, nums[j])
			j++
		}
		ret = append(ret, max)
	}

	return ret
}

func checkTwoNum(a, b int) int {
	if a > b {
		return a
	}

	return b
}
