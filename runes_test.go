package runes_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/talon-one/runes"
)

func marshal(i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return b
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		slice  []rune
		prefix []rune
		result bool
	}{
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0, 1, 2}, true},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{1, 2}, false},
		{[]rune{0}, []rune{0, 1, 2}, false},
	}
	for i, test := range tests {
		// Prepare Integrity check
		before := marshal(test)

		// Result check
		require.Equal(t, test.result, runes.HasPrefix(test.slice, test.prefix), "Test #%d failed", i)

		// Integrity check
		require.EqualValues(t, before, marshal(test), "Integrity check for Test #%d failed", i)
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		slice  []rune
		suffix []rune
		result bool
	}{
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{3, 4, 5}, true},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{3, 4}, false},
		{[]rune{0}, []rune{3, 4, 5}, false},
	}
	for i, test := range tests {
		// Prepare Integrity check
		before := marshal(test)

		// Result check
		require.Equal(t, test.result, runes.HasSuffix(test.slice, test.suffix), "Test #%d failed", i)

		// Integrity check
		require.EqualValues(t, before, marshal(test), "Integrity check for Test #%d failed", i)
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		sliceA []rune
		sliceB []rune
		result bool
	}{
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0, 1, 2, 3, 4, 5}, true},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0, 1, 2, 3, 4, 6}, false},
		{nil, nil, true},
		{[]rune{0}, []rune{3, 4, 5}, false},
		{[]rune{0}, nil, false},
	}
	for i, test := range tests {
		// Prepare Integrity check
		before := marshal(test)

		// Result check
		require.Equal(t, test.result, runes.Equal(test.sliceA, test.sliceB), "Test #%d failed", i)

		// Integrity check
		require.EqualValues(t, before, marshal(test), "Integrity check for Test #%d failed", i)
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		slice  []rune
		sep    []rune
		result int
	}{
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0, 1, 2, 3, 4, 5}, 0},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0}, 0},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{1, 2, 3}, 1},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{0, 1, 2, 3, 4, 5, 6}, -1},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{5}, 5},
		{[]rune{0, 1, 2, 3, 4, 5}, []rune{6}, -1},
	}
	for i, test := range tests {
		// Prepare Integrity check
		before := marshal(test)

		// Result check
		require.Equal(t, test.result, runes.Index(test.slice, test.sep), "Test #%d failed", i)

		// Integrity check
		require.EqualValues(t, before, marshal(test), "Integrity check for Test #%d failed", i)
	}
}
