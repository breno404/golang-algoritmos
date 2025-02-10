package ordenacao

func SelectionSort(arr []int, size int) {
	for i := 0; i < size; i++ {
		min := i

		for j := i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
}
