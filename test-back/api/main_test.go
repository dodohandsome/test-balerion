// main_test.go
package main

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConvertDecimalToThaiBaht(t *testing.T) {
	tests := []struct {
		input    decimal.Decimal
		expected string
	}{
		{
			input:    decimal.NewFromFloat(1234),
			expected: "หนึ่งพันสองร้อยสามสิบสี่บาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(33333.75),
			expected: "สามหมื่นสามพันสามร้อยสามสิบสามบาทเจ็ดสิบห้าสตางค์",
		},
		{
			input:    decimal.NewFromFloat(0),
			expected: "ศูนย์บาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1),
			expected: "หนึ่งบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(1.01),
			expected: "หนึ่งบาทหนึ่งสตางค์",
		},
		{
			input:    decimal.NewFromFloat(11.01),
			expected: "สิบเอ็ดบาทหนึ่งสตางค์",
		},
		{
			input:    decimal.NewFromFloat(11.11),
			expected: "สิบเอ็ดบาทสิบเอ็ดสตางค์",
		},
		{
			input:    decimal.NewFromFloat(21.21),
			expected: "ยี่สิบเอ็ดบาทยี่สิบเอ็ดสตางค์",
		},
		{
			input:    decimal.NewFromFloat(331.00),
			expected: "สามร้อยสามสิบเอ็ดบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(331),
			expected: "สามร้อยสามสิบเอ็ดบาทถ้วน",
		},
		{
			input:    decimal.NewFromFloat(111.01),
			expected: "หนึ่งร้อยสิบเอ็ดบาทหนึ่งสตางค์",
		},
		{
			input:    decimal.NewFromFloat(111.11),
			expected: "หนึ่งร้อยสิบเอ็ดบาทสิบเอ็ดสตางค์",
		},
		{
			input:    decimal.NewFromFloat(111.91),
			expected: "หนึ่งร้อยสิบเอ็ดบาทเก้าสิบเอ็ดสตางค์",
		},
		{
			input:    decimal.NewFromFloat(111.09),
			expected: "หนึ่งร้อยสิบเอ็ดบาทเก้าสตางค์",
		},
	}

	for _, tt := range tests {
		actual := convertDecimalToThaiBaht(tt.input)
		assert.Equal(t, tt.expected, actual, "convertDecimalToThaiBaht(%s)", tt.input.String())
	}
}
