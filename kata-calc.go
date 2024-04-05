package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var input, x, y, op string
	var xrom, yrom bool
	var xint, yint, result int

	fmt.Println("Enter number operation number (split by space)")
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	input = scanner.Text()

	x = (strings.Fields(input)[0])
	op = (strings.Fields(input)[1])
	y = (strings.Fields(input)[2])

	if x[0] < 48 || x[0] > 57 {
		xrom = true
		xint = romtoint(x)
	} else {
		xint, _ = strconv.Atoi(x)

	}

	if y[0] < 48 || y[0] > 57 {
		yrom = true
		yint = romtoint(y)
	} else {
		yint, _ = strconv.Atoi(y)
	}

	if xint < 1 || xint > 10 || yint < 1 || yint > 10 {
		panic("only numbers from 0 to 10")
	}

	if xrom != yrom {
		panic("numbers must be of a same system")
	}

	switch {
	case op == "+":
		result = xint + yint
	case op == "-":
		result = xint - yint
	case op == "*":
		result = xint * yint
	case op == "/":
		result = xint / yint
	}

	if xrom && result < 0 {
		panic("roman result < 0")
	}

	if xrom == yrom && xrom {
		fmt.Println(x, op, y, "=", inttorom(result))
	} else if xrom == yrom && !xrom {
		fmt.Println(x, op, y, "=", result)
	}
}

func inttorom(num int) string {
	rom := ""
	if num == 100 {
		rom = "C"
	} else if num >= 90 {
		rom = "XC"
		num = num - 90
	} else if num >= 50 {
		rom = "L"
		num -= 50
		for num > 10 {
			rom += "X"
			num -= 10
		}
	} else if num >= 40 {
		rom = "XL"
		num -= 40
	} else if num >= 10 {
		for num >= 10 {
			rom += "X"
			num -= 10
		}
	}
	return rom + romless10(num)
}

func romless10(num int) string {
	rom := ""
	if num == 9 {
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
	return rom
}

func romtoint(rom string) int {
	num := 0
	if rom == "IV" || rom == "iv" {
		num = 4
	} else if rom == "IX" || rom == "ix" {
		num = 9
	} else {
		for i := 0; i <= len(rom)-1; i++ {
			if rom[i] == 73 || rom[i] == 105 {
				num += 1
			} else if rom[i] == 86 || rom[i] == 118 {
				num += 5
			} else if rom[i] == 88 || rom[i] == 120 {
				num += 10
			}
		}
	}
	return num
}
