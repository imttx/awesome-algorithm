#!/usr/bin/env python
# -*- coding: utf-8 -*-
# @Date    : 2018-12-21 09:10:07
# 堆


class HeapSort(object):
    def __init__(self, num):
        self.data = [0 for _ in range(num + 1)]
        self.num = num

    def display(self):
        print(self.data)

    def size(self):
        return self.num

    def is_empty(self):
        return self.size() == 0

    def swap(self, i, j):
        self.data[i], self.data[j] = self[j], self.data[i]

    def less(self, i, j):
        return self.data[i] < self.data[j]

    # 上浮
    def swim(self, k):
        while k > 1 and self.less(k / 2, k):
            self.swap(k / 2, k)
            k /= 2

    # 下沉
    def sink(self, k):
        while k * 2 <= self.size():
            j = 2 * k
            if j + 1 < self.size() and self.less(j, j + 1):
                j += 1

            if not self.less(k, j):
                break

            self.swap(k, j)
            k = j
    # 插入一个元素

    def insert(self, n):
        self.num += 1
        self.data.append(n)
        self.swim(1)

    # 删除堆顶部元素
    def del_max(self):
        max = self.data[1]
        self.data.pop()
        self.num -= 1
        self.swap(1, self.num)
        self.sink(1)
        return max


if __name__ == '__main__':
    h = HeapSort(10)
    h.insert(4)
    h.insert(2)
    h.insert(6)
    h.insert(10)
    h.insert(3)
    h.display()
