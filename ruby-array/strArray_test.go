package rbarr

import "testing"

/* string */
func Test_str_pop(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday"}
	arr.pop()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_shift(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"monday", "tuesday", "wednesday", "thursday"}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_push(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday", "thursday", "friday", "saturday"}
	arr.push("friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_unshift(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"thursday", "friday", "saturday", "sunday", "monday", "tuesday", "wednesday", "thursday"}
	arr.unshift("thursday", "friday", "saturday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_delete(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := strArray{"monday", "tuesday", "wednesday"}
	arr.delete("sunday")
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_uniq(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := strArray{"sunday", "monday", "tuesday", "wednesday"}
	arr.uniq()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	//// Abnormal System
	//var arr_blank strArray = strArray{}
	//// Back up the function pointer to os.Exit()
	//oldExit := osExitForStr
	//
	//// Return the function pointer that was backed up to osExit after the test.
	//defer func() { osExitForStr = oldExit }()
	//
	//// When osExit() is executed, the exit code is recorded in the variable called.
	//var status int
	//exit := func(code int) {
	//	status = code
	//}
	//osExitForStr = exit
	//// Execute the method under test and check the results.
	//arr_blank.pop()
	//if exp := 1; status != exp {
	//	t.Errorf("Expected exit code: %d, status: %d", exp, status)
	//}
}

func Test_str_size(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := 5
	size := arr.size()
	if size != expect {
		t.Errorf("result=%v, expect=%v", size, expect)
	}

	// Abnormal System
	var arr_blank = strArray{}
	expect_zero := 0
	blank_size := arr_blank.size()
	if blank_size != expect_zero {
		t.Errorf("result=%v, expect=%v", size, expect)
	}
}

func Test_str_include(t *testing.T) {
	// Normal System
	var arr1 = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect1 := true
	include1 := arr1.include("sunday")
	if include1 != expect1 {
		t.Errorf("result=%v, expect=%v", include1, expect1)
	}

	var arr2 = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect2 := false
	include2 := arr2.include("saturday")
	if include2 != expect2 {
		t.Errorf("result=%v, expect=%v", include2, expect2)
	}

	// Abnormal System
	var arr_blank = strArray{}
	expect_blank := false
	blank_include := arr_blank.include("saturday")
	if blank_include != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_include, expect_blank)
	}
}

func Test_str_first(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := "sunday"
	first := arr.first()
	if first != expect {
		t.Errorf("result=%v, expect=%v", first, expect)
	}

	// Abnormal System
	var arr_blank = strArray{}
	expect_blank := "nil"
	blank_first := arr_blank.first()
	if blank_first != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_first, expect_blank)
	}
}

func Test_str_last(t *testing.T) {
	// Normal System
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "sunday"}
	expect := "sunday"
	last := arr.last()
	if last != expect {
		t.Errorf("result=%v, expect=%v", last, expect)
	}

	// Abnormal System
	var arr_blank = strArray{}
	expect_blank := "nil"
	blank_last := arr_blank.last()
	if blank_last != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_last, expect_blank)
	}
}

func Test_str_combination(t *testing.T) {
	// Normal Systems
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := [][]string{{"sunday", "monday"}, {"sunday", "tuesday"}, {"monday", "tuesday"},
		{"sunday", "wednesday"}, {"monday", "wednesday"}, {"tuesday", "wednesday"}, {"sunday", "thursday"},
		{"monday", "thursday"}, {"tuesday", "thursday"}, {"wednesday", "thursday"}}
	combination := arr.combination(2)
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			if combination[i][j] != expect[i][j] {
				t.Errorf("result=%v, expect=%v", combination[i][j], expect[i][j])
			}
		}
	}

	// Abnormal Systems
	var over_arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	var over_expect = strArray{}
	over_combination := over_arr.combination(6)
	if len(over_combination) != len(over_expect) {
		t.Errorf("result=%v, expect=%v", over_combination, over_expect)
	}
}

func Test_str_reverse(t *testing.T) {
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"thursday", "wednesday", "tuesday", "monday", "sunday"}
	rev := arr.reverse()
	for i, v := range rev {
		if v != expect[i] {
			t.Errorf("result=%v, expect=%v", v, expect[i])
		}
	}
}

func Test_str_empty(t *testing.T) {
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	if arr.empty() != false {
		t.Errorf("result=%v, expect=%v", arr.empty(), false)
	}

	var blank_arr = strArray{}
	if blank_arr.empty() != true {
		t.Errorf("result=%v, expect=%v", blank_arr.empty(), true)
	}
}

func Test_str_sort(t *testing.T) {
	var arr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	expect := strArray{"monday", "sunday", "thursday", "tuesday", "wednesday"}
	arr.sort()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, expect=%v", arr[i], expect[i])
		}
	}
}
