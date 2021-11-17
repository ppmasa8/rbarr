package rbarr

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

/* int */
func Test_int_pop(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4}
	returnInt := arr.pop()
	if returnInt != 5 {
		t.Errorf("Return integer is wrong. result=%v, expect=%v", returnInt, 5)
	}
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("Return array is wrong. result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.pop()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_shift(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := intArray{2, 3, 4, 5}
	arr.shift()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.shift()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_push(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := intArray{1, 2, 3, 4, 5, 6, 7, 8}
	arr.push(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.push()

	expected := "argument is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_unshift(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3 ,4, 5}
	expect := intArray{6, 7, 8, 1, 2, 3 ,4, 5}
	arr.unshift(6, 7, 8)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.unshift()

	expected := "argument is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_delete(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3 ,4, 5}
	expect := intArray{1, 2, 3, 4}
	arr.delete(5)
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.delete(0)

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_uniq(t *testing.T) {
	// Normal System
	var arr = intArray{1, 1, 3 ,4, 5}
	expect := intArray{1, 3 ,4, 5}
	arr.uniq()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, index=%v", arr, i)
		}
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.uniq()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_sum(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 15
	sum := arr.sum()
	if sum != expect {
		t.Errorf("result=%v, expect=%v", sum, expect)
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.sum()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_max(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 5
	max := arr.max()
	if max != expect {
		t.Errorf("result=%v, expect=%v", max, expect)
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.max()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_min(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 1
	min := arr.min()
	if min != expect {
		t.Errorf("result=%v, expect=%v", min, expect)
	}

	// Abnormal system
	var blankArr = intArray{}

	// Change log output destination to buffer.
	var buf bytes.Buffer
	log.SetOutput(&buf)

	// Set the flag to 0 because the default is to output the date.
	defaultFlags := log.Flags()
	log.SetFlags(0)

	// Return changed content when Test_printLog testing finished.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(defaultFlags)
	}()

	blankArr.min()

	expected := "intArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
}

func Test_int_size(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 5
	size := arr.size()
	if size != expect {
		t.Errorf("result=%v, expect=%v", size, expect)
	}

	// Abnormal System
	var arr_blank = intArray{}
	expect_zero := 0
	blank_size := arr_blank.size()
	if blank_size != expect_zero {
		t.Errorf("result=%v, expect=%v", size, expect)
	}
}

func Test_int_include(t *testing.T) {
	// Normal System
	var arr1 = intArray{1, 2, 3, 4, 5}
	expect1 := true
	include1 := arr1.include(1)
	if include1 != expect1 {
		t.Errorf("result=%v, expect=%v", include1, expect1)
	}

	var arr2 = intArray{1, 2, 3, 4, 5}
	expect2 := false
	include2 := arr2.include(28)
	if include2 != expect2 {
		t.Errorf("result=%v, expect=%v", include2, expect2)
	}

	// Abnormal System
	var arr_blank = intArray{}
	expect_blank := false
	blank_include := arr_blank.include(3)
	if blank_include != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_include, expect_blank)
	}
}

func Test_int_first(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 1
	first := arr.first()
	if first != expect {
		t.Errorf("result=%v, expect=%v", first, expect)
	}

	// Abnormal System
	var arr_blank = intArray{}
	expect_blank := -1
	blank_first := arr_blank.first()
	if blank_first != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_first, expect_blank)
	}
}

func Test_int_last(t *testing.T) {
	// Normal System
	var arr = intArray{1, 2, 3, 4, 5}
	expect := 5
	last := arr.last()
	if last != expect {
		t.Errorf("result=%v, expect=%v", last, expect)
	}

	// Abnormal System
	var arr_blank = intArray{}
	expect_blank := -1
	blank_last := arr_blank.last()
	if blank_last != expect_blank {
		t.Errorf("result=%v, expect=%v", blank_last, expect_blank)
	}
}

func Test_int_combination(t *testing.T) {
	// Normal Systems
	var arr = intArray{1, 2, 3, 4, 5}
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
	var over_arr = intArray{1, 2, 3, 4, 5}
	var over_expect = intArray{}
	over_combination := over_arr.combination(6)
	if len(over_combination) != len(over_expect) {
		t.Errorf("result=%v, expect=%v", over_combination, over_expect)
	}
}

func Test_int_reverse(t *testing.T) {
	var arr = intArray{1, 2, 3, 4, 5}
	expect := intArray{5, 4, 3, 2, 1}
	rev := arr.reverse()
	for i, v := range rev {
		if v != expect[i] {
			t.Errorf("result=%v, expect=%v", v, expect[i])
		}
	}
}

func Test_int_empty(t *testing.T) {
	var arr = intArray{1, 2, 3, 4, 5}
	if arr.empty() != false {
		t.Errorf("result=%v, expect=%v", arr.empty(), false)
	}

	var blank_arr = intArray{}
	if blank_arr.empty() != true {
		t.Errorf("result=%v, expect=%v", blank_arr.empty(), true)
	}
}

func Test_int_sort(t *testing.T) {
	var arr = intArray{5, 9, 4, 6, 1, 10, 3, 2, 7, 8}
	expect := intArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr.sort()
	for i, _ := range arr {
		if arr[i] != expect[i] {
			t.Errorf("result=%v, expect=%v", arr[i], expect[i])
		}
	}
}