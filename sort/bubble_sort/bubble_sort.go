// 冒泡排序
package main

import "fmt"

func main() {
	list := []int{10, 7, 3, 9, 2, 3, 1, 6}
	fmt.Println(list)

	sorted := false
	length := len(list)
	for i := 0; i < length && !sorted; i++ {
		min := i
		for j := i; j < length; j++ {
			if list[j] < list[min] {
				sorted = false
				min = j
			}
		}
		list[i], list[min] = list[min], list[i]
	}

	fmt.Println(list)
}
