package array

type intArray []int

// instance method Array#pop
// pop -> object | nil
func (p *intArray) pop() int {
	slice := *p
	last  := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	*p = slice
	return last
}

// instance method Array#shift
// shift -> object | nil
func (p *intArray) shift() int {
	slice := *p
	first := slice[0]
	slice = slice[1:]
	*p = slice
	return first
}