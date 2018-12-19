#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 归并排序，纯粹练手


def mergeSort(data):
    if len(data) <= 1:
        return data

    mid = len(data) // 2
    left = mergeSort(data[:mid])
    right = mergeSort(data[mid:])

    return mergeSortArray(left, right)


def mergeSortArray(data1, data2):
    sortedData = []
    l = r = 0

    while len(data1) > l and len(data2) > r:
        if data1[l] < data2[r]:
            sortedData.append(data1[l])
            l += 1
        else:
            sortedData.append(data2[r])
            r += 1

    sortedData += data1[l:]
    sortedData += data2[r:]

    return sortedData


if __name__ == '__main__':
    data = [10, 3, 7, 1, 2, 5, 9, 11, 5]
    print(data)
    print(mergeSort(data))
