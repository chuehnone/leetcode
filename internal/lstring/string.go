package lstring

import (
	"fmt"
	"strconv"
)

func BalancedStringSplit(s string) int {
	cnt := 0

	sLen := len(s)
	var last string
	push := 0
	for i := 0; i < sLen; i++ {
		ch := string(s[i])
		if push == 0 {
			last = ch
			cnt++
		}

		if ch == last {
			push++
		} else {
			push--
		}
	}

	return cnt
}

func ToHexspeak(num string) string {
	const returnErr = "ERROR"
	dec, err := strconv.Atoi(num)
	if err != nil {
		return returnErr
	}

	hex := fmt.Sprintf("%X", dec)

	sLen := len(hex)
	ans := ""
	for i := 0; i < sLen; i++ {
		if hex[i] == '0' {
			ans += "O"
		} else if hex[i] == '1' {
			ans += "I"
		} else if hex[i] >= 'A' && hex[i] <= 'F' {
			ans += string(hex[i])
		} else {
			return returnErr
		}
	}

	return ans
}

func MaxNumberOfBalloons(text string) int {
	balloonTable := map[byte]int{
		'b': 0,
		'a': 0,
		'l': 0,
		'o': 0,
		'n': 0,
	}

	textLen := len(text)
	for i := 0 ; i < textLen ; i++ {
		c := text[i]
		if _, exist := balloonTable[c]; exist {
			balloonTable[c]++
		}
	}

	min := balloonTable['b']
	for key, value := range balloonTable {
		if key == 'l' || key == 'o' {
			value = value / 2
		}
		if min > value {
			min = value
		}
	}

	return min
}
