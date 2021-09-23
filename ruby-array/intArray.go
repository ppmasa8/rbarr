package rbarr

import (
	"os"
)

type intArray []int

var osExitForInt = os.Exit

/* int */
// instance method Array#pop
// pop -> object | nil
func (p *intArray) pop() int {
	if len(*p) == 0 {
		osExitForInt(1)
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
		osExitForInt(1)
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
	if len(*p) == 0 { osExitForInt(1) }
	slice := *p
	for _, v := range val {
		slice = append(slice, v)
	}
	*p = slice
}

// instance method Array#unshift
// unshift(*obj) -> self
func (p *intArray) unshift(val ...int) {
	if len(*p) == 0 { osExitForInt(1) }
	slice := *p
	slice = append(slice[:0], val...)
	*p = slice
}

// instance method Array#delete
// delete(val) { ... } -> object
func (p *intArray) delete(val int) {
	if len(*p) == 0 { osExitForInt(1) }
	temp := *p; res := intArray{}
	for _, v := range temp {
		if v == val {
			continue
		}
		res = append(res, v)
	}
	*p = res
}

// instance method Array#uniq
// uniq -> Array
func (p *intArray) uniq() {
	if len(*p) == 0 { osExitForInt(1) }
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
		osExitForInt(1)
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
		osExitForInt(1)
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

// instance method Array#min
// min -> object | nil
func (p *intArray) min() int {
	if len(*p) == 0 {
		osExitForInt(1)
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

// instance method Array#size
// size -> Integer
func (p *intArray) size() int {
	if len(*p) == 0 { return 0 }

	temp := *p
	var cnt int
	for range temp { cnt++ }
	return cnt
}

// instance method Array#include
// include(val) -> bool
func (p *intArray) include(val int) bool {
	if len(*p) == 0 { return false }

	temp := *p
	for _, v := range temp {
		if v == val {
			return true
		}
	}
	return false
}

// instance method Array#first
// first -> object | -1
func (p *intArray) first() int {
	if len(*p) == 0 { return -1 }
	temp := *p
	return temp[0]
}

// instance method Array#last
// last -> object | -1
func (p *intArray) last() int {
	if len(*p) == 0 { return -1 }
	temp := *p
	return temp[len(temp) - 1]
}

// instance method Array#combination
// combination(n) -> Enumerator
func (p *intArray) combination(n int) [][]int {
	temp := *p
	cnt := 0
	result := [][]int{}

	// Full bit search
	m := len(temp)
	// Here, one bit represents the pattern for each of the m elements.
	for bits := 0; bits < (1 << uint64(m)); bits++ {
		// Loop for each element of bits
		for i := 0; i < m; i++ {
			// Check whether the state of the i-th element of the bit is on or not.
			if (bits >> uint64(i)) & 1 == 1 { cnt++ }
		}
		// If the number of items that are on is n, then output.
		if cnt == n {
			unit := intArray{}
			for i := 0; i < m; i++ {
				if (bits >> uint64(i)) & 1 == 1 { unit = append(unit, temp[i]) }
			}
			result = append(result, unit)
			unit = intArray{}
		}
		cnt = 0
	}
	return result
}

// instance method Array#reverse
// reverse -> Array
func (p *intArray) reverse() intArray {
	temp := *p
	res := intArray{}
	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}
	return res
}

// instance method Array#empty
// empty -> bool
func (p *intArray) empty() bool {
	return len(*p) == 0
}

//
// sort
//
func merge(a intArray, b intArray) intArray {
	var r = make(intArray, len(a)+len(b))
	var i = 0
	var j = 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			r[i+j] = a[i]
			i++
		} else {
			r[i+j] = b[j]
			j++
		}
	}

	for i < len(a) {
		r[i+j] = a[i]
		i++
	}
	for j < len(b) {
		r[i+j] = b[j]
		j++
	}

	return r
}

func Mergesort(temp intArray) intArray {
	if len(temp) < 2 {
		return temp
	}
	var mid = len(temp) / 2
	var a = Mergesort(temp[:mid])
	var b = Mergesort(temp[mid:])
	return merge(a, b)
}

func (p *intArray) sort() intArray {
	temp := *p
	temp = Mergesort(temp)
	*p = temp
	return *p
}
