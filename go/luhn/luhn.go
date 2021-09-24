package luhn

import (
	"strconv"
	"strings"
)

func Valid(credit_card string) bool {

	numbers := strings.Replace(credit_card, " ", "", -1)

	if len(numbers) <= 1 {
		return false
	}

	x := 1
	sum := 0
	for i := len(numbers) - 1; i >= 0; i-- {
		if value, err := strconv.Atoi(string(numbers[i])); err == nil {
			if x%2 == 0 && err == nil {
				if value*2 > 9 {
					sum += value*2 - 9
				} else {
					sum += value * 2
				}
			} else {
				sum += value
			}
			x += 1
		} else {
			return false
		}
	}

	if sum%10 != 0 {
		return false
	}
	return true
}
