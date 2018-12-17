// 冒泡排序
package main

import "fmt"

func main() {
	list := []int{10, 7, 3, 9, 2, 3, 1, 6}
	fmt.Println(list)

	length := len(list)
	for i := 0; i < length; i++ {
		for j := 1; j < length-i; j++ {
			if list[j-1] > list[j] {
				list[j-1], list[j] = list[j], list[j-1]
			}
		}
	}

	fmt.Println(list)
}
