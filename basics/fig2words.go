package main

import (
	"fmt"
	"strings"
)

var H100 string = "Hundred "

func convert2digits(n int) string {
	var upto20 = [...]string{"", "One ", "Two", "Three ", "Four ", "Five ", "Six ", "Seven ", "Eight ", "Nine ", "Ten ",
					         "Eleven ", "Twelve ", "Thirteen ", "Fourteen ", "Fifteen ", "Sixteen ", "Seventeen ", "Eighteen ", "Nineteen "}
	var tens = [...]string{"", "Ten ", "Twenty ", "Thirty ", "Forty ", "Fifty ", "Sixty ", "Seventy ", "Eighty ", "Ninety "}

	if (n < len(upto20)) {
		return upto20[n]
	}
	return tens[n / 10] + upto20[n % 10]
}

func convert3digits(n int) string {
	if (n < 100) {
		return convert2digits(n)
	}
	return fix_and(convert2digits(n / 100) + H100 + convert2digits(n % 100))
}

func fix_and(s string) string {
	if (strings.Contains(s, H100) && !strings.HasSuffix(s, H100)) {
		s = strings.ReplaceAll(s, H100, H100 + "and ")
	}
	return s
}

func fig2words(n int) string {
	if (n == 0) {
		return "Zero"
	}
	return convert3digits(n)
}

func main() {
	var num int
	fmt.Scanln(&num)
	fmt.Println(fig2words(num))
}
