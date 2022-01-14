package Searching

import (
	"hash/fnv"
	"strconv"
)

// RabinKarpWithoutHash searches for pattern in a given text string without using language built-in hash function
func RabinKarpWithoutHash(text, pattern string) []int {
	arr := make([]int, 0)
	d := 256
	q := 101
	m := len(pattern)
	n := len(text)
	p := 0
	t := 0
	h := 1
	i := 0
	j := 0

	for i = 0; i < m-1; i++ {
		h = (h * d) % q
	}

	for i = 0; i < m; i++ {
		p = ((d * p) + int(pattern[i])) % q
		t = ((d * t) + int(text[i])) % q
	}

	for i = 0; i <= n-m; i++ {
		if p == t {
			for j = 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					break
				}
			}
			if j == m {
				arr = append(arr, i+1)
			}
		}

		if i < n-m {
			t = ((d * (t - int(text[i])*h)) + (int(text[i+m]))) % q
			if t < 0 {
				t += q
			}
		}

	}

	return arr
}

// RabinKarpWithHash searches for pattern in a given text string by using language built-in hash function
func RabinKarpWithHash(text, pattern string) []int {
	arr := make([]int, 0)
	m := len(pattern)
	n := len(text)
	p := ""
	t := ""
	i := 0
	j := 0

	p = hash(pattern[0:m])
	t = hash(text[0:m])

	for i = 0; i <= n-m; i++ {
		if p == t {
			for j = 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					break
				}
			}
			if j == m {
				arr = append(arr, i+1)
			}
		}

		if i < n-m {
			t = hash(text[i : i+m+1])
		}

	}

	return arr
}

func hash(s string) string {
	h := fnv.New32a()
	_, err := h.Write([]byte(s))
	if err != nil {
		return ""
	}
	return strconv.Itoa(int(h.Sum32()))
}