// This is here to document what existed before, you should run the benchmark
// tests in go/bench_test.go instead.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s inputfile\n", os.Args[0])
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Printf("Usage: %s inputfile\n", os.Args[0])
		log.Fatal(err)
	}
	dec := json.NewDecoder(file)
	startTime := time.Now()
	var j JSONBlob
	if err := dec.Decode(&j); err != nil && err != io.EOF {
		log.Fatal(err)
	}
	duration := time.Since(startTime)
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	speed := float32(fileInfo.Size()) / 1024.0 / 1024.0 / float32(duration.Seconds())
	fmt.Printf("%.02f MB/s in %s\n", speed, duration)
}
