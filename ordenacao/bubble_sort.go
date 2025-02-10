package ordenacao

func BubbleSort(arr []int, size int) {

	for i := 0; i < size-1; i++ {
		swapped := false

		for j := 0; j < size-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}
