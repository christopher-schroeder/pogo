package pogo

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	start := time.Now()
	var r uint16

	for i := 0; i < 1000000000; i++ {
		r = GetRank(0, 4, 8, 12, 16, 20, 24)
	}
	fmt.Println(r)

	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)
}