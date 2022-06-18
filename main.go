package main

import "fmt"

func main() {
	var someArr = []int32{1, 3, 9, 9, 9, 9, 3, 3, 3, 3, 2, 4, 9, 4, 4, 9, 9, 1, 1, 1, 8, 8, 8, 8, 8, 8, 8}

	num := hawk(someArr)

	fmt.Println(num)
}

func hawk(arr []int32) int32 {
	var i int32 = 0
	var num int32 = 0
	var bArr = make([]int32, 5)

	for i = 1; i < 6; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == i {
				num++
			}
		}

		bArr[i-1] = num
		num = 0
	}

	var tmp int32 = 0
	var ltmp int32 = 0

	for k, e := range bArr {
		if e > tmp {
			tmp = e
			ltmp = int32(k+1)
		}
	}

	return ltmp
}