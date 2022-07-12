package main

func main() {
	/*cards := newDeck()
	cards.saveToFile("my_deck.txt")*/

	cards := newDeckFromFile("my_deck.txt")
	cards.shuffle()
	cards.print()
}
