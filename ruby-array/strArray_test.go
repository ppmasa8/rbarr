package rbarr

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

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

	// Abnormal system
	var blankArr = strArray{}

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

	expected := "strArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
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

	// Abnormal system
	var blankArr = strArray{}

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

	expected := "strArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
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

	// Abnormal system
	var blankArr = strArray{}

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

	// Abnormal system
	var blankArr = strArray{}

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

	// Abnormal system
	var blankArr = strArray{}

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

	blankArr.delete("0")

	expected := "strArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
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

	// Abnormal system
	var blankArr = strArray{}

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

	expected := "strArray is null"
	actual := strings.TrimRight(buf.String(), "\n")

	if actual != expected {
		t.Fatalf("failed! actual = %s, expected = %s", actual, expected)
	}
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
	var arrBlank = strArray{}
	expectZero := 0
	blankSize := arrBlank.size()
	if blankSize != expectZero {
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
	var arrBlank = strArray{}
	expectBlank := false
	blankInclude := arrBlank.include("saturday")
	if blankInclude != expectBlank {
		t.Errorf("result=%v, expect=%v", blankInclude, expectBlank)
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
	var arrBlank = strArray{}
	expectBlank := "nil"
	blankFirst := arrBlank.first()
	if blankFirst != expectBlank {
		t.Errorf("result=%v, expect=%v", blankFirst, expectBlank)
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
	var arrBlank = strArray{}
	expectBlank := "nil"
	blankLast := arrBlank.last()
	if blankLast != expectBlank {
		t.Errorf("result=%v, expect=%v", blankLast, expectBlank)
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
	var overArr = strArray{"sunday", "monday", "tuesday", "wednesday", "thursday"}
	var overExpect = strArray{}
	overCombination := overArr.combination(6)
	if len(overCombination) != len(overExpect) {
		t.Errorf("result=%v, expect=%v", overCombination, overExpect)
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

	var blankArr = strArray{}
	if blankArr.empty() != true {
		t.Errorf("result=%v, expect=%v", blankArr.empty(), true)
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
