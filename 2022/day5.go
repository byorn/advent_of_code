package _2022

import (
	"strconv"
	"strings"

	util "github.com/byorn/advent_of_code/util"
)

func Day5(inputFile string) string {
	inputText := util.ReadFile(inputFile)

	return getPeekFromAllStacks(inputText)
}

func getPeekFromAllStacks(entries []string) string {
	stack1 := util.NewStack([]util.Item{"B", "W", "N"})
	stack2 := util.NewStack([]util.Item{"L", "Z", "S", "P", "T", "D", "M", "B"})
	stack3 := util.NewStack([]util.Item{"Q", "H", "Z", "W", "R"})
	stack4 := util.NewStack([]util.Item{"W", "D", "V", "J", "Z", "R"})
	stack5 := util.NewStack([]util.Item{"S", "H", "M", "B"})
	stack6 := util.NewStack([]util.Item{"L", "G", "N", "J", "H", "V", "P", "B"})
	stack7 := util.NewStack([]util.Item{"J", "Q", "Z", "F", "H", "D", "L", "S"})
	stack8 := util.NewStack([]util.Item{"W", "S", "F", "J", "G", "Q", "B"})
	stack9 := util.NewStack([]util.Item{"Z", "W", "M", "S", "C", "D", "J"})

	cratesMap := map[int]util.Stack{
		1: stack1,
		2: stack2,
		3: stack3,
		4: stack4,
		5: stack5,
		6: stack6,
		7: stack7,
		8: stack8,
		9: stack9,
	}

	// fmt.Println(cratesMap)

	for i, entry := range entries {
		if i >= 10 {
			str1 := strings.Replace(entry, "move", "", -1)
			str2 := strings.Replace(str1, "from", "", -1)
			str3 := strings.Replace(str2, "to", "", -1)
			str4 := strings.ReplaceAll(str3, " ", "")

			to, _ := strconv.Atoi(string(str4[len(str4)-1]))
			from, _ := strconv.Atoi(string(str4[len(str4)-2]))
			amount, _ := strconv.Atoi(str4[:len(str4)-2])

			for k := 0; k < amount; k++ {

				var item util.Item

				if stackFrom, ok := cratesMap[from]; ok {
					item, _ = stackFrom.Pop()
					cratesMap[from] = stackFrom
				}

				if stackTo, ok := cratesMap[to]; ok {

					stackTo.Push(item)
					cratesMap[to] = stackTo
				}
			}
		}
	}
	var a1, a2, a3, a4, a5, a6, a7, a8, a9 string
	if stack1, ok := cratesMap[1]; ok {
		answer1, _ := stack1.Peek()
		a1, _ = answer1.(string)

	}

	if stack2, ok := cratesMap[2]; ok {
		answer2, _ := stack2.Peek()
		a2, _ = answer2.(string)
	}

	if stack3, ok := cratesMap[3]; ok {
		answer3, _ := stack3.Peek()
		a3, _ = answer3.(string)
	}

	if stack4, ok := cratesMap[4]; ok {
		answer4, _ := stack4.Peek()
		a4, _ = answer4.(string)
	}

	if stack5, ok := cratesMap[5]; ok {
		answer5, _ := stack5.Peek()
		a5, _ = answer5.(string)
	}

	if stack6, ok := cratesMap[6]; ok {
		answer6, _ := stack6.Peek()
		a6, _ = answer6.(string)
	}

	if stack7, ok := cratesMap[7]; ok {
		answer7, _ := stack7.Peek()
		a7, _ = answer7.(string)
	}

	if stack8, ok := cratesMap[8]; ok {
		answer8, _ := stack8.Peek()
		a8, _ = answer8.(string)
	}

	if stack9, ok := cratesMap[9]; ok {
		answer9, _ := stack9.Peek()
		a9, _ = answer9.(string)
	}

	return a1 + a2 + a3 + a4 + a5 + a6 + a7 + a8 + a9
}
