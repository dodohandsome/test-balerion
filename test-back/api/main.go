package main

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

var digitWords = []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
var unitWords = []string{"", "สิบ", "ร้อย", "พัน", "หมื่น", "แสน"}

func convertGroup(n string) string {
	res := ""
	l := len(n)

	for i, ch := range n {
		pos := l - i
		digit := int(ch - '0')

		if digit == 0 {
			continue
		}

		if pos == 2 {
			switch digit {
			case 1:
				res += "สิบ"
			case 2:
				res += "ยี่สิบ"
			default:
				res += digitWords[digit] + "สิบ"
			}
			continue
		}

		if pos == 1 {
			if digit == 1 {
				if l > 1 {
					res += "เอ็ด"
				} else {
					res += digitWords[digit]
				}
			} else {
				res += digitWords[digit]
			}
			continue
		}

		res += digitWords[digit] + unitWords[pos-1]
	}

	return res
}

func convertSatang(n string) string {
	res := ""
	l := len(n)
	digitFirst := 0
	for i, ch := range n {
		pos := l - i
		digit := int(ch - '0')

		if digit == 0 {
			continue
		}

		if pos == 2 {
			digitFirst = digit
			switch digit {
			case 1:
				res += "สิบ"
			case 2:
				res += "ยี่สิบ"
			default:
				res += digitWords[digit] + "สิบ"
			}
			continue
		}

		if pos == 1 {
			if digitFirst > 0 && digit == 1 {
				res += "เอ็ด"
			} else {
				res += digitWords[digit]
			}

		}
	}

	return res
}

func convertNumber(n string) string {
	n = strings.TrimLeft(n, "0")
	if n == "" {
		return "ศูนย์"
	}
	if len(n) <= 6 {
		return convertGroup(n)
	}
	prefix := n[:len(n)-6]
	suffix := n[len(n)-6:]
	return convertNumber(prefix) + "ล้าน" + convertGroup(suffix)
}

func convertDecimalToThaiBaht(d decimal.Decimal) string {
	s := d.StringFixed(2)
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	fractionPart := parts[1]

	result := convertNumber(integerPart) + "บาท"
	if fractionPart == "00" {
		result += "ถ้วน"
	} else {
		result += convertSatang(fractionPart) + "สตางค์"
	}
	return result
}

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}
	for _, input := range inputs {
		fmt.Println("Input:", input.String())
		fmt.Println("Output:", convertDecimalToThaiBaht(input))
	}
}
