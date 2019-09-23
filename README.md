# pogo
Texas hold em poker hand evaluator in go

# Installation

```
go get github.com/christopher-schroeder/pogo
```
# Instruction



|    |   |   |   |   |
|----|---|---|---|---|
|  A |   |   |   |   |
|  K |   |   |   |   |
|  Q |   |   |   |   |
|  J |   |   |   |   |
|  T |   |   |   |   |
|  9 |   |   |   |   |
|  8 |   |   |   |   |
|  7 |   |   |   |   |
|  6 |   |   |   |   |
|  5 |   |   |   |   |
|  4 |   |   |   |   |
|  3 |   |   |   |   |
|  2 |   |   |   |   |


# Example
```go
package main

import (
	"fmt"
	"github.com/christopher-schroeder/pogo"
)

func main() {
	fmt.Println(pogo.GetRank(10,12,14,27,7,51,1))
}
```

# How does it work?

pogo is a direct go implementation of https://github.com/kennethshackleton/SKPokerEval. Please have a look, if you are interested in the algorithm.
