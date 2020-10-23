package kansuji_test

import (
	"testing"

	"github.com/GreenYun/kansuji"
)

func TestTraditionalChineseFloat(t *testing.T) {
	intTestCases := []struct {
		number float64
		prec   int
		want   string
	}{
		{0, -1, "零"},
		{-1.01, -1, "負一點零一"},
		{1234567, 5, "一百二十三萬四千五百六十七點零零零零零"},
		{-1001101.987654321, 8, "負一百萬一千一百零一點九八七六五四三二"},
	}

	for _, tc := range intTestCases {
		got := kansuji.TraditionalChineseFloat(tc.number, tc.prec)
		if got != tc.want {
			t.Errorf("TraditionalChineseFloat(%f, %d) got \"%s\" but wanted \"%s\"", tc.number, tc.prec, got, tc.want)
		}
	}
}

func TestTraditionalChineseInt(t *testing.T) {
	intTestCases := []struct {
		number int64
		want   string
	}{
		{0, "零"},
		{-1, "負一"},
		{0x7fffffffffffffff, "九百二十二京三千三百七十二兆零三百六十八億五千四百七十七萬五千八百零七"},
		{101000010100100000, "十京一千兆零一百零一億零一十萬"},
		{-10010001010100101, "負一京零一十兆零一十億一千零一十萬零一百零一"},
	}

	for _, tc := range intTestCases {
		got := kansuji.TraditionalChineseInt(tc.number)
		if got != tc.want {
			t.Errorf("TraditionalChineseInt(%d) got \"%s\" but wanted \"%s\"", tc.number, got, tc.want)
		}
	}
}

func TestTraditionalChineseUint(t *testing.T) {
	intTestCases := []struct {
		number uint64
		want   string
	}{
		{0, "零"},
		{0xffffffffffffffff, "一千八百四十四京六千七百四十四兆零七百三十七億零九百五十五萬一千六百一十五"},
	}

	for _, tc := range intTestCases {
		got := kansuji.TraditionalChineseUint(tc.number)
		if got != tc.want {
			t.Errorf("TraditionalChineseUint(%d) got \"%s\" but wanted \"%s\"", tc.number, got, tc.want)
		}
	}
}
