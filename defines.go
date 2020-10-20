package kansuji

func chars() map[byte]rune {
	return map[byte]rune{
		'0': []rune("零")[0],
		'1': []rune("一")[0],
		'2': []rune("二")[0],
		'3': []rune("三")[0],
		'4': []rune("四")[0],
		'5': []rune("五")[0],
		'6': []rune("六")[0],
		'7': []rune("七")[0],
		'8': []rune("八")[0],
		'9': []rune("九")[0],
	}
}

func multipliers() []rune { return []rune("十百千") }

func multipliersAgain() [][]rune {
	return [][]rune{
		nil,
		[]rune("萬万万"),
		[]rune("億亿"),
		[]rune("兆"),
		[]rune("京"),
		[]rune("垓"),
		[]rune("秭"),
		[]rune("穰"),
		[]rune("溝沟"),
		[]rune("澗涧"),
		[]rune("正"),
		[]rune("載载"),
		[]rune("極极"),
	}
}

func upperChars() map[byte][]rune {
	return map[byte][]rune{
		'0': []rune("零"),
		'1': []rune("壹壹壱"),
		'2': []rune("貳贰弐"),
		'3': []rune("參叁参"),
		'4': []rune("肆"),
		'5': []rune("伍"),
		'6': []rune("陸陆"),
		'7': []rune("柒柒漆"),
		'8': []rune("捌"),
		'9': []rune("玖"),
	}
}

func upperMultipliers() [][]rune {
	return [][]rune{
		[]rune("拾拾拾"),
		[]rune("佰佰陌"),
		[]rune("仟仟阡"),
	}
}

func point() []rune { return []rune("點点点") }

func negative() []rune { return []rune("負负") }

const positive = "正"

type group [4]rune
