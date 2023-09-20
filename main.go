package main

import (
	"fmt"
	"strconv"
)

var number_1 string
var number_2 string
var operation string

var roman = map[string]int{
	"C":    100,
	"XC":   90,
	"L":    50,
	"XL":   40,
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1,
}
var roman_2 = map[int]string{
	100: "C",
	90:  "XC",
	50:  "L",
	40:  "XL",
	10:  "X",
	9:   "IX",
	8:   "VIII",
	7:   "VII",
	6:   "VI",
	5:   "V",
	4:   "IV",
	3:   "III",
	2:   "II",
	1:   "I",
}
var convIntToRoman = [14]int{
	100,
	90,
	50,
	40,
	10,
	9,
	8,
	7,
	6,
	5,
	4,
	3,
	2,
	1,
}

func main() {

	fmt.Print("Введите пример: ")
	fmt.Scan(&number_1, &operation, &number_2)

	num_1, exists := roman[number_1]
	num_2, exist := roman[number_2]

	if exists {
		if exist {
			if num_1 <= 10 && num_2 <= 10 {
				roman_calc(num_1, operation, num_2)
			} else {
				fmt.Println("Одно из введенных чисел больше X")
			}
		} else {
			fmt.Println("Оба числа должны быть в арабской или римской системе счисления.")
		}
	} else if exist {

		fmt.Println("Оба числа должны быть в арабской или римской системе счисления.")

	} else {
		number_1, err := strconv.Atoi(number_1)
		if err != nil {
			panic(err)
		}
		number_2, err_2 := strconv.Atoi(number_2)
		if err_2 != nil {
			panic(err_2)
		}

		if number_1 <= 10 && number_2 <= 10 {
			arabic_calc(number_1, operation, number_2)
		} else {
			fmt.Println("Одно из введенных чисел больше 10")
		}
	}
}

func arabic_calc(number_1 int, operation string, number_2 int) {

	if operation == "+" {
		sum := number_1 + number_2
		fmt.Println(sum)
	} else if operation == "/" {
		sum := number_1 / number_2
		fmt.Println(sum)
	} else if operation == "*" {
		sum := number_1 * number_2
		fmt.Println(sum)
	} else if operation == "-" {
		sum := number_1 - number_2
		fmt.Println(sum)
	} else {
		fmt.Println("Такой математической операции нет в программе...")
	}

}

func roman_calc(num_1 int, operation string, num_2 int) {

	if operation == "+" {
		sum := num_1 + num_2
		convToRoman(sum)
	} else if operation == "/" {
		sum := num_1 / num_2
		convToRoman(sum)
	} else if operation == "*" {
		sum := num_1 * num_2
		convToRoman(sum)
	} else if operation == "-" {
		sum := num_1 - num_2
		convToRoman(sum)
	} else {
		fmt.Println("Такой математической операции нет в программе...")
	}

}

func convToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		fmt.Println("В римской системе счисления нет числа 0!")
	} else if romanResult < 0 {
		fmt.Println("В римской системе счисления нет отрицательных чисел!")
	} else {
		for romanResult > 0 {
			for _, elem := range convIntToRoman {
				for i := elem; i <= romanResult; {
					if value, ok := roman_2[elem]; ok {
						romanNum += value
						romanResult -= elem
					}
				}
			}
		}
		fmt.Println(romanNum)
	}
}
