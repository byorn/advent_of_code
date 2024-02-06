package _2023

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

type BSTLeaf struct {
	value string
	bid   string
	left  *BSTLeaf
	right *BSTLeaf
}

func (root *BSTLeaf) Insert(value string, bid string) {
	if root.value == "" {
		root.value = value
		root.bid = bid
	} else if result := CompareHands(value, root.value); result < 0 {
		if root.left == nil {
			root.left = &BSTLeaf{value: value, bid: bid}
		} else {
			root.left.Insert(value, bid)
		}
	} else if result := CompareHands(value, root.value); result > 0 {
		if root.right == nil {
			root.right = &BSTLeaf{value: value, bid: bid}
		} else {
			root.right.Insert(value, bid)
		}
	}
}

func (root *BSTLeaf) InsertPart2(value string, bid string) {
	if root.value == "" {
		root.value = value
		root.bid = bid
	} else if result := CompareHands(value, root.value, true); result < 0 {
		if root.left == nil {
			root.left = &BSTLeaf{value: value, bid: bid}
		} else {
			root.left.InsertPart2(value, bid)
		}
	} else if result := CompareHands(value, root.value, true); result > 0 {
		if root.right == nil {
			root.right = &BSTLeaf{value: value, bid: bid}
		} else {
			root.right.InsertPart2(value, bid)
		}
	}
}

func (root *BSTLeaf) TraverseInOrder() []string {
	strArr := []string{}
	if root.left != nil {
		strArr = append(strArr, root.left.TraverseInOrder()...)
	}

	strArr = append(strArr, root.bid)

	if root.right != nil {
		strArr = append(strArr, root.right.TraverseInOrder()...)
	}

	return strArr
}

func Day7(inputFile string) (int, int) {
	inputText := util.ReadFile(inputFile)
	fmt.Printf("%d \n", len(inputText))
	return day7Part1(inputText), day7Part2(inputText)
}

func ReplaceJoker(hand string) string {
	if strings.ContainsAny(hand, "J") {
		cardCount := map[string]int{}

		for _, c := range hand {
			cardCount[string(c)] = cardCount[string(c)] + 1
		}

		keyValuesSorted := SortMapByValuesAndKey(cardCount)

		maxCard := keyValuesSorted[0].Key

		if maxCard != "J" {
			return strings.ReplaceAll(hand, "J", maxCard)
		} else {
			if len(keyValuesSorted) > 1 {
				nextMaxCard := keyValuesSorted[1].Key
				return strings.ReplaceAll(hand, "J", nextMaxCard)
			} else {
				return strings.ReplaceAll(hand, "J", "A")
			}
		}

	}
	return hand
}

func CompareHands(hand1, hand2 string, isJJoker ...bool) int {
	origHand1 := hand1
	origHand2 := hand2
	if len(isJJoker) > 0 && isJJoker[0] {
		hand1 = ReplaceJoker(hand1)
		hand2 = ReplaceJoker(hand2)
	}
	if GetHandType(hand1) > GetHandType(hand2) {
		return 1
	} else if GetHandType(hand1) == GetHandType(hand2) {
		if len(isJJoker) > 0 && isJJoker[0] {
			return CompareByFirstCardStrength(origHand1, origHand2, true)
		} else {
			return CompareByFirstCardStrength(hand1, hand2)
		}
	} else {
		return -1
	}
}

func CompareByFirstCardStrength(hand1, hand2 string, isJJoker ...bool) int {
	valJ := 11

	if len(isJJoker) > 0 && isJJoker[0] {
		valJ = 1
	}

	cardStrength := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"J": valJ,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
	}

	for i := 0; i < len(hand1); i++ {
		//
		if cardStrength[hand1[i:i+1]] > cardStrength[hand2[i:i+1]] {
			return 1
		} else if cardStrength[hand1[i:i+1]] < cardStrength[hand2[i:i+1]] {
			return -1
		} else {
			continue
		}
	}
	return 0
}

func GetHandType(hand string) int {
	handTypeRank := 0

	handType := map[string]int{}

	for _, c := range hand {
		handType[string(c)] = handType[string(c)] + 1
	}

	keyValuesSorted := util.SortMapByValues(handType)

	handTypeStr := ""

	for _, kv := range keyValuesSorted {
		handTypeStr = handTypeStr + strconv.Itoa(kv.Value)
	}

	switch handTypeStr {
	case "5":
		{
			handTypeRank = 7
			break
		}
	case "41":
		{
			handTypeRank = 6
			break
		}
	case "32":
		{
			handTypeRank = 5
			break
		}
	case "311":
		{
			handTypeRank = 4
			break
		}
	case "221":
		{
			handTypeRank = 3
			break
		}
	case "2111":
		{
			handTypeRank = 2
			break
		}
	case "11111":
		{

			handTypeRank = 1
			break
		}
	}
	return handTypeRank
}

/*
Return the max total from all group totals
*/
func day7Part1(input []string) int {
	root := BSTLeaf{value: strings.Fields(input[0])[0], bid: strings.Fields(input[0])[1]}
	input1 := input[1:]
	for _, line := range input1 {
		card, bid := strings.Fields(line)[0], strings.Fields(line)[1]
		root.Insert(card, bid)
	}

	cardsSortedAsc := root.TraverseInOrder()
	sum := 0
	for i, bid := range cardsSortedAsc {
		bidInt, _ := strconv.Atoi(bid)
		sum += (i + 1) * bidInt
	}
	return sum
}

func day7Part2(input []string) int {
	root := BSTLeaf{value: strings.Fields(input[0])[0], bid: strings.Fields(input[0])[1]}
	input1 := input[1:]
	for _, line := range input1 {
		card, bid := strings.Fields(line)[0], strings.Fields(line)[1]
		root.InsertPart2(card, bid)
	}

	cardsSortedAsc := root.TraverseInOrder()
	sum := 0
	for i, bid := range cardsSortedAsc {
		bidInt, _ := strconv.Atoi(bid)
		sum += (i + 1) * bidInt
	}
	return sum
}

func SortMapByValuesAndKey(mymap map[string]int) []util.KeyValue {
	keyValues := []util.KeyValue{}

	cardMap := map[string]int{
		"A": 14,
		"K": 13,
		"Q": 12,
		"T": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"J": 1,
	}

	for Key, Value := range mymap {
		keyValues = append(keyValues,
			util.KeyValue{Key: Key, Value: Value})
	}

	sort.Slice(keyValues, func(i, j int) bool {
		if keyValues[i].Value > keyValues[j].Value {
			return true
		}
		if keyValues[i].Value < keyValues[j].Value {
			return false
		}
		if keyValues[i].Value == keyValues[j].Value {
			return cardMap[keyValues[i].Key] > cardMap[keyValues[j].Key]
		}
		return false
	})
	return keyValues
}
