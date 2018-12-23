#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 插入排序

'''
核心：通过构建有序序列，对于未排序序列，在已排序序列中从后向前扫描(对于单向链表则只能从前往后遍历)，找到相应位置并插入。实现上通常使用in-place排序(需用到O(1)的额外空间)

	1. 从第一个元素开始，该元素可认为已排序
	2. 取下一个元素，对已排序数组从后往前扫描
	3. 若从排序数组中取出的元素大于新元素，则移至下一位置
	4. 重复步骤3，直至找到已排序元素小于或等于新元素的位置
	5. 插入新元素至该位置
	6. 重复2~5
'''


def insertionSort(data):
    for k, v in enumerate(data):
        idx = k
        while idx > 0 and data[idx-1] > v:
            data[idx] = data[idx-1]
            idx -= 1
        data[idx] = v

    return data


if __name__ == '__main__':
    data = [10, 4, 6, 2, 1, 8]
    print(data)
    print(insertionSort(data))
