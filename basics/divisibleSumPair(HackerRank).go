// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem

package main

import "fmt"

func isDivisiblePair(num1, num2, k int) bool {
	return (num1 + num2) % k == 0
}

func countDivisiblePairs(arr []int, n, k int) int {
	pairCount := 0
	
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (isDivisiblePair(arr[i], arr[j], k)) {
				pairCount++
			}
		}
	}
	return pairCount
}

func main() {
	var n, k int
	var arr []int
	fmt.Scan(&n)
	fmt.Scan(&k)

	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	fmt.Println(countDivisiblePairs(arr, n, k))
}