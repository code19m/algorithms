def find_smallest(arr: list) -> int:
    smallest_index = 0
    smallest = arr[smallest_index]
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest_index = i
    return smallest_index

def selection_sort(arr: list) -> list:
    new_arr = []
    for i in range(len(arr)):
        smallest = find_smallest(arr)
        new_arr.append(arr.pop(smallest))
    return new_arr
