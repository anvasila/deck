package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// Create and return a list of playing cards.
// Essentially an array of strings
func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

// Receiver!!!
func (d deck) print( ) {
	for i,card := range d{
		fmt.Println( i, card )
	}
}

// Create a 'hand' of cards.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize],d[handSize:]
}

// Save a list of cards to a file on the local machine
func (d deck) saveToFile( filename string ) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (d deck) toString() string {
	return strings.Join([]string(d),",")
}

// Load a list of cards from the local machine
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile( fileName )
	if err != nil {
		// Option #1 - Log the error and create new deck from a call to newDeck(),
		// save the deck and return it
		log.Println(err)
		d := newDeck()
		fmt.Println("ftanei")
		d.saveToFile(fileName)
		return d

		// Option #2 - Log the error and entirely quit the program
		//log.Fatal(err) // Fatal include the os.Exit(1)
	}

	return stringToDeck( string(bs) )
}

func stringToDeck( s string ) deck {
	return deck( strings.Split( s, "," ) )
}