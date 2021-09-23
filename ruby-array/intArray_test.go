package rbarr

import (
	"testing"
)

/* int */
func Test_int_pop(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4}
	return_int := arr.pop()
	if return_int != 5 {
		t.Errorf("Return integer is wrong. result=%v, expect=%v", return_int, 5)
	}
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("Return array is wrong. result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_shift(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{2, 3, 4, 5}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_push(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4, 5, 6, 7, 8}
	arr.push(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_unshift(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3 ,4, 5}
	expect := intArray{6, 7, 8, 1, 2, 3 ,4, 5}
	arr.unshift(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_delete(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3 ,4, 5}
	expect := intArray{1, 2, 3, 4}
	arr.delete(5)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_uniq(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 1, 3 ,4, 5}
	expect := intArray{1, 3 ,4, 5}
	arr.uniq()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_sum(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 15
	sum := arr.sum()
	if sum != expect {
		t.Errorf("result=%v, expect=%v", sum, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_max(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 5
	max := arr.max()
	if max != expect {
		t.Errorf("result=%v, expect=%v", max, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_min(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 1
	min := arr.min()
	if min != expect {
		t.Errorf("result=%v, expect=%v", min, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForInt

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForInt = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForInt = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_int_size(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 5
	size := arr.size()
	if size != expect {
		t.Errorf("result=%v, expect=%v", size, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	expect_zero := 0
	blank_size := arr_blank.size()
	if blank_size != expect_zero {
		t.Errorf("result=%v, expect=%v", size, expect)
	}
}

func Test_int_include(t *testing.T) {
	// Normal System
	var arr1 intArray = intArray{1, 2, 3, 4, 5}
	expect1 := true
	include1 := arr1.include(1)
	if include1 != expect1 {
		t.Errorf("result=%v, expect=%v", include1, expect1)
	}

	var arr2 intArray = intArray{1, 2, 3, 4, 5}
	expect2 := false
	include2 := arr2.include(28)
	if include2 != expect2 {
		t.Errorf("result=%v, expect=%v", include2, expect2)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	expect_blank := false
	blank_include := arr_blank.include(3)
	if blank_include != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_include, expect_blank)
	}
}

func Test_int_first(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 1
	first := arr.first()
	if first != expect {
		t.Errorf("result=%v, expect=%v", first, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	expect_blank := -1
	blank_first := arr_blank.first()
	if blank_first != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_first, expect_blank)
	}
}

func Test_int_last(t *testing.T) {
	// Normal System
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := 5
	last := arr.last()
	if last != expect {
		t.Errorf("result=%v, expect=%v", last, expect)
	}

	// Abnormal System
	var arr_blank intArray = intArray{}
	expect_blank := -1
	blank_last := arr_blank.last()
	if blank_last != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_last, expect_blank)
	}
}

func Test_int_combination(t *testing.T) {
	// Normal Systems
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := [][]int{{1, 2}, {1, 3}, {2, 3}, {1, 4}, {2, 4}, {3, 4}, {1, 5}, {2, 5}, {3, 5}, {4, 5}}
	combination := arr.combination(2)
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			if combination[i][j] != expect[i][j] {
				t.Errorf("result=%v, expect=%v", combination[i][j], expect[i][j])
			}
		}
	}

	// Abnormal Systems
	var over_arr intArray = intArray{1, 2, 3, 4, 5}
	var over_expect intArray = intArray{}
	over_combination := over_arr.combination(6)
	if len(over_combination) != len(over_expect) {
		t.Errorf("result=%v, expect=%v", over_combination, over_expect)
	}
}

func Test_int_reverse(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	expect := intArray{5, 4, 3, 2, 1}
	rev := arr.reverse()
	for i, v := range rev {
		if v != expect[i] {
			t.Errorf("result=%v, expect=%v", v, expect[i])
		}
	}
}

func Test_int_empty(t *testing.T) {
	var arr intArray = intArray{1, 2, 3, 4, 5}
	if arr.empty() != false {
		t.Errorf("result=%v, expect=%v", arr.empty(), false)
	}

	var blank_arr intArray = intArray{}
	if blank_arr.empty() != true {
		t.Errorf("result=%v, expect=%v", blank_arr.empty(), true)
	}
}

func Test_int_sort(t *testing.T) {
	var arr intArray = intArray{5, 9, 4, 6, 1, 10, 3, 2, 7, 8}
	expect := intArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr.sort()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, expect=%v", arr[i], expect[i])
		}
	}
}