/**
  @author: hs
  @date: 2023/10/22
  @note:

modification history
--------------------
**/
package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 4, 3, 3}
	quickSort(arr)
	fmt.Println("arr: ", arr)
}

func quickSort(arr []int) []int {
	return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		_quickSort(arr, left, partitionIndex-1)
		_quickSort(arr, partitionIndex+1, right)
	}

	return arr
}

func partition(arr []int, left, right int) int {
	pivot := left
	index := left + 1

	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, index, i)
			index++
		}
	}
	swap(arr, pivot, index-1)

	return index-1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}