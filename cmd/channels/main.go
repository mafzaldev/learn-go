package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{
		"amazon",
		"flipkart",
		"bigbasket",
		"grofers",
		"swiggy",
	}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}

	sendNotification(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(1 * time.Second)
		var chickPrice = rand.Float32() * 20

		if chickPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendNotification(chickenChannel chan string) {
	fmt.Println("Deal:", <-chickenChannel, "has chicken below", MAX_CHICKEN_PRICE, "bucks!")
}
