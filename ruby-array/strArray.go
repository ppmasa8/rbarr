package rbarr

import "os"

type strArray []string

var osExitForStr = os.Exit

/* string */
// instance method Array#pop
// pop -> object | nil
func (p *strArray) pop() string {
	if len(*p) == 0 {
		osExitForStr(1)
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
		osExitForStr(1)
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
	if len(*p) == 0 { osExitForStr(1) }
	slice := *p
	for _, v := range val {
		slice = append(slice, v)
	}
	*p = slice
}

// instance method Array#unshift
// unshift(*obj) -> self
func (p *strArray) unshift(val ...string) {
	if len(*p) == 0 { osExitForStr(1) }
	slice := *p
	slice = append(slice[:0], val...)
	*p = slice
}

// instance method Array#delete
// delete(val) { ... } -> object
func (p *strArray) delete(val string) {
	if len(*p) == 0 { osExitForStr(1) }
	temp := *p; res := strArray{}
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
func (p *strArray) uniq() {
	if len(*p) == 0 { osExitForStr(1) }
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

// instance method Array#size
// size -> Integer
func (p *strArray) size() int {
	if len(*p) == 0 { return 0 }

	temp := *p
	var cnt int
	for range temp { cnt++ }
	return cnt
}

// instance method Array#include
// include(val) -> bool
func (p *strArray) include(val string) bool {
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
// first -> object | "nil"
func (p *strArray) first() string {
	if len(*p) == 0 { return "nil" }
	temp := *p
	return temp[0]
}

// instance method Array#last
// last -> object | "nil"
func (p *strArray) last() string {
	if len(*p) == 0 { return "nil" }
	temp := *p
	return temp[len(temp) - 1]
}

// instance method Array#combination
// combination(n) -> Enumerator
func (p *strArray) combination(n int) [][]string {
	temp := *p
	cnt := 0
	result := [][]string{}

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
			unit := strArray{}
			for i := 0; i < m; i++ {
				if (bits >> uint64(i)) & 1 == 1 { unit = append(unit, temp[i]) }
			}
			result = append(result, unit)
			unit = strArray{}
		}
		cnt = 0
	}
	return result
}

// instance method Array#reverse
// reverse -> Array
func (p *strArray) reverse() strArray {
	temp := *p
	res := strArray{}
	for i := len(temp) - 1; i >= 0; i-- {
		res = append(res, temp[i])
	}
	return res
}
