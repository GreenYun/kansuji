package kansuji

var numbers = [...]rune{
	0x96f6, // `零`
	0x4e00, // `一`
	0x4e8c, // `二`
	0x4e09, // `三`
	0x56db, // `四`
	0x4e94, // `五`
	0x516d, // `六`
	0x4e03, // `七`
	0x516b, // `八`
	0x4e5d, // `九`
}

var multipliers = [...]rune{
	0x5341, // `十`
	0x767e, // `百`
	0x5343, // `千`
}

var multipliers2 = [...][3]rune{
	{0x842c, 0x4e07, 0x4e07}, // `萬`, `万`, `万`
	{0x5104, 0x4ebf, 0x5104}, // `億`, `亿`, `億`
	{0x5146, 0x5146, 0x5146}, // `兆`, `兆`, `兆`
	{0x4eac, 0x4eac, 0x4eac}, // `京`, `京`, `京`
	{0x5793, 0x5793, 0x5793}, // `垓`, `垓`, `垓`
	{0x79ed, 0x79ed, 0x79ed}, // `秭`, `秭`, `秭`
	{0x7a70, 0x7a70, 0x7a70}, // `穰`, `穰`, `穰`
	{0x6e9d, 0x6c9f, 0x6e9d}, // `溝`, `沟`, `溝`
	{0x6f97, 0x6da7, 0x6f97}, // `澗`, `涧`, `澗`
	{0x6b63, 0x6b63, 0x6b63}, // `正`, `正`, `正`
	{0x8f09, 0x8f7d, 0x8f09}, // `載`, `载`, `載`
	{0x6975, 0x6781, 0x6975}, // `極`, `极`, `極`
}

var financialNumbers = [...][3]rune{
	{0x96f6, 0x96f6, 0x96f6}, // `零`, `零`, `零`
	{0x58f9, 0x58f9, 0x58f1}, // `壹`, `壹`, `壱`
	{0x8cb3, 0x8d30, 0x5f10}, // `貳`, `贰`, `弐`
	{0x53c3, 0x53c1, 0x53c2}, // `參`, `叁`, `参`
	{0x8086, 0x8086, 0x8086}, // `肆`, `肆`, `肆`
	{0x4f0d, 0x4f0d, 0x4f0d}, // `伍`, `伍`, `伍`
	{0x9678, 0x9646, 0x9678}, // `陸`, `陆`, `陸`
	{0x67d2, 0x67d2, 0x6f06}, // `柒`, `柒`, `漆`
	{0x634c, 0x634c, 0x634c}, // `捌`, `捌`, `捌`
	{0x7396, 0x7396, 0x7396}, // `玖`, `玖`, `玖`
}

var financialMultipliers = [...][3]rune{
	{0x62fe, 0x62fe, 0x62fe}, // `拾`, `拾`, `拾`
	{0x4f70, 0x4f70, 0x964c}, // `佰`, `佰`, `陌`
	{0x4edf, 0x4edf, 0x9621}, // `仟`, `仟`, `阡`
}

var point = [3]rune{0x9ede, 0x70b9, 0x70b9} // `點`, `点`, `点`

var negative = [3]rune{0x8ca0, 0x8d1f, 0x8ca0} // `負`, `负`, `負`

const positive = rune(0x6b63) // `正`
