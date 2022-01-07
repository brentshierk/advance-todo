package Searching

import (
	"math"
)


func jumpSearch(arr []int, s int) int {

	var blockSize = int(math.Sqrt(float64(len(arr))))

	var i = 0
	for {

		if arr[i] >= s {
			break
		}

		if i > len(arr) {
			break
		}

		i += blockSize
	}

	for j := i; j > 0; j-- {
		if arr[j] == s {
			return j
		}
	}

	return -1
}
//for testing
//    var n = []int{2, 9, 11, 21, 22, 32, 36, 48, 76, 90}