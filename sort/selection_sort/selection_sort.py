#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 选择排序, 纯粹练手
# 核心：不断地选择剩余元素中的最小者。
# 1.找到数组中最小元素并将其和数组第一个元素交换位置。
# 2.在剩下的元素中找到最小元素并将其与数组第二个元素交换，直至整个数组排序。


def selection_sort(list):
    # 从小到大排序
    for i in range(len(list)):
        min = i
        for j in range(i + 1, len(list)):
            if list[j] < list[min]:
                min = j
        list[min], list[i] = list[i], list[min]

    return list


if __name__ == '__main__':
    list = [12, 10, 15, 6, 22, 11, 4, 2, 10]
    print("排序前", list)
    selection_sort(list)
    print("排序后", list)
