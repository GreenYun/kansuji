package kansuji

import (
	"strconv"
	"strings"
)

// SimplifiedChineseFloat returns kansuji of a `float64` in simplified Chinese style.
//
// `prec` has the same meaning of `strconv.FormatFloat`.
func SimplifiedChineseFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	strParts := strings.Split(str, ".")

	sb.WriteString(simplifiedChineseStringInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point[1])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(numbers[strParts[1][i]-'0'])
		}
	}

	return sb.String()
}

// SimplifiedChineseInt returns kansuji of a `int64` in simplified Chinese style.
func SimplifiedChineseInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[1])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(simplifiedChineseStringInt(str))

	return sb.String()
}

// SimplifiedChineseUint returns kansuji of a `uint64` in simplified Chinese style.
func SimplifiedChineseUint(in uint64) string {
	var sb strings.Builder

	str := strconv.FormatUint(in, 10)
	sb.WriteString(simplifiedChineseStringInt(str))

	return sb.String()
}

// SimplifiedChineseFinancialFloat returns kansuji of a `float64` in simplified Chinese financial style.
//
// `prec` has the same meaning of `strconv.FormatFloat`.
func SimplifiedChineseFinancialFloat(in float64, prec int) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[1])
	}

	str := strconv.FormatFloat(in, 'f', prec, 64)
	strParts := strings.Split(str, ".")

	sb.WriteString(simplifiedChineseStringFinancialInt(strParts[0]))

	if len(strParts) > 1 && strParts[1] != "" {
		sb.WriteRune(point[1])
		for i := 0; i < len(strParts[1]); i++ {
			sb.WriteRune(financialNumbers[strParts[1][i]][1])
		}
	}

	return sb.String()
}

// SimplifiedChineseFinancialInt returns kansuji of a `int64` in simplified Chinese financial style.
func SimplifiedChineseFinancialInt(in int64) string {
	var sb strings.Builder

	isNegative := in < 0
	if isNegative {
		in = -in
		sb.WriteRune(negative[1])
	}

	str := strconv.FormatInt(in, 10)
	sb.WriteString(simplifiedChineseStringFinancialInt(str))

	return sb.String()
}

// SimplifiedChineseFinancialUint returns kansuji of a `uint64` in simplified Chinese financial style.
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

	allzeros := true
	if b == 2 && integer[g][1] == numbers[1] {
		b--
		if b > 0 {
			sb.WriteRune(multipliers[b-1])
		}
		allzeros = false
	}
	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; integer[g][b] == numbers[0]; b-- {
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
			if integer[g][b] == numbers[0] {
				if b == 0 {
					break
				}
				if integer[g][b-1] == numbers[0] {
					continue
				}
				sb.WriteRune(integer[g][b])
				continue
			}
			allzeros = false
			sb.WriteRune(integer[g][b])
			if b > 0 {
				sb.WriteRune(multipliers[b-1])
			}
		}
		b = 3
		if g > 0 && !allzeros {
			sb.WriteRune(multipliers2[g-1][1])
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
		integer[g][b] = financialNumbers[in[i]-'0'][1]
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
	if b == 2 && integer[g][1] == financialNumbers[1][1] {
		b--
		if b > 0 {
			sb.WriteRune(financialMultipliers[b-1][1])
		}
		allzeros = false
	}
	if b > 0 {
		b--
	} else {
		g--
		b = 3
	}
	for ; integer[g][b] == numbers[0]; b-- {
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
			if integer[g][b] == numbers[0] {
				if b == 0 {
					break
				}
				if integer[g][b-1] == numbers[0] {
					continue
				}
				sb.WriteRune(integer[g][b])
				continue
			}
			allzeros = false
			sb.WriteRune(integer[g][b])
			if b > 0 {
				sb.WriteRune(financialMultipliers[b-1][1])
			}
		}
		b = 3
		if g > 0 && !allzeros {
			sb.WriteRune(multipliers2[g-1][1])
		}

		allzeros = true
	}

	return sb.String()
}
