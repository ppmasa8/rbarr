# rbarr
It can handle arrays like ruby.

## Installation
```
$ go get github.com/ppmasa8/rbarr
```

## Usage
```go
import {
  "github.com/ppmasa8/rbarr"
}

array := [5]int{1, 2, 3, 4, 5}

array.pop()           // return 5, array = [1, 2, 3, 4]

array.shift()         // return 1, array = [2, 3, 4, 5]

array.push(6, 7)      // return nil, array = [1, 2, 3, 4, 5, 6, 7]

array.unshift(0, 1)   // return nil, array = [0, 1, 1, 2, 3, 4, 5]

array := [5]int{1, 1, 3, 4, 5}

array.uniq()          // return nil, array = [1, 3, 4, 5]
```

## License
MIT license

## Author
ppmasa8
