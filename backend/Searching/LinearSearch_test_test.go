package Searching_test

import (
	"backend/Searching"
	"testing"
)

func TestLinearsearch(t *testing.T) {
	arr :=[]int{4,7,8,2,1,99,6,22,10}
	ans := Searching.Linearsearch(arr,6)
	if ans != true {
		t.Errorf("LinearSearch(arr,6)=%d; want true")
	}
}

func TestLinearsearch2(t *testing.T) {
	arr :=[]int{4,7,8,2,1,99,6,22,10}
	ans := Searching.Linearsearch(arr,33)
	if ans != false {
		t.Errorf("LinearSearch(arr,6)=%d; want false")
	}
}