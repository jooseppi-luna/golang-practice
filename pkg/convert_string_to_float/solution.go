package convert_string_to_float

import "fmt"
import "strings"

var runeToFloat = map[rune]float64{
	'0': 0.0,
	'1': 1.0,
	'2': 2.0,
	'3': 3.0,
	'4': 4.0,
	'5': 5.0,
	'6': 6.0,
	'7': 7.0,
	'8': 8.0,
	'9': 9.0,
}

func StringToFloat(num string) float64 {
	numNum := 0.0
	sign := 1.0
	decimalOffset := 10.0

	if num[0] == '-' {
		sign = -1.0
		num = num[1:]
	}

	splitNum := strings.Split(num, ".")

	fmt.Printf("splitNum: %+v\n", splitNum)

	for _, digit := range splitNum[0] {
		numNum = numNum*10.0 + runeToFloat[digit]
	}

	if len(splitNum) == 1 {
		return sign*numNum
	}

	for _, digit := range splitNum[1] {
		if runeToFloat[digit] != 0.0 {
			numNum = numNum + runeToFloat[digit] / decimalOffset
		}
		decimalOffset = decimalOffset*10.0
	}

	return sign*numNum
}

