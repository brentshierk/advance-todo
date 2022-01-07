package Searching

import "fmt"

func BinarySearch(sortedList []int, lookingFor int) int {
	var lo int = 0
	var hi int = len(sortedList) - 1

	for lo <= hi {
		var mid int = lo + (hi-lo)/2
		var midValue int = sortedList[mid]
		fmt.Println("Middle value is:", midValue)

		if midValue == lookingFor {
			return mid
		} else if midValue > lookingFor {
			// We want to use the left half of our list
			hi = mid - 1
		} else {
			// We want to use the right half of our list
			lo = mid + 1
		}
	}

	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	return -1
}

//the second form of implementation is usings Go's standard lib "sort" that has a Search
//function that is already built in and makes use of binary tree data struct

//n int, f func(int) bool) int {
	// Define f(-1) == false and f(n) == true.
	// Invariant: f(i-1) == false, f(j) == true.
	//i, j := 0, n
	//for i < j {
		//h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		//if !f(h) {
		//	i = h + 1 // preserves f(i-1) == false
		//} else {
			//j = h // preserves f(j) == true
//		}
//	}
//	// i == j, f(i-1) == false, and f(j) (= f(i)) == true  =>  answer is i.
//	//return i
//}
//

// implement this test
//a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
//x := 6
