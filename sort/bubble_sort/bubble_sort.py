# 冒泡排序


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
