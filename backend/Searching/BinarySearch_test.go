package Searching

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1,2,3,4,5}
	ans := BinarySearch(arr,4)
	if ans != 4{
		t.Errorf("BinarySearch(arr,4)= %d;want 4",ans)
	}
}

func TestBinarySearch2(t *testing.T) {
	arr := []int{1,2,3,4,5}
	ans := BinarySearch(arr,6)
	if ans != -1{
		t.Errorf("BinarySearch(arr,6)= %d;want -1",ans)
	}
}
