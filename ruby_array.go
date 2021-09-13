package rbarr

import (
	"os"
)

type intArray []int

type strArray []string

var osExit = os.Exit

/* int */
// instance method Array#pop
// pop -> object | nil
func (p *intArray) pop() int {
	if len(*p) == 0 {
		osExit(1)
		return -1
	}
	slice := *p
	last  := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	*p = slice
	return last
}

// instance method Array#shift
// shift -> object | nil
func (p *intArray) shift() int {
	if len(*p) == 0 {
		osExit(1)
		return -1
	}
	slice := *p
	first := slice[0]
	slice = slice[1:]
	*p = slice
	return first
}

// instance method Array#append
// push(*obj) -> self
func (p *intArray) push(val ...int) {
	if len(*p) == 0 { osExit(1) }
	slice := *p
	for _, v := range val {
		slice = append(slice, v)
	}
	*p = slice
}

// instance method Array#unshift
// unshift(*obj) -> self
func (p *intArray) unshift(val ...int) {
	if len(*p) == 0 { osExit(1) }
	slice := *p
	slice = append(slice[:0], val...)
	*p = slice
}

func (p *intArray) uniq() {
	if len(*p) == 0 { osExit(1) }
	temp := *p
	m := make(map[int]bool)
	var uniq []int

	for _, v := range temp {
		if !m[v] {
			m[v] = true
			uniq = append(uniq, v)
		}
	}
	*p = uniq
}

// instance method Enumerable#sum
// sum() -> object
func (p *intArray) sum() int {
	if len(*p) == 0 {
		osExit(1)
		return -1
	}
	temp := *p
	var sum int

	for _, v := range temp {
		sum = sum + v
	}
	return sum
}

// instance method Array#max
// max -> object
func (p *intArray) max() int {
	if len(*p) == 0 {
		osExit(1)
		return -1
	}
	temp := *p
	max := temp[0]

	for _, v := range temp {
		if max < v {
			max = v
		}
	}
	return max
}

func (p *intArray) min() int {
	if len(*p) == 0 {
		osExit(1)
		return -1
	}
	temp := *p
	min := temp[0]

	for _, v := range temp {
		if min > v {
			min = v
		}
	}
	return min
}


/* string */
// instance method Array#pop
// pop -> object | nil
func (p *strArray) pop() string {
	if len(*p) == 0 {
		osExit(1)
		return "-1"
	}
	slice := *p
	last  := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	*p = slice
	return last
}

// instance method Array#shift
// shift -> object | nil
func (p *strArray) shift() string {
	if len(*p) == 0 {
		osExit(1)
		return "-1"
	}
	slice := *p
	first := slice[0]
	slice = slice[1:]
	*p = slice
	return first
}

// instance method Array#append
// push(*obj) -> self
func (p *strArray) push(val ...string) {
	if len(*p) == 0 { osExit(1) }
	slice := *p
	for _, v := range val {
		slice = append(slice, v)
	}
	*p = slice
}

// instance method Array#unshift
// unshift(*obj) -> self
func (p *strArray) unshift(val ...string) {
	if len(*p) == 0 { osExit(1) }
	slice := *p
	slice = append(slice[:0], val...)
	*p = slice
}

func (p *strArray) uniq() {
	if len(*p) == 0 { osExit(1) }
	temp := *p
	m := make(map[string]bool)
	var uniq []string

	for _, v := range temp {
		if !m[v] {
			m[v] = true
			uniq = append(uniq, v)
		}
	}
	*p = uniq
}
