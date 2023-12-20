package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mpCardBid := map[string]int64{}
	for scanner.Scan() {
		line := scanner.Text()
		mappings := strings.Split(line, " ")
		mpCardBid[mappings[0]] = getInt(mappings[1])
	}

	mpRankCard := map[int64][]string{}
	for card := range mpCardBid {
		mpLabelCount := make(map[rune]int)
		for _, label := range card {
			mpLabelCount[label]++
		}

		switch len(mpLabelCount) {
		case 1:
			// five of a kind
			mpRankCard[7] = append(mpRankCard[7], card)
		case 2:
			diff := 0
			for _, val := range mpLabelCount {
				if diff == 0 {
					diff += val
				} else {
					diff -= val
				}
			}
			// four of a kind 41 diff=3
			if diff == -3 || diff == 3 {
				mpRankCard[6] = append(mpRankCard[6], card)
			}

			// full house 32 diff=1
			if diff == -1 || diff == 1 {
				mpRankCard[5] = append(mpRankCard[5], card)
			}
		case 3:
			mul := 1
			for _, val := range mpLabelCount {
				mul *= val
			}
			// three of a kind 311 mul=3
			if mul == 3 {
				mpRankCard[4] = append(mpRankCard[4], card)
			}
			// two pair 221 mul=4
			if mul == 4 {
				mpRankCard[3] = append(mpRankCard[3], card)
			}
		case 4:
			// one pair
			mpRankCard[2] = append(mpRankCard[2], card)
		case 5:
			// high card
			mpRankCard[1] = append(mpRankCard[1], card)
		}
	}

	rankCardSlice := [][]string{}
	score := int64(0)
	rank := int64(1)

	keys := make([]int64, 0)
	for k := range mpRankCard {
		keys = append(keys, k)
	}
	slices.Sort(keys)

	for _, key := range keys {
		cards := mpRankCard[key]
		sort.Slice(cards, func(i, j int) bool { return cardCondition(cards[i], cards[j]) })
		rankCardSlice = append(rankCardSlice, cards)
		for _, card := range cards {
			score += int64(rank * int64(mpCardBid[card]))
			rank++
		}
	}

	fmt.Println(rankCardSlice)
	// fmt.Println(mpRankCard)
	fmt.Println(score)

	/*
	   Five of a kind: five cards same label: AAAAA
	   Four of a kind: four cards same label and one card different: AA8AA
	   Full house: three cards same label, remaining two cards different label: 23332
	   Three of a kind: three cards same label, remaining two cards each different from any other card: TTT98
	   Two pair: two cards one label, two other cards second label, remaining card third label: 23432
	   One pair: two cards one label, other three cards different label from the pair and each other: A23A4
	   High card: all cards' labels are distinct: 23456
	*/
}

func cardCondition(card1 string, card2 string) bool {
	ranking := map[rune]int{}
	ranking['A'] = 13
	ranking['K'] = 12
	ranking['Q'] = 11
	ranking['J'] = 10
	ranking['T'] = 9
	ranking['9'] = 8
	ranking['8'] = 7
	ranking['7'] = 6
	ranking['6'] = 5
	ranking['5'] = 4
	ranking['4'] = 3
	ranking['3'] = 2
	ranking['2'] = 1

	for i := 0; i < 5; i++ {
		if card1[i] != card2[i] {
			return ranking[rune(card1[i])] < ranking[rune(card2[i])]
		}
	}

	return true
}

func getInt(str string) int64 {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return int64(val)
}
