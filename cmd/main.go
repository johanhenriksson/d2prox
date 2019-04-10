package main

import (
	"fmt"
	"time"

	"github.com/johanhenriksson/d2prox"
)

func main() {
	fmt.Println("d2prox v0.1.0 by @johanhenriksson")
	fmt.Println()

	bnet := d2prox.NewBnet()
	go d2prox.Serve(bnet)

	realm := d2prox.NewRealm()
	go d2prox.Serve(realm)

	game := d2prox.NewGame()
	go d2prox.Serve(game)

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
