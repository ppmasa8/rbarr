# rbarr
It can handle arrays like ruby.
https://docs.ruby-lang.org/ja/latest/class/Array.html

## Installation
```
$ go get github.com/ppmasa8/rbarr
```

## Index
### for []int
- pop
- shift
- push
- unshift
- delete
- sum
- max
- min
- size
- uniq
- include
- first
- last
- combination(alfa)

### for []string
- pop
- shift
- push
- unshift
- delete
- size
- uniq
- include
- first
- last

## Usage
```go
import {
  "github.com/ppmasa8/rbarr"
}

array := [5]intArray{1, 2, 3, 4, 5}

array.pop()           // return 5, array = [1, 2, 3, 4]

array.shift()         // return 1, array = [2, 3, 4, 5]

array.push(6, 7)      // return nil, array = [1, 2, 3, 4, 5, 6, 7]

array.unshift(0, 1)   // return nil, array = [0, 1, 1, 2, 3, 4, 5]

array.delete(1)       // return nil, array = [2, 3, 4, 5]

array.sum()           // return 15

array.max()           // return 5

array.min()           // return 1

array.size()          // return 5

array.include(3)      // return true

array.first()         // return 1

array.last()          // return 5

array.combination(2)  // return [[1 2] [1 3] [2 3] [1 4] [2 4] [3 4] [1 5] [2 5] [3 5] [4 5]]

array := [5]intArray{1, 1, 3, 4, 5}

array.uniq()          // return nil, array = [1, 3, 4, 5]
```

## License
MIT license

## Author
https://github.com/ppmasa8
