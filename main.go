package main

import "fmt"

func main() {
	var qqch = []int32{1, 3, 9, 9, 9, 9, 3, 3, 3, 3, 2, 4, 9, 4, 4, 9, 9, 1, 1, 1, 8, 8, 8, 8, 8, 8, 8}

	num := flamingo(qqch)

	fmt.Println(num)
}

func hawk(arr []int32) int32 {
	var num int32 = 0
	//mp := map[int32]int32 {}



}

func flamingo(arr []int32) int32 {
	var tmp int32 = 0
	var ltmp int32 = 0
	var num int32 = 0
	mp := map[int32]int32 {}
	
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] == arr[i] {
				num++
			}
		}

		if num > tmp || num == tmp {
			tmp = num
			ltmp = arr[i]
			mp[arr[i]] = num
		}

		num = 0
		
	}

	for k, e := range mp {
		if e == tmp {
			if k < ltmp {
				ltmp = k
			}
		}
	}

	return ltmp

}