// 选择排序, 纯粹练手
package main

import "fmt"

func main() {
	datalist := []int{12, 10, 15, 6, 22, 11, 4, 2, 10}
	fmt.Println(datalist)

	// 从大到小排序
	dataLen := len(datalist)
	for i := 0; i < dataLen; i++ {
		max := i
		for j := i + 1; j < dataLen; j++ {
			if datalist[j] > datalist[max] {
				max = j
			}
		}
		datalist[i], datalist[max] = datalist[max], datalist[i]
	}

	fmt.Println(datalist)
}
