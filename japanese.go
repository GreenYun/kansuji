package kansuji

import (
	"strconv"
	"strings"
)

const (
	NoOnes int = iota
	CommonOnes
	AllOnes
)

func JapaneseFloat(in float64, ones, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(japaneseStringInt(strParts[0], ones))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[2])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(chars()[strParts[1][i]])
		}
	}

	return sb.String()
}

func JapaneseInt(in int64, ones int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(japaneseStringInt(str, ones))

	return sb.String()
}

func JapaneseUint(in uint64, ones int) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(japaneseStringInt(str, ones))

	return sb.String()
}

func JapaneseFinancialFloat(in float64, ones, prec int) string {
	return ToJapaneseLaw(JapaneseFloat(in, ones, prec))
}

func JapaneseFinancialInt(in int64, ones int) string {
	return ToJapaneseLaw(JapaneseInt(in, ones))
}

func JapaneseFinancialUint(in uint64, ones int) string {
	return ToJapaneseLaw(JapaneseUint(in, ones))
}

func JapaneseOldFinancialFloat(in float64, ones, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(japaneseStringOldFinancialInt(strParts[0], ones))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[2])
		for i := 0; i < len(strParts[1]); i++ {
			rs := upperChars()[strParts[1][i]]
			if len(rs) >= 3 {
				sb.WriteRune(rs[2])
			} else {
				sb.WriteRune(rs[0])
			}
		}
	}

	return sb.String()
}

func JapaneseOldFinancialInt(in int64, ones int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(japaneseStringOldFinancialInt(str, ones))

	return sb.String()
}

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
		integer[g][b] = chars()[in[i]]
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
			if integer[g][b] == chars()['0'] {
				continue
			}
			if !(integer[g][b] == chars()['1'] && b > 0) || ones == AllOnes {
				sb.WriteRune(integer[g][b])
			} else {
				if ones == CommonOnes && g >= 1 && b >= 3 {
					sb.WriteRune(integer[g][b])
				}
			}
			if b > 0 {
				sb.WriteRune(multipliers()[b-1])
			}
		}
		b = 3
		if r := multipliersAgain()[g]; r != nil {
			if len(r) >= 3 {
				sb.WriteRune(r[2])
			} else {
				sb.WriteRune(r[0])
			}
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
		r := upperChars()[in[i]]
		if len(r) >= 3 {
			integer[g][b] = r[2]
		} else {
			integer[g][b] = r[0]
		}
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
			if integer[g][b] == chars()['0'] {
				continue
			}
			if !(integer[g][b] == upperChars()['1'][2] && b > 0) || ones == AllOnes {
				sb.WriteRune(integer[g][b])
			} else {
				if ones == CommonOnes && g >= 1 && b >= 3 {
					sb.WriteRune(integer[g][b])
				}
			}
			if b > 0 {
				sb.WriteRune(upperMultipliers()[b-1][2])
			}
		}
		b = 3
		if r := multipliersAgain()[g]; r != nil {
			if len(r) >= 3 {
				sb.WriteRune(r[2])
			} else {
				sb.WriteRune(r[0])
			}
		}
	}

	return sb.String()
}
