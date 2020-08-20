package leetcode

func maxArea(hights []int) int {
	max, left, right := 0, 0, len(hights)-1
	for left < right {
		width := right - left
		hight := 0

		if hights[left] < hights[right] {
			hight = hights[left]
			left++
		} else {
			hight = hights[right]
			right--
		}

		if tmp := hight * width; tmp > max {
			max = tmp
		}
	}

	return max
}
