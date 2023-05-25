package main

import "golang.org/x/exp/constraints"

func BinarySearch[T constraints.Ordered](a []T, x T) int {
	start, mid, end := 0, 0, len(a)-1
	for start <= end {
		mid = (start + end) >> 1
		switch {
		case a[mid] > x:
			end = mid - 1
		case a[mid] < x:
			start = mid + 1
		default:
			return mid
		}
	}
	return -1
}
