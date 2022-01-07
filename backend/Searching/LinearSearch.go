package Searching


func Linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

// for testing
// var n = []int{9, 1, 33, 21, 78, 42, 4}