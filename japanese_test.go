package kansuji_test

import (
	"testing"

	"github.com/GreenYun/kansuji"
)

func TestJapaneseFloat(t *testing.T) {
	intTestCases := []struct {
		number float64
		ones   int
		prec   int
		want   string
	}{
		{0, 0, -1, "零"},
		{-1.01, 0, -1, "負一点零一"},
		{1234567, 1, 5, "百二十三万四千五百六十七点零零零零零"},
		{-1001101.987654321, 2, 8, "負一百万一千一百一点九八七六五四三二"},
	}

	for _, tc := range intTestCases {
		got := kansuji.JapaneseFloat(tc.number, tc.ones, tc.prec)
		if got != tc.want {
			t.Errorf("JapaneseFloat(%f, %d, %d) got \"%s\" but wanted \"%s\"", tc.number, tc.ones, tc.prec, got, tc.want)
		}
	}
}

func TestJapaneseInt(t *testing.T) {
	intTestCases := []struct {
		number int64
		ones   int
		want   string
	}{
		{0, 0, "零"},
		{-1, 0, "負一"},
		{0x7fffffffffffffff, 0, "九百二十二京三千三百七十二兆三百六十八億五千四百七十七万五千八百七"},
		{101010001010101101, 0, "十京千十兆十億千十万千百一"},
		{101010001010101101, 1, "十京一千十兆十億一千十万千百一"},
		{-1001101, 2, "負一百万一千一百一"},
	}

	for _, tc := range intTestCases {
		got := kansuji.JapaneseInt(tc.number, tc.ones)
		if got != tc.want {
			t.Errorf("JapaneseInt(%d, %d) got \"%s\" but wanted \"%s\"", tc.number, tc.ones, got, tc.want)
		}
	}
}

func TestJapaneseUint(t *testing.T) {
	intTestCases := []struct {
		number uint64
		ones   int
		want   string
	}{
		{0, 0, "零"},
		{0xffffffffffffffff, 0, "千八百四十四京六千七百四十四兆七百三十七億九百五十五万千六百十五"},
		{0xffffffffffffffff, 1, "一千八百四十四京六千七百四十四兆七百三十七億九百五十五万千六百十五"},
		{0xffffffffffffffff, 2, "一千八百四十四京六千七百四十四兆七百三十七億九百五十五万一千六百一十五"},
	}

	for _, tc := range intTestCases {
		got := kansuji.JapaneseUint(tc.number, tc.ones)
		if got != tc.want {
			t.Errorf("JapaneseUint(%d, %d) got \"%s\" but wanted \"%s\"", tc.number, tc.ones, got, tc.want)
		}
	}
}
