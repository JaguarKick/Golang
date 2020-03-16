package main

import "fmt"

type game struct {
	gameTitle string
	gamePublisher string
	gamePlatform string
}

func(g *game) updateTitle(title string) {
	(*g).gameTitle = title
}

func(g *game) updatePublisher(publisher string) {
	(*g).gamePublisher = publisher
}

func(g *game) updatePlatform(platform string) {
	(*g).gamePlatform = platform
}

func main() {
	var gow game
	gow.updateTitle("God of War");
	gow.updatePublisher("Santa Monica")
	gow.updatePlatform("PS4")
	fmt.Println(gow)
}
