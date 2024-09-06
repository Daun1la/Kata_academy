package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	m := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII",
		9: "IX", 10: "X", 11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
		16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX", 21: "XXI",
		22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII",
		28: "XXVIII", 29: "XXIX", 30: "XXX", 31: "XXXI", 32: "XXXII", 33: "XXXIII",
		34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII",
		39: "XXXIX", 40: "XL", 41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV",
		45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
		51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI",
		57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX", 61: "LXI", 62: "LXII",
		63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII",
		69: "LXIX", 70: "LXX", 71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV",
		75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX",
		80: "LXXX", 81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV",
		85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX",
		90: "XC", 91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV",
		96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	accept := [24]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "+", "-", "*", "/"}

	for {
		fmt.Println("Введите значение в формате '1 + 5' или 'II * VI'")
		text, _ := reader.ReadString('\n')
		text = strings.ToUpper(text)
		words := strings.Fields(text)

		// Проверка на три введенных символа
		if len(words) != 3 {
			fmt.Println("Неверное количество значений!")
			panic("Cтрока не является математической операцией")
		}

		// Проверка на ввод числа от 1 до 10 и правильных мат. символов
		count := 0
		for _, item := range accept {
			if (item == words[0]) || (item == words[1]) || (item == words[2]) {
				count += 1
			}
		}
		if words[0] != words[2] {
			if count < 3 {
				panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно и математические операторы '+', '-', '*'. '/'")
			}
		}
		if words[0] == words[2] {
			if count < 2 {
				panic("Калькулятор должен принимать на вход числа от 1 до 10 включительно и математические операторы '+', '-', '*'. '/'")
			}
		}

		// Проверка на воод арабских или римских чисел одновременно
		firstFlag := false
		lastFlag := false
		for _, item := range roman {
			if item == words[0] {
				firstFlag = true
			}
			if item == words[2] {
				lastFlag = true
			}
		}
		if (!firstFlag && lastFlag) || (firstFlag && !lastFlag) {
			panic("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
		} else {
			// Переводим римские числа в арабские
			for idx, value := range m {
				if value == words[0] {
					words[0] = fmt.Sprintf("%d", idx)
				}
				if value == words[2] {
					words[2] = fmt.Sprintf("%d", idx)
				}
			}
			first, _ := strconv.Atoi(words[0])
			last, _ := strconv.Atoi(words[len(words)-1])

			// firstFlag взведется в true, если в исходной сроке есть римская цифра
			if !firstFlag {
				switch words[1] {
				case "+":
					fmt.Println(first + last)
				case "-":
					fmt.Println(first - last)
				case "*":
					fmt.Println(first * last)
				case "/":
					fmt.Println(first / last)
				}
			} else {
				tmp := 0
				switch words[1] {
				case "+":
					tmp = first + last
				case "-":
					tmp = first - last
					if tmp < 1 {
						panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
					}
				case "*":
					tmp = first * last
				case "/":
					tmp = first / last
					if tmp < 1 {
						panic("Результатом работы калькулятора с римскими числами могут быть только положительные числа")
					}
				}
				for i, result := range m {
					if i == tmp {
						fmt.Println(result)
					}
				}
			}
		}
	}
}
