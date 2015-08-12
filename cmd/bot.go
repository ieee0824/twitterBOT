package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/shirase-aoi/twitterBOT"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	postTime := time.Now().Unix()
	postTime = 0
	for {
		if time.Now().Unix()-postTime > (30*60)+time.Now().Unix()%10 {
			bot.RandomPost()
			fmt.Println("post")
		}
		time.Sleep(time.Minute * 1)
	}
}
