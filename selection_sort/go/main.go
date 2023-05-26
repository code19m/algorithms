package selection_sort

func Selection_Sort(array []int, size int) []int {
	var min_index int
	var temp int
	for i := 0; i < size-1; i++ {
		min_index = i
		// Find index of minimum element
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	return array
}
