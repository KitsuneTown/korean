package korean

import (
	"strings"
	"unicode/utf8"
)

func Josa(str, target string) string {
	jong := HasJongSeong(str)

	if jong {
		return getJosa(target, jongSeongJosa)
	} else {
		return getJosa(target, noJongSeongJosa)
	}
}

var (
	noJongSeongJosa = []string{
		"나", "란", "든가", "든지", "나마", "네",
		"를", "가", "와", "로", "는", "아",
	}
	jongSeongJosa = []string{
		"을", "과", "으로", "은", "야",
		"이나", "이란", "이든가", "이든지", "이나마", "이네", "이",
	}
)

func getJosa(target string, josa []string) string {
	for _, v := range josa {
		if josa := strings.Replace(v, "이", "", -1); utf8.RuneCountInString(josa) > 0 {
			if strings.Contains(target, josa) {
				return v
			}
			continue
		}

		if strings.Contains(target, v) {
			return v
		}
	}

	return ""
}
