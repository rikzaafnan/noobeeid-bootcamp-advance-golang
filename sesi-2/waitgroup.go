package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// membuat object waitgroup
	wg := sync.WaitGroup{}
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	// // memberitahu, jumlah goroutine yang akan dibuat
	// wg.Add(3)

	// go speak(2, "Goroutine 1", &wg)
	// go speak(2, "Goroutine 2", &wg)
	// go speak(2, "Goroutine 3", &wg)

	// atau dengna cara
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go speak(2, fmt.Sprintf("Speak Goroutine %v", i+1), &wg)
	}

	// karena disini kita tidak membutuhkan goroutine,
	// jadi value waitgroupnya kita isi nil
	speak(2, "Hello", nil)

	// // proses menunggu sampai semua waitgroup telah di `done`-in
	wg.Wait()

	fmt.Println("Doneee")
}

// menambahkan parameter WaitGroup
// dengan value pointer
func speak(total int, message string, wg *sync.WaitGroup) {
	// check apakah waitgroup nil atau engga
	if wg != nil {
		// proses untuk menandakan bahwa goroutine telah selesai
		defer wg.Done()
	}
	for i := 0; i < total; i++ {
		fmt.Println("Speak", message, "- (", i+1, ")")
	}

}
