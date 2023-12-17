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

	classifiedHands := classifyHands(camelHands, false)
	fmt.Printf("Total Winnings are: %d\n", calculateWinnings(classifiedHands, len(camelHands)))
}

func Part2Solve(input string) {

	file, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
	}

	stringInput := string(file)

	camelHands := getHands(stringInput)

	classifiedHands := classifyHands(camelHands, true)

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

func classifyHands(hands []handOfCamel, part2 bool) map[string][]handOfCamel {
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

			if part2 {

				if strings.Contains(hand.hand, "J") {

					// First find all the cases where jokers make a hand

					// NOTE:
					//        - Its impossible to have 2 pairs if there is a joker in the hand
					//		  - Its impossible to have a High Card with a joker in the hand

					// If there are 4 jokers in a hand then the only hand we can make is 5 of a kind and we are done
					if strings.Count(hand.hand, "J") == 4 || strings.Count(hand.hand, "J") == 5 {
						classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand, part2)
						found3k = false
						foundPair = false
						highcard = false
						break
					}

					// If there are 3 jokers in a hand then the minimum hand we can make is 4 of a kind
					// If the other 2 cards are a pair then we have five of a kind and we are done
					// Otherwise its 4 of a kind and we are done
					if strings.Count(hand.hand, "J") == 3 {

						for _, xcard := range hand.hand {

							if xcard != 'J' {
								if strings.Count(hand.hand, string(xcard)) == 2 {
									classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand, part2)

									break
								} else {
									classifiedHands["4K"] = insertInRankedOrder(classifiedHands["4K"], hand, part2)

									break
								}

							}
						}

					}
					// If there are 2 jokers in a hand then the minimum had we can make is 3 of a kind
					// If the other 3 cards contain contain a pair then we have four of a kind and we are done
					// If the other 3 cards are three of a kind then we have five of a kind and we are done

					if strings.Count(hand.hand, "J") == 2 {
						for _, xcard := range hand.hand {
							if xcard != 'J' {
								if strings.Count(hand.hand, string(xcard)) == 3 {
									classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand, part2)
									break

								} else if strings.Count(hand.hand, string(xcard)) == 2 {
									classifiedHands["4K"] = insertInRankedOrder(classifiedHands["4K"], hand, part2)
									break

								} else if strings.Count(hand.hand, string(xcard)) == 1 {
									// Check if the characters apart from the jokers are unique
									// if they are we have 3 of a kind. Otherwise continue the loop

									if checkUnique(strings.Replace(hand.hand, "J", "", -1)) {
										classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand, part2)
										break
									}

									continue

								} else {
									classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand, part2)
									break
								}
							}
						}
					}

					// If there is 1 joker in a hand then the minimum hand we can make is a pair
					// If all the remaing cards are the same then we have 5 of a kind
					if strings.Count(hand.hand, "J") == 1 {
						//pairs := []string{}

						for _, xcard := range hand.hand {
							if xcard != 'J' {
								if strings.Count(hand.hand, string(xcard)) == 4 {
									classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand, part2)
									break

								} else if strings.Count(hand.hand, string(xcard)) == 3 {
									classifiedHands["4K"] = insertInRankedOrder(classifiedHands["4K"], hand, part2)
									break
									// There are 4 cards in the hand to check. They could either be all different
									// which means we make a pair or they are 2 pair which means we make a full house
								} else if strings.Count(hand.hand, string(xcard)) == 2 {

									unique, uniqueCount := uniqueCount(strings.Replace(hand.hand, "J", "", -1))

									if unique {
										classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand, part2)

										break
									} else {
										// two of the three remaining cards are unique so we have two pair
										if uniqueCount == 2 {
											classifiedHands["FH"] = insertInRankedOrder(classifiedHands["FH"], hand, part2)

											break
										} else if uniqueCount == 3 {
											// Just a pair
											classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand, part2)

											break
										}
									}

								} else if strings.Count(hand.hand, string(xcard)) == 1 {
									if checkUnique(strings.Replace(hand.hand, "J", "", -1)) {
										classifiedHands["1P"] = insertInRankedOrder(classifiedHands["1P"], hand, part2)
										break
									}

								}
							}
						}
					}
					// We should break out of the main loop here as well, since we have completely classified
					// every hand with a joker in it
					found3k = false
					foundPair = false
					highcard = false

					break
				}

			}

			// The rest of the hands wont contain a joker, so its the same as Part 1 but with a different
			// calculation to cater for the joker value

			// Four and Five of a kind are easy
			if c == 4 {
				// insert the card in the right place in the array terms of first card value

				classifiedHands["4K"] = insertInRankedOrder(classifiedHands["4K"], hand, part2)

				found3k = false
				foundPair = false
				highcard = false
				break
			}

			// five of a kind is always five of a kind. No need to transform for part 2
			if c == 5 {
				classifiedHands["5K"] = insertInRankedOrder(classifiedHands["5K"], hand, part2)
				found3k = false
				foundPair = false
				highcard = false
				break
			}

			if c == 2 && foundPair && !slices.Contains(pairCards, string(card)) {

				// two pair
				classifiedHands["2P"] = insertInRankedOrder(classifiedHands["2P"], hand, part2)

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

				// A full house can transform into 4 of a kind if the remaining card is a joker

				classifiedHands["FH"] = insertInRankedOrder(classifiedHands["FH"], hand, part2)

				found3k = false
				foundPair = false
				highcard = false
				break
			}

			// If we find 3 of a kind and we have already found a pair, then the hand is a full house

			// It's only possible to have one 3 of a kind in the hand. So we dont need to track which cards
			// are part of the 3 of a kind
			if c == 3 && !found3k && !foundPair {

				// Three of a kind can transform into a 4 of a kind if the remaining card is a joker
				// If the other 2 cards are jokers then we get 4 of a kind

				found3k = true
				highcard = false

				// If we find 3 of a kind and we have already found a pair then we have a full house
			} else if c == 3 && foundPair {
				// A full hose can transform into a 4 of a kind if the remaining card is a joker

				classifiedHands["FH"] = insertInRankedOrder(classifiedHands["FH"], hand, part2)

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
			// If one of the high cards is a joker then we can transform the hand into a pair

			classifiedHands["HC"] = insertInRankedOrder(classifiedHands["HC"], hand, part2)

		}

		if foundPair {
			// if one of the high cards is a joker then we can transform the hand into three of a kind

			classifiedHands["1P"] = insertInRankedOrder(classifiedHands["1P"], hand, part2)

		} else if found3k {
			// if one of the other cards is a joker then we can transform the hand into 4 of a kind
			// if there are 2 jokers in our hand then that is a full house.

			classifiedHands["3K"] = insertInRankedOrder(classifiedHands["3K"], hand, part2)

		}

	}

	return classifiedHands
}
func insertInRankedOrder(hands []handOfCamel, hand handOfCamel, part2 bool) (ranked []handOfCamel) {

	if len(hands) == 0 {
		ranked = append(ranked, hand)
		return ranked

	} else {
		for i, existingHand := range hands {
			if isHigherRank(hand, existingHand, part2) {
				// insert hand before existing hand
				ranked := hands
				if !slices.Contains(ranked, hand) {
					ranked = append(ranked, handOfCamel{})
					copy(ranked[i+1:], ranked[i:])
					ranked[i] = hand
				}

				return ranked
			}
		}
		// we dont return from the above loop so the new hand must be the lowest rank. So simply append it
		if !slices.Contains(ranked, hand) {

			ranked = append(hands, hand)
		}
	}

	return ranked
}

func checkUnique(checkString string) bool {
	allUnique := true

	for _, char := range checkString {
		count := strings.Count(checkString, string(char))
		if count > 1 {
			allUnique = false
			break
		}
	}

	return allUnique
}

func uniqueCount(checkString string) (bool, int) {
	allUnique := true
	uniqueCount := 0
	for i, char := range checkString {

		count := strings.Count(checkString, string(char))
		if count > 1 {
			allUnique = false
		}

		if strings.Contains(checkString[i+1:], string(char)) {
			allUnique = false

		} else {
			uniqueCount++
		}
	}

	return allUnique, uniqueCount
}

func isHigherRank(hand1 handOfCamel, hand2 handOfCamel, part2 bool) bool {
	isHigher := false

	for idx, card := range hand1.hand {
		card1Val, err := getCardValue(card, part2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		card2Val, err := getCardValue(rune(hand2.hand[idx]), part2)
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
func getCardValue(card rune, part2 bool) (int, error) {

	cardValue := 0

	courtVals := map[string]int{"T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}
	if part2 {
		courtVals = map[string]int{"T": 10, "J": 1, "Q": 12, "K": 13, "A": 14}
	}

	var err error

	if unicode.IsDigit(card) {
		cardValue, err = strconv.Atoi(string(card))

	} else {
		cardValue = courtVals[string(card)]
	}

	return cardValue, err
}
