package main

import "fmt"

func main() {
	cards := newDeck()
	hand, _ := cards.deal(5)
	hand.saveToFile("hand.txt")
	fmt.Println("Hand saved: ", hand)
	handFromHardDisk := newDeckFromFile("hand.txt")
	fmt.Println("Loaded from disk: ", handFromHardDisk)
}
