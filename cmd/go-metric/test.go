package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"math/rand"
	"time"
)

func main() {
	println("GO")
	tm := metrics.NewTimer()
	for i := 0; i <= 1000000; i++ {
		if i % 1000 == 0 {
			fmt.Printf("%.1f\n", tm.Rate1())
			println(tm.Count())
		}

		time.Sleep(10 * time.Millisecond)
		tm.Update(time.Duration(rand.Intn(100)) * time.Nanosecond)
	}

	//println(tm.Percentile(0.99))
	//fmt.Printf("%v\n", tm.Percentiles([]float64{0.1, 0.5, 0.99, 0.999}))
	//
	//t := time.Now().UnixNano() / 1000000
	//fmt.Println(t)
}
