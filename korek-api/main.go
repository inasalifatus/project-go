package main

import (
	"fmt"
	"math/rand"
	"os"
	// "time"
)

var expected = 11

func main() {

	c := make(chan *korek)
	done := make(chan *korek)

	go player("User 1", c, done)
	go player("User 2", c, done)
	go player("User 3", c, done)
	go player("User 4", c, done)
	// fmt.Println(<-done)
	clear(c, done)

}

type korek struct {
	hits       int
	lastPlayer string
}

func player(name string, c, done chan *korek) {
	hits := 0
	for {
		random := rand.Intn(100)
		hits += 1
		fmt.Println(random, "ini di hits ke ", hits, "user ", name)
		if random != 0 && hits > 1 {
			if random%expected == 0 {
				fmt.Println("masuk ke expected 1 ", random, expected)
				msg := <-c
				done <- &korek{
					hits:       msg.hits,
					lastPlayer: name,
				}
			} else {
				c <- &korek{
					hits: hits,
				}
			}
			select {
			case msgC := <-c:
				if random%expected == 0 {
					fmt.Println("masuk ke expected 2 ", random, expected)
					msg := <-c
					done <- &korek{
						hits:       msg.hits,
						lastPlayer: name,
					}
				} else {
					msgC.hits = msgC.hits + 1
					c <- &korek{
						hits: msgC.hits,
					}
				}
			case msg := <-done:
				fmt.Println(msg.lastPlayer, "player yang kalah")
				os.Exit(2)
			}
		}
	}
}

func clear(c, done chan *korek) {
	fmt.Println("ini masuk ke clear")
	for {
		select {
		case <-c:
			continue
		case msg := <-done:
			fmt.Printf("pemain ke %s kalah", msg.lastPlayer)
			close(c)
			close(done)
			break
		}
	}
}
