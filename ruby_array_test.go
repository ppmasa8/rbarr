package array

import "testing"

func Test_pop(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_shift(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{2, 3, 4, 5}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}
