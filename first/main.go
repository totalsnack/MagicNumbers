package main

import "fmt"

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int

	intSet := make(map[int]bool, len(arr))
	for _, element := range arr {
		if _, ok := intSet[element]; !ok {
			result = append(result, element)
		}
		intSet[element] = true
	}
	fmt.Println("initial slice: ", arr)
	fmt.Println("result       : ", result)
}
