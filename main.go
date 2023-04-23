package main

import (
	"errors"
	"fmt"
)

func isOpenBracket(symb rune) bool {
	return symb == '(' || symb == '['
}

func isClosedBracket(symb rune) bool {
	return symb == ')' || symb == ']'
}

func isLetter(symb rune) bool {
	return symb >= 'a' && symb <= 'z'
}

func isOperationSign(symb rune) bool {
	return symb == '+' || symb == '-' || symb == '*' || symb == '/'
}

func read(notation string, i *int) (symb rune, err error) {

	r := *i
	*i++
	if *i >= len(notation) {
		return '?', errors.New("read(): i is out of notation string range")
	}
	return rune(notation[r]), nil
}

func isRightNotation(notation string, ip *int) bool {
	symb := rune(notation[*ip])
	if isLetter(symb) {
		read(notation, ip)
	}

	if isOpenBracket(symb) {
		var closingBracket rune
		if symb == '(' {
			closingBracket = ')'
		} else {
			closingBracket = ']'
		}

		read(notation, ip)

		if !isRightNotation(notation, ip) {
			return false
		}
		if *ip >= len(notation) {
			return false
		}
		if !isOperationSign(rune(notation[*ip])) {
			return false
		}
		read(notation, ip)

		if !isRightNotation(notation, ip) {
			return false
		}
		if *ip >= len(notation) {
			return false
		}

		if rune(notation[*ip]) != closingBracket {
			return false
		}
		read(notation, ip)
	}
	return true
}

func main() {
	fmt.Println("Правильная скобочная запись арифметических выражений с двумя видами скобок." +
		" Друг за другом на одном уровне вложенности могут стоять не более двух скобок." +
		" Каждая бинарная операция вместе с операндами берется в скобки." +
		" В правильной записи не могут присутствовать “лишние” (двойные) скобки и одна буква не может браться в скобки." +
		"\n\nПример." +
		"\nПравильная запись: ([(a+b)-(c*[(a-b)+(c*c)])]/((a-(b+c))*[a*(b+c)]))" +
		"\nНеправильная запись [(a)*([b-c*d]([a-c]/(a+b)*((c-d))))]")
	for {
		var notation string
		fmt.Print("Введите скобочную запись: ")
		fmt.Scanln(&notation)
		if len(notation) == 0 {
			break
		}
		var i = 0
		res := isRightNotation(notation, &i)
		if res == false || i < len(notation) {
			fmt.Println("Неправильная запись")
		} else {
			fmt.Println("Правильная запись")
		}
	}
}
