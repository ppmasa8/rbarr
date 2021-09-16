package rbarr

import "testing"

/* string */
func Test_str_pop(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday"}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_shift(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"monday", "tuesday", "wednesday", "thursday"}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_push(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	arr.push("friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_unshift(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"thursday", "friday", "saturday", "sunday", "monday", "tuesday", "wednesday", "thursday"}
	arr.unshift("thursday", "friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_delete(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := strArray{"monday", "tuesday", "wednesday"}
	arr.delete("sunday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_uniq(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday"}
	arr.uniq()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	// Back up the function pointer to os.Exit()
	oldExit := osExitForStr

	// Return the function pointer that was backed up to osExit after the test.
	defer func() { osExitForStr = oldExit }()

	// When osExit() is executed, the exit code is recorded in the variable called.
	var status int
	exit := func(code int) {
		status = code
	}
	osExitForStr = exit
	// Execute the method under test and check the results.
	arr_blank.pop()
	if exp := 1; status != exp {
		t.Errorf("Expected exit code: %d, status: %d", exp, status)
	}
}

func Test_str_size(t *testing.T) {
	// Normal System
	var arr strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := 5
	size := arr.size()
	if size != expect {
		t.Errorf("result=%v, expect=%v", size, expect)
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	expect_zero := 0
	blank_size := arr_blank.size()
	if blank_size != expect_zero {
		t.Errorf("result=%v, expect=%v", size, expect)
	}
}

func Test_str_include(t *testing.T) {
	// Normal System
	var arr1 strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect1 := true
	include1 := arr1.include("sunday")
	if include1 != expect1 {
		t.Errorf("result=%v, expect=%v", include1, expect1)
	}

	var arr2 strArray = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect2 := false
	include2 := arr2.include("saturday")
	if include2 != expect2 {
		t.Errorf("result=%v, expect=%v", include2, expect2)
	}

	// Abnormal System
	var arr_blank strArray = strArray{}
	expect_blank := false
	blank_include := arr_blank.include("saturday")
	if blank_include != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_include, expect_blank)
	}

}
