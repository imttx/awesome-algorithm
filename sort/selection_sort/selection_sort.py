#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 选择排序, 纯粹练手


def selection_sort(list):
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
