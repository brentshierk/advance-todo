package Searching

func NaiveMatch(task string,pattern string) []int {
	var occurrence []int
	for i := range task {
		for j := range pattern {
			if task[i+j] != pattern[j] {
				break
			} else if j == (len(pattern) - 1) {
				occurrence = append(occurrence, i+1)
			}
		}
	}
	return occurrence
}
