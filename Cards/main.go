package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"

	card := "Ace of Spades" // compiler determines string is going to be assigned to it. We only use := when defining a new variable.

	card = "Five of Diamonds"

	fmt.Println(card)

	new_card := newCard()

	fmt.Println(new_card)

	// arrays && slices

	cards := []string{"Ace of Diamonds", newCard()}

	fmt.Println(cards)

	cards = append(cards, "Six of Spades")

	fmt.Println(cards)

	for i := range cards {
		fmt.Println(i, cards[i])
	}

	my_cards := deck{"Ace of Diamonds", newCard()}

	fmt.Println(my_cards)

	my_cards.print()

	my_cards = newDeck()

	my_cards.print()

	hand, remainingDeck := deal(my_cards, 5)

	hand.print()

	remainingDeck.print()

	// converting string to byte slice
	greeting := "Hi there!"
	fmt.Println([]byte(greeting))

	fmt.Println(my_cards.toString())

	my_cards.saveToFile("my_cards")

	my_cards = newDeckFromFile("my_cards")

	my_cards.shuffle()

	my_cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
