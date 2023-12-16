package puzzle7

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

type handOfCamel struct {
	hand string
	bid  int
}

func Part1Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	camelHands := getHands(stringInput)

	classifiedHands := classifyHands(camelHands)

	fmt.Printf("Total Winnings are: %d\n", calculateWinnings(classifiedHands, len(camelHands)))
}

// Need to know how many hands are being ranked up front so we can apply the rank multipler
func calculateWinnings(hands map[string][]handOfCamel, totalHands int) (winnings int) {
	// Hand order:
	// High Card  = HC
	// Pair = 1P
	// Two Pair = 2P
	// Three of a Kind = 3K
	// Full House = FH
	// Four of a Kind = 4K
	// Five of a Kind = 5K

	// Calcuate winnings based on ranking of hands with FH being highest and HC being lowest
	currentRank := totalHands

	// Count hand rank order

	for _, hand := range hands["5K"] {
		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["4K"] {

		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["FH"] {

		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["3K"] {
		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["2P"] {
		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["1P"] {

		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	for _, hand := range hands["HC"] {

		winnings = winnings + hand.bid*currentRank
		currentRank--
	}

	return
}

func getHands(input string) (hands []handOfCamel) {

	scanner := bufio.NewScanner(strings.NewReader(input))

	for scanner.Scan() {
		line := scanner.Text()

		cards := line[:5]
		bid, _ := strconv.Atoi(line[6:])

		hand := handOfCamel{cards, bid}
		hands = append(hands, hand)

	}

	return

}

func classifyHands(hands []handOfCamel) map[string][]handOfCamel {
	// High Card  = HC
	// Pair = 1P
	// Two Pair = 2P
	// Three of a Kind = 3K
	// Full House = FH
	// Four of a Kind = 4K
	// Five of a Kind = 5K

	classifiedHands := make(map[string][]handOfCamel)

	for _, hand := range hands {

		foundPair := false
		pairCards := []string{}

		found3k := false
		highcard := true

		for _, card := range hand.hand {

			c := strings.Count(hand.hand, string(card))

			// Four and Five of a kind are easy
			if c == 4 {
				// insert the card in the right place in the array terms of first card value
				classifiedHands["4K"] = insertInRankedOrder(classifiedHands["4K"], hand)
				found3k = false
				foundPair = false
				highcard = false
				break
			}

			if c == 5 {
				classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand)
				found3k = false
				foundPair = false
				highcard = false
				break
			}

			// If we find a pair and its different from a pair we have already found
			// then we have 2 pair
			if c == 2 && foundPair && !slices.Contains(pairCards, string(card)) {
				// two pair
				classifiedHands["2P"] = insertInRankedOrder(classifiedHands["2P"], hand)
				found3k = false
				foundPair = false
				highcard = false
				break

				// if we find a pair and havent already found a pair then at minimum the hand is a pair
				// but could still be a full house or two pair. So we dont rank it yet.
			} else if c == 2 && !foundPair {
				foundPair = true
				highcard = false
				pairCards = append(pairCards, string(card))

				//if we find a pair and we have already found a 3 of a kind then we have a full house
			} else if c == 2 && found3k {
				classifiedHands["FH"] = insertInRankedOrder(classifiedHands["FH"], hand)
				found3k = false
				foundPair = false
				highcard = false
				break
			}

			// T55J5
			// If we find 3 of a kind and we have already found a pair, then the hand is a full house

			// It's only possible to have one 3 of a kind in the hand. So we dont need to track which cards
			// are part of the 3 of a kind
			if c == 3 && !found3k && !foundPair {
				found3k = true
				highcard = false

				// If we find 3 of a kind and we have already found a pair then we have a full house
			} else if c == 3 && foundPair {
				classifiedHands["FH"] = insertInRankedOrder(classifiedHands["FH"], hand)
				found3k = false
				foundPair = false
				highcard = false
				break
			}

		}

		if found3k && foundPair {
			fmt.Println("Found a pair and three of a kind but didnt classify as a full house!")
			fmt.Println("This should never happen and is a BUG!")
			os.Exit(1)
		}

		// If we exit this loop and found pair or found3k are true then we have three of a kind or just a pair
		// or just high card

		if highcard {
			classifiedHands["HC"] = insertInRankedOrder(classifiedHands["HC"], hand)
		}

		if foundPair {
			classifiedHands["1P"] = insertInRankedOrder(classifiedHands["1P"], hand)

		} else if found3k {

			classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand)

		}

	}

	return classifiedHands
}

func insertInRankedOrder(hands []handOfCamel, hand handOfCamel) (ranked []handOfCamel) {

	if len(hands) == 0 {
		ranked = append(ranked, hand)
		return ranked

	} else {
		for i, existingHand := range hands {
			if isHigherRank(hand, existingHand) {
				// insert hand before existing hand
				ranked := hands

				ranked = append(ranked, handOfCamel{})
				copy(ranked[i+1:], ranked[i:])
				ranked[i] = hand

				return ranked
			}
		}
		// we dont return from the above loop so the new hand must be the lowest rank. So simply append it
		ranked = append(hands, hand)
	}

	return ranked
}

func isHigherRank(hand1 handOfCamel, hand2 handOfCamel) bool {
	isHigher := false

	for idx, card := range hand1.hand {
		card1Val, err := getCardValue(card)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		card2Val, err := getCardValue(rune(hand2.hand[idx]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if card1Val == card2Val {
			continue
		}

		if card1Val > card2Val {
			isHigher = true
		}
		break
	}
	return isHigher
}
func getCardValue(card rune) (int, error) {

	cardValue := 0
	courtVals := map[string]int{"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
	var err error

	if unicode.IsDigit(card) {
		cardValue, err = strconv.Atoi(string(card))

	} else {
		cardValue = courtVals[string(card)]
	}

	return cardValue, err
}
