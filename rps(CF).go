// https://codeforces.com/contest/1380/problem/B

package main

import (
	"fmt"
	"strings"
)

func findOptimalSequence(RPSseq string) string {
	lenOfSeq := len(RPSseq)

	rockCount := strings.Count(RPSseq, "R")
	paperCount := strings.Count(RPSseq, "P")
	scissorCount := strings.Count(RPSseq, "S")
	
	if rockCount > paperCount && rockCount > scissorCount {
		return strings.Repeat("P", lenOfSeq)
	} else if paperCount > scissorCount {
		return strings.Repeat("S", lenOfSeq)
	} else {
		return strings.Repeat("R", lenOfSeq)
	}
}

func main() {
	var testCaseCount int
	fmt.Scanln(&testCaseCount)

	for tc := 0; tc < testCaseCount; tc++ {
		var RPSseq string
		fmt.Scanln(&RPSseq)
		fmt.Println(findOptimalSequence(RPSseq))
	}
}
