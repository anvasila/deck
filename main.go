package main

func main(){
	//cards := newDeck()
	//
	//hand, remainingCards := deal(cards, 5)
	//
	//hand.print()
	//remainingCards.print()

	// Create And Save Deck
	//cards := newDeck()
	//cards.saveToFile("deck")
	//fmt.Println( cards.toString() )

	// Read from file Deck. If no file exist create one
	//cards := newDeckFromFile("deck")
	//cards.print()

	// Create a deck and shuffle it
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
