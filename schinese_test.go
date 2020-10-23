package kansuji_test

import (
	"testing"

	"github.com/GreenYun/kansuji"
)

func TestSimplifiedChineseFloat(t *testing.T) {
	intTestCases := []struct {
		number float64
		prec   int
		want   string
	}{
		{0, -1, "零"},
		{-1.01, -1, "负一点零一"},
		{1234567, 5, "一百二十三万四千五百六十七点零零零零零"},
		{-1001101.987654321, 8, "负一百万一千一百零一点九八七六五四三二"},
	}

	for _, tc := range intTestCases {
		got := kansuji.SimplifiedChineseFloat(tc.number, tc.prec)
		if got != tc.want {
			t.Errorf("SimplifiedChineseFloat(%f, %d) got \"%s\" but wanted \"%s\"", tc.number, tc.prec, got, tc.want)
		}
	}
}

func TestSimplifiedChineseInt(t *testing.T) {
	intTestCases := []struct {
		number int64
		want   string
	}{
		{0, "零"},
		{-1, "负一"},
		{0x7fffffffffffffff, "九百二十二京三千三百七十二兆零三百六十八亿五千四百七十七万五千八百零七"},
		{101000010100100000, "十京一千兆零一百零一亿零一十万"},
		{-10010001010100101, "负一京零一十兆零一十亿一千零一十万零一百零一"},
	}

	for _, tc := range intTestCases {
		got := kansuji.SimplifiedChineseInt(tc.number)
		if got != tc.want {
			t.Errorf("SimplifiedChineseInt(%d) got \"%s\" but wanted \"%s\"", tc.number, got, tc.want)
		}
	}
}

func TestSimplifiedChineseUint(t *testing.T) {
	intTestCases := []struct {
		number uint64
		want   string
	}{
		{0, "零"},
		{0xffffffffffffffff, "一千八百四十四京六千七百四十四兆零七百三十七亿零九百五十五万一千六百一十五"},
	}

	for _, tc := range intTestCases {
		got := kansuji.SimplifiedChineseUint(tc.number)
		if got != tc.want {
			t.Errorf("SimplifiedChineseUint(%d) got \"%s\" but wanted \"%s\"", tc.number, got, tc.want)
		}
	}
}
