package kansuji

import (
	"strconv"
	"strings"
)

func SimplifiedChineseFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(simplifiedChineseStringInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[1])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(chars()[strParts[1][i]])
		}
	}

	return sb.String()
}

func SimplifiedChineseInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(simplifiedChineseStringInt(str))

	return sb.String()
}

func SimplifiedChineseUint(in uint64) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(simplifiedChineseStringInt(str))

	return sb.String()
}

func SimplifiedChineseFinancialFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	str = strings.Trim(str, "0")
	strParts := strings.Split(str, ".")

	sb.WriteString(simplifiedChineseStringFinancialInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point()[1])
		for i := 0; i < len(strParts[1]); i++ {
			rs := upperChars()[strParts[1][i]]
			if len(rs) >= 2 {
				sb.WriteRune(rs[1])
			} else {
				sb.WriteRune(rs[0])
			}
		}
	}

	return sb.String()
}

func SimplifiedChineseFinancialInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative()[1])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(simplifiedChineseStringFinancialInt(str))

	return sb.String()
}

func SimplifiedChineseFinancialUint(in uint64) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(simplifiedChineseStringFinancialInt(str))

	return sb.String()
}

func simplifiedChineseStringInt(in string) string {
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
			if len(r) >= 2 {
				sb.WriteRune(r[1])
			} else {
				sb.WriteRune(r[0])
			}
		}

		allzeros = true
	}

	return sb.String()
}

func simplifiedChineseStringFinancialInt(in string) string {
	if in == "" || in == "0" {
		return "零"
	}

	var integer [12]group
	var i, g, b int
	for i, g, b = len(in)-1, 0, 0; i >= 0; i-- {
		r := upperChars()[in[i]]
		if len(r) >= 2 {
			integer[g][b] = r[1]
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
			if len(r) >= 2 {
				sb.WriteRune(r[1])
			} else {
				sb.WriteRune(r[0])
			}
		}

		allzeros = true
	}

	return sb.String()
}
