#!/usr/bin/env python
# -*- coding: utf-8 -*-
#
# 快速排序，纯粹练手


class QuickSort(object):
    def __init__(self, data):
        self.__data = data

    def run(self):
        self.__run_helper(0, len(self.__data) - 1)

    def __run_helper(self, l, r):
        if l >= r:
            return

        p = self.__partision(l, r)
        self.__run_helper(l, p - 1)
        self.__run_helper(p + 1, r)

    def __partision(self, l, r):
        ti, t = l, self.__data[l]

        l += 1
        while l <= r:
            while self.__data[l] < t:
                l += 1
            while self.__data[r] >= t:
                r -= 1

            if l > r:
                break

            self.__data[l], self.__data[r] = self.__data[r], self.__data[l]

        self.__data[r], self.__data[ti] = self.__data[ti], self.__data[r]
        return r


if __name__ == '__main__':
    data = [3, 2, 111, 3, -1, 0, 0, 1, 0, 2, 4]
    print(data)
    sort = QuickSort(data)
    sort.run()
    print(data)
