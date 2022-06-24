package sorting

func main() {

}

func BubbleSort(nums []int) []int { // TODO: Fix me.
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}

	return nums
}
