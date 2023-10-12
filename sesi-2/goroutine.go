package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// set max processor yang akan digunakan
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	// running in new goroutine
	go speak(2, "Goroutine 1")
	// running in new goroutine
	go speak(2, "Goroutine 2")
	// running in new goroutine
	go speak(2, "Goroutine 3")

	// running in main goroutine
	speak(2, "Hello")

	time.Sleep(1 * time.Second)
}

func speak(total int, message string) {
	for i := 0; i < total; i++ {
		fmt.Println("Speak", message, "- (", i+1, ")")
	}
}
