#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date    : 2018-12-22 11:16:42
#
# 堆排序 纯粹练手


class HeapSort(object):
    def __init__(self, data):
        self.data = data

    def sort(self):
        # 建立最大堆
        for i in range(len(self.data)//2-1, -1, -1):
            # 从第一个非叶子结点从下至上，从右至左调整结构
            self.sink(i, len(self.data))

        # 调整下沉元素
        for i in range(len(self.data)-1, 0, -1):
            # 堆顶元素移到最后一位
            self.swap(0, i)

            # 重新调整堆结构
            self.sink(0, i)

    # 下沉元素
    def sink(self, i, length):
        tmp = self.data[i]
        k = 2 * i + 1
        while k < length:
            if k + 1 < length and self.data[k] < self.data[k+1]:
                k += 1

            if self.data[k] > tmp:
                self.data[i] = self.data[k]
                i = k
                k = 2*k+1
            else:
                break

        self.data[i] = tmp

    def swap(self, i, j):
        self.data[i], self.data[j] = self.data[j], self.data[i]

    def display(self):
        print(self.data)


if __name__ == '__main__':
    data = [10, 3, 4, 8, 7, 20, 1, 9, 2]
    heap = HeapSort(data)
    heap.display()
    heap.sort()
    heap.display()
