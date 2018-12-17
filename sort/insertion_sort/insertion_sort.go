// 插入排序

package main

import "fmt"

func main() {
	data := []int{10, 4, 6, 2, 1, 8}
	fmt.Println(data)

	for k, v := range data {
		idx := k
		for idx > 0 && data[idx-1] > v {
			data[idx] = data[idx-1]
			idx--
		}
		data[idx] = v
	}

	fmt.Println(data)
}
