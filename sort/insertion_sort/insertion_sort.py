#!/usr/bin/env python
# -*- coding: utf-8 -*-

# 插入排序


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
