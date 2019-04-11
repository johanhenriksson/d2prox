package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/johanhenriksson/d2prox"
	"github.com/johanhenriksson/d2prox/ip"
)

func main() {
	// greet
	fmt.Println("core | d2prox by @johanhenriksson")

	// env config: realm hostname
	realmHost := GetEnv("REALM", "europe.battle.net")
	fmt.Println("core | destination realm:", realmHost)

	// env config: local machine proxy
	local := GetEnvBool("LOCAL", false)
	if local {
		fmt.Println("core | running on localhost")
	} else {
		// resolve public IP
		ip, err := ip.ResolvePublicIP()
		if err != nil {
			fmt.Println("core | Error resolving public IP address!")
			return
		}
		fmt.Println("core | running as public. ip address resolved to", ip)
	}

	// set up battle.net proxy
	bnet := d2prox.NewBnet(realmHost)
	go d2prox.Serve(bnet)

	// set up realm server proxy
	realm := d2prox.NewRealm()
	go d2prox.Serve(realm)

	// set up game server proxy
	game := d2prox.NewGame()
	go d2prox.Serve(game)

	// wait forever.
	// todo: console command loop
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

// GetEnvBool gets an environment variable and casts it to boolean
func GetEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		switch strings.ToLower(value) {
		case "1":
			return true
		case "yes":
			return true
		case "true":
			return true
		}
		return false
	}
	return fallback
}

// GetEnv gets an environment variable and returns it as a string
func GetEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
