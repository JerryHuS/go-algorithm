package main

func sortByBubble(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return
}

func sortByQuick(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	left, right := low, high
	pivot := (high-low)/2 + low

	for left <= right {
		for arr[left] < arr[pivot] {
			left++
		}
		for arr[right] > arr[pivot] {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}
