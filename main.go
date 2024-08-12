package main

import "fmt"

func main() {
	t := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotateArray(6, t))
}

func rotateArray(times int, arr []int) []int {
	if len(arr) < 2 || times < 1 { // there’s no way to operate with only one element in array or a time less than 1
		return arr
	}
	modTimes := times % len(arr)
	return append(arr[(len(arr)-modTimes):], arr[:(len(arr))-modTimes]...) // We’ve just create an array on-fly without use a new variable
}
