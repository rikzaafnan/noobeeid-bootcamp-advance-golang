package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
)

type Player struct {
	Name string
	Hit  int
}

func main() {
	// set max processor yang akan digunakan
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("running goroutine in", runtime.NumCPU(), "cpu")

	var wg sync.WaitGroup
	gameOver := make(chan bool)

	// a := Player{Name: "A", Hit: 0}
	// b := Player{Name: "B", Hit: 0}

	players := []Player{
		{Name: "a", Hit: 0},
		{Name: "b", Hit: 0},
	}

	for _, v := range players {
		fmt.Println(v)
		wg.Add(1)
		go playPingPong(v.Name, &v, gameOver, &wg)
	}

	wg.Wait()

	// done := make(chan bool)

	// wg.Add(2)
	// go playPingPong("Player A", &a, done, &wg)
	// go playPingPong("Player B", &a, done, &wg)
	// wg.Wait()

	// close(gameOver)
	<-gameOver
	// fmt.Printf("Game Over!\nTotal Hits:\nPlayer A: %d\nPlayer B: %d\n", a.Hit, b.Hit)
	fmt.Printf("Game Over!")

}

func playPingPong(playerName string, player *Player, gameOver chan bool, wg *sync.WaitGroup) {

	defer wg.Done()
	for {

		randValue := rand.Intn(50)
		player.Hit++
		fmt.Printf("%s = Hit %d // counter=%d\n", playerName, player.Hit, randValue)
		if randValue%11 == 0 {
			fmt.Printf("%s kalah, total hit: %d, kalah di nomor %d\n", playerName, player.Hit, randValue)
			gameOver <- true
			break

		}
		// player

	}

}

/*
PINGPONG APPS
2 player => 2 goroutine

kondisi kalah : saat flag/counter/random number, habis dibagi 11 (random % 11 == 0 THAN break)

Contoh :
Player A = Hit 1 // counter 8
Player B = Hit 2 // counter 3
Player A = Hit 3 // counter 24
Player B = Hit 4 // counter 33

Player B kalah, total hit : 4, kalah di nomor 33

Player A = Hit 1 // counter 8
Player B = Hit 2 // counter 9
Player A = Hit 3 // counter 22

Player A kalah, total hit : 3, kalah di nomor 22

Requirement :
- Struct Player {
Name string
Hit int
}

a := Player{Name: A, Hit:0}
b := Player{Name: B, Hit:0}

a.Hit++
b.Hit++

*/
