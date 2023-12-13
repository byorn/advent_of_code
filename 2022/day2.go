/*
For example, suppose you were given the following strategy guide:

A Y
B X
C Z
This strategy guide predicts and recommends the following:

In the first round, your opponent will choose Rock (A), and you should choose Paper (Y). This ends in a win for you with a score of 8 (2 because you chose Paper + 6 because you won).
In the second round, your opponent will choose Paper (B), and you should choose Rock (X). This ends in a loss for you with a score of 1 (1 + 0).
The third round is a draw with both players choosing Scissors, giving you a score of 3 + 3 = 6.
In this example, if you were to follow the strategy guide, you would get a total score of 15 (8 + 1 + 6).
*/

/*
X - 1
Y - 2
Z - 3

won - 6
draw - 3
lose - 0
*/

package _2022

import (
	"fmt"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

var gameResult = map[string]string{
	"AX": "D",
	"AY": "W",
	"AZ": "L",
	"BX": "L",
	"BY": "D",
	"BZ": "W",
	"CX": "W",
	"CY": "L",
	"CZ": "D",
}

var gameResultPart2 = map[string][]string{
	"AX": {"L", "Z"},
	"AY": {"D", "X"},
	"AZ": {"W", "Y"},
	"BX": {"L", "X"},
	"BY": {"D", "Y"},
	"BZ": {"W", "Z"},
	"CX": {"L", "Y"},
	"CY": {"D", "Z"},
	"CZ": {"W", "X"},
}

var gameResultPoints = map[string]int{
	"W": 6,
	"D": 3,
	"L": 0,
}

var symbolPoints = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

func Day2_Part1(inputFile string) int {
	inputText := util.ReadFile(inputFile)
	fmt.Println("okay testing neovim go")
	return getTotalScore(inputText)
}

func Day2_Part2(inputFile string) int {
	inputText := util.ReadFile(inputFile)
	return getTotalScorePart2(inputText)
}

func getTotalScore(scores []string) int {
	symbolTotalScore := 0
	gameResultTotalScore := 0
	for _, score := range scores {

		play := strings.Replace(score, " ", "", -1)
		gameResult, symbol := getGameResultAndSymbol(play)
		gameResultTotalScore += gameResult
		symbolTotalScore += symbol

	}

	return symbolTotalScore + gameResultTotalScore
}

func getTotalScorePart2(scores []string) int {
	symbolTotalScore := 0
	gameResultTotalScore := 0
	for _, score := range scores {

		play := strings.Replace(score, " ", "", -1)
		play = play[0:1] + gameResultPart2[play][1]
		gameResult, symbol := getGameResultAndSymbol(play)
		gameResultTotalScore += gameResult
		symbolTotalScore += symbol

	}

	return symbolTotalScore + gameResultTotalScore
}

func getGameResultAndSymbol(play string) (int, int) {
	result := gameResult[play]
	gameResultPoint := gameResultPoints[result]
	symbolPoint := symbolPoints[play[1:]]
	return gameResultPoint, symbolPoint
}
