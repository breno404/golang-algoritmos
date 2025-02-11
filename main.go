package main

import (
	"algoritmos/inversao"
	"algoritmos/ordenacao"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	command := strings.Join(args, " ")

	arr := []int{1, 25, 3, 8, 32, 6, 7, 27, 9, 10}

	if command == "inverter array" {
		fmt.Printf("Array original: %v\n", arr)

		for i := 0; i < len(arr)/2; i++ {
			last := 1 + i
			inversao.Exec(&arr[i], &arr[len(arr)-last])
		}

		fmt.Printf("Array invertido: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n)")
		fmt.Println("Complexidade no caso medio O(n)")
		fmt.Println("Complexidade no pior caso O(n)")
	}

	if command == "insertion sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.InsertionSort(arr, len(arr))

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n)")
		fmt.Println("Complexidade no caso medio O(n^2)")
		fmt.Println("Complexidade no pior caso O(n^2)")
	}

	if command == "selection sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.SelectionSort(arr, len(arr))

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n^2)")
		fmt.Println("Complexidade no caso medio O(n^2)")
		fmt.Println("Complexidade no pior caso O(n^2)")
	}

	if command == "bubble sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.BubbleSort(arr, len(arr))

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n)")
		fmt.Println("Complexidade no caso medio O(n^2)")
		fmt.Println("Complexidade no pior caso O(n^2)")
	}

	if command == "quick sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.QuickSort(arr, 0, len(arr)-1)

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n log n)")
		fmt.Println("Complexidade no caso medio O(n log n)")
		fmt.Println("Complexidade no pior caso O(n^2)")
	}

	if command == "merge sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.MergeSort(arr, 0, len(arr)-1)

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n log n)")
		fmt.Println("Complexidade no caso medio O(n log n)")
		fmt.Println("Complexidade no pior caso O(n log n)")
	}

	if command == "heap sort" {
		fmt.Printf("Array original: %v\n", arr)

		ordenacao.HeapSort(arr)

		fmt.Printf("Array ordenado: %v\n", arr)
		fmt.Println("Complexidade no melhor caso O(n log n)")
		fmt.Println("Complexidade no caso medio O(n log n)")
		fmt.Println("Complexidade no pior caso O(n log n)")
	}
}
