package main

import (
	"fmt"
	"log"
)

func main() {
	input := "-1 2 3 4 5"
	var (
		result        string
		num, max, min int32
		neg           int32 = 1
		length              = len(input) - 1
	)
	StrInt := func(num *int32, char int32) {
		if char >= '0' && char <= '9' {
			*num *= 10         // збільшуємо розряд числа з кожною значущою цифрою
			*num += char - '0' // ASCII code --> 0-9
		} else {
			log.Fatal("невірний формат даних [0-9 ]")
		}

	}
	for i, char := range input {
		switch {
		case i == length: // останнє число у вхідних даних (або перше та єдине)
			StrInt(&num, char)
			fallthrough // gotcha!
		case char == ' ':
			num *= neg // маємо сформоване число зі знаком

			if i == 0 { // TODO: MinMax(). Можливо, буде доречно використати *int32
				min = num
				max = num
			} else {
				if min > num {
					min = num
				}
				if max < num {
					max = num
				}
			}
			num = 0 // обнуляємо лічильники для базового випадку
			neg = 1
		default:
			if char == '-' {
				neg = -1
				continue
			}
			StrInt(&num, char)
		}
	}
	if max != min {
		result = fmt.Sprintf("%v %v", max, min)
	} else {
		result = fmt.Sprintf("%v", max)
	}
	fmt.Println(result)
}
