package main

func main() {
	nums := []int{1, 3, 2, 1, 3, 2, 2}

	println(numberOfPairs(nums))
}
func numberOfPairs(nums []int) []int {
	m := make(map[int]int)
	for i := range nums {
		m[i]++
	}
	ans := []int{0, 0}
	for _, v := range m {

		if v%2 == 0 {
			//当前为偶数对

			ans[0] += v / 2
		} else {
			ans[0] += v / 2
			ans[1]++
		}
	}
	return ans
}
