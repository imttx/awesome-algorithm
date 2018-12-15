# 冒泡排序


def bubble(list):
    length = len(list)
    sorted = False
    for x in range(length):
        if not sorted:
            break

        min = x
        for y in range(x, length):
            if list[y] < list[min]:
                sorted = False
                min = y
        list[x], list[min] = list[min], list[x]

    return list


if __name__ == '__main__':
    list = [10, 4, 8, 2, 5, 3, 9]
    print(list)
    print(bubble(list))
