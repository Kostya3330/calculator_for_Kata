package main

import (
	"fmt"
	"strconv"
)

func main() {
	romeNumMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

	var znak, a, b, c string
	fmt.Scanf("%s %s %s %s", &a, &znak, &b, &c)
	if c != "" {
		panic("Введено больше двух операндов")
	}
	
	rightOper(znak)             // проверяем на знак
	if arabOrRome(a, b, znak) { // проверяем ввод операндов
		cloneA, _ := strconv.Atoi(a)
		cloneB, _ := strconv.Atoi(b)

		switch znak {
		case "+":
			fmt.Println(cloneA + cloneB)
		case "-":
			fmt.Println(cloneA - cloneB)
		case "*":
			fmt.Println(cloneA * cloneB)
		case "/":
			fmt.Println(cloneA / cloneB)

		}
	} else {
		var arabNum int
		switch znak {
		case "+":
			arabNum = romeNumMap[a] + romeNumMap[b]
		case "-":
			arabNum = romeNumMap[a] - romeNumMap[b]
		case "*":
			arabNum = romeNumMap[a] * romeNumMap[b]
		case "/":
			arabNum = romeNumMap[a] / romeNumMap[b]
		}
		if arabNum < 1 {
			panic("Результат операции над римскими числами меньше единицы ")
		}
		fmt.Println(convertToRoman(arabNum))
	}

}
func arabOrRome(a, b, znak string) bool {
	romeNumMap := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	arabicNumMap := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}

	if _, ok := arabicNumMap[a]; ok {
		if _, ok := arabicNumMap[b]; ok {
			return true
		}
	}
	if _, ok := romeNumMap[a]; ok {
		if _, ok := romeNumMap[b]; ok {
			return false
		}
	}

	panic("Введены неверные данные!")
}
func rightOper(znak string) bool {
	operatorMap := map[string]int{"+": 1, "-": 2, "*": 3, "/": 4}
	if _, ok := operatorMap[znak]; !ok {
		panic("Неверная арифметическая операция")
	}
	return true
}

func convertToRoman(arabNum int) string {
	var result string
	var arabRom = []struct {
		arabic   int
		romanNum string
	}{{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
		{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}
	for arabNum > 0 {
		for _, value := range arabRom {
			if arabNum >= value.arabic {
				result += value.romanNum
				arabNum -= value.arabic
				break
			}
		}
	}
	return result
}

