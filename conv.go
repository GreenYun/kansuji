package kansuji

import "strings"

func ToHongKongStandard(in string) string {
	return strings.ReplaceAll(in, "參", "叄")
}

func ToCantoneseVariant(in string) string {
	s := strings.ReplaceAll(in, "貳拾", "廿")
	s = strings.ReplaceAll(s, "贰拾", "廿")
	s = strings.ReplaceAll(s, "貮拾", "廿")
	s = strings.ReplaceAll(s, "弐拾", "廿")
	s = strings.ReplaceAll(s, "二十", "廿")
	s = strings.ReplaceAll(s, "參拾", "卅")
	s = strings.ReplaceAll(s, "叄拾", "卅")
	s = strings.ReplaceAll(s, "叁拾", "卅")
	s = strings.ReplaceAll(s, "弎拾", "卅")
	s = strings.ReplaceAll(s, "三十", "卅")
	s = strings.ReplaceAll(s, "肆拾", "卌")
	s = strings.ReplaceAll(s, "四十", "卌")

	return s
}

func ToJapaneseLaw(in string) string {
	s := strings.ReplaceAll(in, "一", "壱")
	s = strings.ReplaceAll(s, "二", "弐")
	s = strings.ReplaceAll(s, "三", "参")
	s = strings.ReplaceAll(s, "十", "拾")

	return s
}
