package rbarr

import "testing"

// int
func Test_int_pop(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_int_shift(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{2, 3, 4, 5}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_int_push(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4, 5, 6, 7, 8}
	arr.push(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_int_unshift(t *testing.T) {
	var arr intArray = intArray{1, 2, 3 ,4, 5}
	expect := intArray{6, 7, 8, 1, 2, 3 ,4, 5}
	arr.unshift(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}


// str
func Test_str_pop(t *testing.T) {
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday"}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_str_shift(t *testing.T) {
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"monday", "tuesday", "wednesday", "thursday"}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_str_push(t *testing.T) {
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	arr.push("friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}

func Test_str_unshift(t *testing.T) {
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"thursday", "friday", "saturday", "sunday", "monday", "tuesday", "wednesday", "thursday"}
	arr.unshift("thursday", "friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}
}
