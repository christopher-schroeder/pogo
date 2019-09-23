# pogo
Texas hold em poker hand evaluator in go

# Installation

```
go get github.com/christopher-schroeder/pogo
```
# Instruction

Each card is encoded as a single uint8 by the following scheme.

|    |  ♠ |  ♥ |  ♦ |  ♣ |
|----|----|----|----|----|
|  A |  0 |  1 |  2 |  3 |
|  K |  4 |  5 |  6 |  7 |
|  Q |  8 |  9 | 10 | 11 |
|  J | 12 | 13 | 14 | 15 |
|  T | 16 | 17 | 18 | 19 |
|  9 | 20 | 21 | 22 | 23 |
|  8 | 24 | 25 | 26 | 27 |
|  7 | 28 | 29 | 30 | 31 |
|  6 | 32 | 33 | 34 | 35 |
|  5 | 36 | 37 | 38 | 39 |
|  4 | 40 | 41 | 42 | 43 |
|  3 | 44 | 45 | 46 | 47 |
|  2 | 48 | 49 | 50 | 51 |

The pogo.GetRank function with seven uint8 parameters (encoded cards) caluclates a uint16 rank for a given set of seven cards. Higher poker hands get higher ranks, while equal hands tie equal ranks. The higherst possible score is 7462 (Royal Flush).

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
