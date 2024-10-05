package main

import "fmt"

func main() {
	arr := []int{1, 1, 0, -1, -1}
	fmt.Println(PlusMinus(arr))
}

func PlusMinus(arr []int) (float64, float64, float64) {
	var positif, negative, zero float64

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			positif++
		} else if arr[i] < 0 {
			negative++
		} else {
			zero++
		}
	}

	lengthArr := float64(len(arr))
	return positif / lengthArr, negative / lengthArr, zero / lengthArr
}
