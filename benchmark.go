package utils

import (
	"log"
	"time"
)

func Benchmark(taskName string, fn func()) {
	start := time.Now()
	fn()
	log.Printf("%s took %s", taskName, time.Since(start))
}
