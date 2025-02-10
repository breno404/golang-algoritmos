package ordenacao

func QuickSort(arr []int, left int, right int) {

	if left >= right {
		return
	}

	pivo := arr[(left+right)/2]
	i, j := left, right

	for i <= j {
		for arr[i] < pivo {
			i++
		}

		for arr[j] > pivo {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	if left < right {
		QuickSort(arr, left, j)
		QuickSort(arr, i, right)
	}

}
