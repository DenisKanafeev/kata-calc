//This is a simple calculator for one-digit numbers.
//It understands romans and integers, but not at the same time.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	input, x, y, op    string
	xrom, yrom         bool
	xint, yint, result int
)

const (
	welcome         = "This is a simple calculator for one-digit numbers. It understands romans and integers, but not at the same time."
	bye             = "OK, bye then!"
	invop           = "invalid operation"
	nosnum          = "no second number provided"
	nofnum          = "no first number provided"
	numrangeerr     = "only numbers from 1 to 10 allowed"
	numsyserr       = "numbers must be of a same system"
	romlessthanzero = "roman result < 0"
)

func main() {

	fmt.Println(welcome) //показываем юзеру привественное сообщение с описанием функционала

	for {
		xrom, yrom = false, false                                         //в начале цикла сбрасываем значения xrom и yrom на false
		fmt.Println("Enter number operation number. To stop enter stop.") //запрашиваем у юзера строку
		scanner := bufio.NewScanner(os.Stdin)
		_ = scanner.Scan()
		input = scanner.Text()

		if input == "stop" { //если юзер ввёл строку "stop", завершаем работу
			fmt.Println(bye) //и прощаемся
			break
		}

		input = strings.ReplaceAll(input, " ", "") //удаляем все пробелы из строки

		count := 0
		for i := 0; i <= len(input)-1; i++ { //при помощи счётчика проверяем, сколько арифметичеких операций указано в строке
			if input[i] == 42 || input[i] == 43 || input[i] == 45 || input[i] == 47 { //символы операций смотрим в ASCII
				count++
			}
		}

		if count != 1 { //если количество операций не равно 1, завершаем работу
			panic(invop) //выводим сообщение об ошибке
		}

		switch { //проверяем, какую именно операцию указал юзер, и записываем её в переменную op
		case strings.Contains(input, "+"):
			op = "+"
		case strings.Contains(input, "-"):
			op = "-"
		case strings.Contains(input, "*"):
			op = "*"
		case strings.Contains(input, "/"):
			op = "/"
		}

		x = strings.Split(input, op)[0] //делим строку по оператору и часть строки до операции записываем в переменную x
		y = strings.Split(input, op)[1] //делим строку по оператору и часть строки после операции записываем в переменную y

		if x == "" { //если до операции ничего не указано, завершаем работу
			panic(nofnum) //и выводим сообщение об ошибке (не указано первое число)
		}
		if y == "" { //если после операции ничего не указано, завершаем работу
			panic(nosnum) //и выводим сообщение об ошибке (не указано второе число)
		}

		if x[0] < 48 || x[0] > 57 { //если первый символ первого числа не является цифрой (проверяем по ASCII)
			xrom = true        //значит юзер указал первое число в римской системе, отражаем это в переменной xrom
			xint = romtoint(x) //переводим его в int с помощью переводчика и записываем в переменную xint
		} else {
			xint, _ = strconv.Atoi(x) //если же первый сомвол это цифра, переводим строку в число и записываем в переменную xint
		}

		if y[0] < 48 || y[0] > 57 { //то же самое делаем со вторым числом
			yrom = true
			yint = romtoint(y)
		} else {
			yint, _ = strconv.Atoi(y)
		}

		if xint < 1 || xint > 10 || yint < 1 || yint > 10 { //если одно из чисел меньше 1 или больше 10
			panic(numrangeerr) //завершаем работу и показываем сообщение об ошибке
		}

		if xrom != yrom { //если числа указаны в разных системах
			panic(numsyserr) //завершаем работу и показываем сообщение об ошибке
		}

		switch { //вычисляем результат операции с числами и записываем его в переменную result
		case op == "+":
			result = xint + yint
		case op == "-":
			result = xint - yint
		case op == "*":
			result = xint * yint
		case op == "/":
			result = xint / yint
		}

		if xrom && result < 0 { //если юзер указывал римские цифры, а результат получился меньше 0
			panic(romlessthanzero) //завершаем работу и показываем сообщение об ошибке
		}

		if xrom { //если юзер вводил римские цифры
			fmt.Println(x, op, y, "=", inttorom(result)) //выводим результат в римской системе с помощью обратного переводчика
		} else {
			fmt.Println(x, op, y, "=", result) //если же нет, то просто выводим результат
		}
	}
}

func inttorom(num int) string { //это функция-переводчик для перевода int чисел до 100 (int) в римские (string)
	rom := ""       //сейчас будем считать десятки в числе и записывать их в переменную rom
	if num == 100 { //100 это исключение, поэтому переводим его буквально
		rom = "C"
	} else if num >= 90 { //если число больше 90, записываем в rom значение "XC", а оставшиеся единицы сохраняем в num
		rom = "XC"
		num = num - 90
	} else if num >= 50 { //если число больше 50, записываем в rom значение "XC", а оставшиеся единицы сохраняем в num
		rom = "L"
		num -= 50
		for num > 10 {
			rom += "X"
			num -= 10
		}
	} else if num >= 40 { //если число больше 40, записываем в rom значение "XC", а оставшиеся единицы сохраняем в num
		rom = "XL"
		num -= 40
	} else if num >= 10 { //если число больше 10
		for num >= 10 {
			rom += "X" //добавляем в rom значение "X"
			num -= 10  //пока десятки не закончатся
		}
	}
	return rom + inttoromless10(num) //добавляем к десяткам единицы с помощью переводчика единиц и возвращаем
}

func inttoromless10(num int) string { //это функция-переводчик для перевода int цифр (int) в римские цифры (string)
	rom := ""
	if num == 9 { //она буквально берёт каждую цифру и переводит римскую строку
		rom = "IX"
	} else if num == 8 {
		rom = "VIII"
	} else if num == 7 {
		rom = "VII"
	} else if num == 6 {
		rom = "VI"
	} else if num == 5 {
		rom = "V"
	} else if num == 4 {
		rom = "IV"
	} else if num == 3 {
		rom = "III"
	} else if num == 2 {
		rom = "II"
	} else if num == 1 {
		rom = "I"
	}
	return rom //возвращаем количество единиц в римской системе
}

func romtoint(rom string) int { //это функция-переводчик цифр из римских (string) в обычные (int)
	num := 0
	if rom == "IV" || rom == "iv" || rom == "iV" || rom == "Iv" { //цифры 4 и 9 в римской системе являются исключениями
		num = 4 //потому что для их получения требуется вычитание
	} else if rom == "IX" || rom == "ix" || rom == "iX" || rom == "Ix" { //поэтому переводим их буквально
		num = 9 //при этом допускаем, что юзер ввёл их разными сочетаниями больших и маленьких букв
	} else {
		for i := 0; i <= len(rom)-1; i++ { //для остальных цифр используем цикл, где они складываются
			if rom[i] == 73 || rom[i] == 105 { //опять же позволяя юзеру использовать разные сочетания больших и маленьких букв
				num += 1
			} else if rom[i] == 86 || rom[i] == 118 {
				num += 5
			} else if rom[i] == 88 || rom[i] == 120 {
				num += 10
			}
		}
	}
	return num //возвращаем получившуюся цифру в int формате
}
