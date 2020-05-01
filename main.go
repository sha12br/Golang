package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"   // alternate way to the above declaration
	// card = "Five of Diamonds" // assigning a new value
	// card = newCard()          // u can also initialize a func to a var as card := newCard()

	// fmt.Println(card)
	// printState() will execute from state.go from same package
	// go run main.go state.go --> to execute the message

	// Slice
	// cards := newDeck()

	// // cards.print() // this calls print function from deck.go

	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()
	// cards := newDeck()

	// cards.saveToFile("myCards")

	cards := newDeckFromFile("myCards")
	cards.print()

}
