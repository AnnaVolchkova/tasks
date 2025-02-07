package brackets

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestIsValid(t *testing.T) {
	tt := []struct{
		name string
		input string
		expectedResult bool
	}{
		{
			name: "пустая строка",
			input: "",
			expectedResult: false,
		},
		{
			name: "невалидная строка",
			input: "(}{)",
			expectedResult: false,
		},
		{
			name: "валидная строка",
			input: "[(){[]}()]",
			expectedResult: true,
		},
		{
			name: "случайный набор символов",
			input: "ffjjjbjjfj[[[[((((((()))))))]]]",
			expectedResult: false,
		},
		{
			name: "одна открывающая скобка",
			input: "(",
			expectedResult: false,
		},
		{
			name: "одна закрывающая скобка",
			input: ")",
			expectedResult: false,
		},
		{
			name: "только открывающие скобки",
			input: "{{{{{{{{",
			expectedResult: false,
		},
		{
			name: "только закрывающие скобки",
			input: "}}}}}}}}",
			expectedResult: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsValid(tc.input)
			require.Equal(t, tc.expectedResult, actual)
		})
	}
}

// написать программу, которая при помощи стека проверяла валидность последовательностей скобок
// изменить сигнатуру linkedlist, чтобы она работала со строками, либо с рунами

// что означают валидные скобки, а что означают невалидные скобки
// придумать алгоритм, как это проверять
// написать это в коде

