package main

import (
	"fmt"
	"math/rand"
	"time"
)

type binFunc func(int, int) int

func main() {
	rand.Seed(time.Now().UnixNano())

	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / (y * y + 1) },
		func(x, y int) int { return x % (y * y + 1) },
	}

	x, y := 69691, 69693
	for i := 0; i < 1000; i++ {
		x = fns[rand.Intn(len(fns))](x, y)
		y = fns[rand.Intn(len(fns))](x, y)
		fmt.Printf("%d %d ", x, y)
	}
}
