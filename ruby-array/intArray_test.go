package rbarr

import "testing"

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
