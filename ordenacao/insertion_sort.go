package ordenacao

func InsertionSort(arr []int, size int) {

	for i := 1; i < size; i++ {

		temp := arr[i]

		j := i - 1

		for j >= 0 && temp < arr[j] {

			arr[j+1] = arr[j]

			j--
		}

		arr[j+1] = temp
	}
}
