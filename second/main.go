package main

import (
	"fmt"
)

func main() {
	input := "1 9 3 4 -5"
	var (
		start, end    int
		num, max, min int32
		length        = len(input) - 1
		numStr        string
		result        string
	)

	for i, char := range input {
		if char == ' ' || i == length {
			end, start = i+1, end
			numStr = input[start:end]
			num = StrToInt(numStr)

			if i == 0 { // TODO: MaxMin()
				max = num
				min = num
			} else {
				if num > max {
					max = num
				}
				if num < min {
					min = num
				}
			}
		}
	}
	if max != min {
		result = fmt.Sprintf("%v %v", max, min)
	} else {
		result = fmt.Sprintf("%v", max)
	}
	fmt.Println("результат", result)
}

func StrToInt(str string) (num int32) {
	var neg int32 = 1
	for _, char := range str {
		if char == '-' {
			neg = -1
			continue
		}
		if char >= '0' && char <= '9' {
			num *= 10         // збільшуємо розряд числа з кожною значущою цифрою
			num += char - '0' // ASCII --> 0-9
		}
	}
	num *= neg
	return
}
