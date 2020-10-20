package kansuji

import (
	"strconv"
	"strings"
)

func TraditionalChineseFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(traditionalChineseStringInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[0])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(chars()[strParts[1][i]])
		}
	}

	return sb.String()
}

func TraditionalChineseInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(traditionalChineseStringInt(str))

	return sb.String()
}

func TraditionalChineseUint(in uint64) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(traditionalChineseStringInt(str))

	return sb.String()
}

func TraditionalChineseFinancialFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(traditionalChineseStringFinancialInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[0])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(upperChars()[strParts[1][i]][0])
		}
	}

	return sb.String()
}

func TraditionalChineseFinancialInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[0])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(traditionalChineseStringFinancialInt(str))

	return sb.String()
}

func TraditionalChineseFinancialUint(in uint64) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(traditionalChineseStringFinancialInt(str))

	return sb.String()
}

func traditionalChineseStringInt(in string) string {
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

	allzeros := true
	if b == 2 && integer[g][1] == chars()['1'] {
		b--
		if b > 0 {
			sb.WriteRune(multipliers()[b-1])
		}
		allzeros = false
	}
	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; integer[g][b] == chars()['0']; b-- {
		if b == 0 {
			break
		}
	}
	if b < 0 {
		g--
		b = 3
	}
	for ; g >= 0; g-- {
		for ; b >= 0; b-- {
			if integer[g][b] == chars()['0'] {
				if b == 0 {
					break
				}
				if integer[g][b-1] == chars()['0'] {
					continue
				}
				sb.WriteRune(integer[g][b])
				continue
			}
			allzeros = false
			sb.WriteRune(integer[g][b])
			if b > 0 {
				sb.WriteRune(multipliers()[b-1])
			}
		}
		b = 3
		if r := multipliersAgain()[g]; r != nil && !allzeros {
			sb.WriteRune(r[0])
		}

		allzeros = true
	}

	return sb.String()
}

func traditionalChineseStringFinancialInt(in string) string {
	if in == "" || in == "0" {
		return "零"
	}

	var integer [12]group
	var i, g, b int
	for i, g, b = len(in)-1, 0, 0; i >= 0; i-- {
		integer[g][b] = upperChars()[in[i]][0]
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

	allzeros := true
	if b == 2 && integer[g][1] == upperChars()['1'][0] {
		b--
		if b > 0 {
			sb.WriteRune(upperMultipliers()[b-1][0])
		}
		allzeros = false
	}
	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; integer[g][b] == chars()['0']; b-- {
		if b == 0 {
			break
		}
	}
	if b < 0 {
		g--
		b = 3
	}
	for ; g >= 0; g-- {
		for ; b >= 0; b-- {
			if integer[g][b] == chars()['0'] {
				if b == 0 {
					break
				}
				if integer[g][b-1] == chars()['0'] {
					continue
				}
				sb.WriteRune(integer[g][b])
				continue
			}
			allzeros = false
			sb.WriteRune(integer[g][b])
			if b > 0 {
				sb.WriteRune(upperMultipliers()[b-1][0])
			}
		}
		b = 3
		if r := multipliersAgain()[g]; r != nil && !allzeros {
			sb.WriteRune(r[0])
		}

		allzeros = true
	}

	return sb.String()
}
