# 冒泡排序
# 核心：冒泡，持续比较相邻元素，大的挪到后面，因此大的会逐步往后挪，故称之为冒泡。


def bubble(list):
    length = len(list)
    for x in range(length):
        for y in range(1, length-x):
            if list[y-1] > list[y]:
                list[y-1], list[y] = list[y], list[y-1]
    return list


if __name__ == '__main__':
    list = [10, 4, 8, 2, 5, 3, 9]
    print(list)
    print(bubble(list))
