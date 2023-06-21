package stringx_test

import (
	"strings"
	"testing"

	"github.com/eos175/stringx"
)

func TestHasPrefixFast(t *testing.T) {
	testCases := []struct {
		s        string
		prefix   string
		expected bool
	}{
		{"hello world", "he", true},
		{"hello world", "hel", true},
		{"hello world", "hell", true},
		{"hello world", "hello", true},
		{"hello world", "hello ", true},
		{"hello world", "", true},
		{"", "", true},
		{"", "hello", false},
		{"hello world", "goodbye", false},
	}

	for _, tc := range testCases {
		got := stringx.HasPrefixFast(tc.s, tc.prefix)
		if got != tc.expected {
			t.Errorf("hasPrefixFast(%q, %q) = %v; expected %v", tc.s, tc.prefix, got, tc.expected)
		}
	}
}

var sLong = "hello world, this is a very long string"
var sShort = "hello"
var prefixLong = "hello world, this is a very lon"
var prefixShort = "he"
var prefixNot = "goodbye"
var prefixEmpty = ""

func BenchmarkHasPrefixFastLongStringLongPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringx.HasPrefixFast(sLong, prefixLong)
	}
}

func BenchmarkHasPrefixStdLongStringLongPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(sLong, prefixLong)
	}
}

func BenchmarkHasPrefixFastLongStringShortPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringx.HasPrefixFast(sLong, prefixShort)
	}
}

func BenchmarkHasPrefixStdLongStringShortPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(sLong, prefixShort)
	}
}

func BenchmarkHasPrefixFastShortStringLongPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringx.HasPrefixFast(sShort, prefixLong)
	}
}

func BenchmarkHasPrefixStdShortStringLongPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(sShort, prefixLong)
	}
}

func BenchmarkHasPrefixFastShortStringShortPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringx.HasPrefixFast(sShort, prefixShort)
	}
}

func BenchmarkHasPrefixStdShortStringShortPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(sShort, prefixShort)
	}
}

func BenchmarkHasPrefixFastInvalidPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringx.HasPrefixFast(sLong, prefixNot)
	}
}

func BenchmarkHasPrefixStdInvalidPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.HasPrefix(sLong, prefixNot)
	}
}
