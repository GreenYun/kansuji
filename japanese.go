package kansuji

import (
	"strconv"
	"strings"
)

// In Japanese, there are several styles of "ones".
// NoOnes means no "one"s before "ten"s, "hundred"s and "thousand"s.
// CommonOnes puts the optional "one" before every "thousand".
// AllOnes puts every "one" explicitly (usually used in financial situations).
const (
	NoOnes int = iota
	CommonOnes
	AllOnes
)

// JapaneseFloat returns kansuji of a `float64` in Japanese style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
//
// `prec` has the same meaning of `strconv.FormatFloat`.
func JapaneseFloat(in float64, ones, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[2])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	strParts := strings.Split(str, ".")

	sb.WriteString(japaneseStringInt(strParts[0], ones))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point[2])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(numbers[strParts[1][i]-'0'])
		}
	}

	return sb.String()
}

// JapaneseInt returns kansuji of a `int64` in Japanese style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseInt(in int64, ones int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[2])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(japaneseStringInt(str, ones))

	return sb.String()
}

// JapaneseUint returns kansuji of a `uint64` in Japanese style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseUint(in uint64, ones int) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(japaneseStringInt(str, ones))

	return sb.String()
}

// JapaneseFinancialFloat returns kansuji of a `float64` in Japanese financial style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
//
// `prec` has the same meaning of `strconv.FormatFloat`.
func JapaneseFinancialFloat(in float64, ones, prec int) string {
	return ToJapaneseLaw(JapaneseFloat(in, ones, prec))
}

// JapaneseFinancialInt returns kansuji of a `int64` in Japanese financial style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseFinancialInt(in int64, ones int) string {
	return ToJapaneseLaw(JapaneseInt(in, ones))
}

// JapaneseFinancialUint returns kansuji of a `uint64` in Japanese financial style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseFinancialUint(in uint64, ones int) string {
	return ToJapaneseLaw(JapaneseUint(in, ones))
}

// JapaneseOldFinancialFloat returns kansuji of a `float64` in Japanese old style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
//
// `prec` has the same meaning of `strconv.FormatFloat`.
func JapaneseOldFinancialFloat(in float64, ones, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[2])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(japaneseStringOldFinancialInt(strParts[0], ones))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point[2])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(financialNumbers[strParts[1][i]-'0'][2])
		}
	}

	return sb.String()
}

// JapaneseOldFinancialInt returns kansuji of a `int64` in Japanese old style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseOldFinancialInt(in int64, ones int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[2])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(japaneseStringOldFinancialInt(str, ones))

	return sb.String()
}

// JapaneseOldFinancialUint returns kansuji of a `uint64` in Japanese old style.
// `ones` is one of `NoOnes`, `CommonOnes` and `AllOnes`.
func JapaneseOldFinancialUint(in uint64, ones int) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(japaneseStringOldFinancialInt(str, ones))

	return sb.String()
}

func japaneseStringInt(in string, ones int) string {
	if in == "" || in == "0" {
		return "零"
	}

	var integer [12]group
	var i, g, b int
	for i, g, b = len(in)-1, 0, 0; i >= 0; i-- {
		integer[g][b] = numbers[in[i]-'0']
		b++
		if b == 4 {
			g++
			if g > 12 {
				g = 12
				break
			}
			b = 0
		}
	}

	var sb strings.Builder

	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; g >= 0; g-- {
		for ; b >= 0; b-- {
			if integer[g][b] == numbers[0] {
				continue
			}
			if !(integer[g][b] == numbers[1] && b > 0) || ones == AllOnes {
				sb.WriteRune(integer[g][b])
			} else {
				if ones == CommonOnes && g >= 1 && b >= 3 {
					sb.WriteRune(integer[g][b])
				}
			}
			if b > 0 {
				sb.WriteRune(multipliers[b-1])
			}
		}
		b = 3
		if g > 0 {
			sb.WriteRune(multipliers2[g-1][2])
		}
	}

	return sb.String()
}

func japaneseStringOldFinancialInt(in string, ones int) string {
	if in == "" || in == "0" {
		return "零"
	}

	var integer [12]group
	var i, g, b int
	for i, g, b = len(in)-1, 0, 0; i >= 0; i-- {
		integer[g][b] = financialNumbers[in[i]-'0'][2]
		b++
		if b == 4 {
			g++
			if g > 12 {
				g = 12
				break
			}
			b = 0
		}
	}

	var sb strings.Builder

	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; g >= 0; g-- {
		for ; b >= 0; b-- {
			if integer[g][b] == numbers[0] {
				continue
			}
			if !(integer[g][b] == financialNumbers[1][2] && b > 0) || ones == AllOnes {
				sb.WriteRune(integer[g][b])
			} else {
				if ones == CommonOnes && g >= 1 && b >= 3 {
					sb.WriteRune(integer[g][b])
				}
			}
			if b > 0 {
				sb.WriteRune(financialMultipliers[b-1][2])
			}
		}
		b = 3
		if g > 0 {
			sb.WriteRune(multipliers2[g-1][2])
		}
	}

	return sb.String()
}
