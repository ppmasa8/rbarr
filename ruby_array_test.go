package array

import "testing"

func Test_pop(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	exp := intArray{1, 2, 3, 4}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != exp[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}
