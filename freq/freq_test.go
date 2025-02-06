package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestWordFrequency(t *testing.T) {
	tt := []struct{
		name string
		inputArgument string
		expectedMap map[string]int
	} {
		{
			name: "empty string",
			inputArgument: "",
			expectedMap: map[string]int{},
		},
		{
			name: "full string",
			inputArgument: "hello world",
			expectedMap: map[string]int {"hello": 1, "world": 1},
		},
		{
			name: "full string duplicates",
			inputArgument: "hello world world world hello hello",
			expectedMap: map[string]int {"hello": 3, "world": 3},
		},
		{
			name: "case insensitive",
			inputArgument: "Hello world World woRld hELLo hello",
			expectedMap: map[string]int {"hello": 3, "world": 3},
		},
		{
			name: "special characters",
			inputArgument: "Hello, world; World;;; woRld,,, hELLo??? hello,;.!?",
			expectedMap: map[string]int {"hello": 3, "world": 3},
		},
		{
			name: "invalid characters",
			inputArgument: ",,,...!!!???:::",
			expectedMap: map[string]int{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			actualMap := wordFrequency(tc.inputArgument)
			
			for expWord, expFreq := range tc.expectedMap{
				actualFreq, exists := actualMap[expWord]
				require.True(t, exists)
				require.Equal(t, expFreq, actualFreq)
			}

			require.Len(t, actualMap, len(tc.expectedMap))
		})
	}
}