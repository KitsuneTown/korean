package korean

func HasJongSeong(str string) bool {
	runes := []rune(str)
	return (runes[len(runes)-1]-0xAC00)%28 > 0
}

func HasJongSeongArray(str string) (result []bool) {
	runes := []rune(str)

	for _, v := range runes {
		result = append(result, (v-0xAC00)%28 > 0)
	}
	return
}

type SyllableOption struct {
	// 조합형 한글만을 배열에 담습니다.
	OnlyCombinedHangul bool
	// 타 언어 및 문자 부호도 포함할지 선택합니다.
	IncludeOtherLanguage bool
	// 쌍자음 및 조합형 모음을 분리할지 선택합니다.
	// 가령, "ㄲ" 는 "ㄱ", "ㄱ"으로,
	// "ㅝ"는 "ㅜ"와 "ㅓ"로 분리됩니다.
	SeparateHangul bool
}

func GetSyllables(str string, opt ...SyllableOption) []string {
	syllable := GetSyllableArray(str, opt...)
	result := make([]string, 0)

	for _, v := range syllable {
		result = append(result, v...)
	}

	return result
}

func GetSyllableArray(str string, opt ...SyllableOption) [][]string {
	result := make([][]string, 0)

	runes := []rune(str)

	var option SyllableOption
	if len(opt) > 0 {
		option = opt[0]
	} else {
		option = SyllableOption{false, false, false}
	}

	var (
		CHOSEONG = []string{
			"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ",
			"ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ",
			"ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ",
			"ㅋ", "ㅌ", "ㅍ", "ㅎ",
		}
		JUNGSEONG = []string{
			"ㅏ", "ㅐ", "ㅑ", "ㅒ", "ㅓ",
			"ㅔ", "ㅕ", "ㅖ", "ㅗ", "ㅘ",
			"ㅙ", "ㅚ", "ㅛ", "ㅜ", "ㅝ",
			"ㅞ", "ㅟ", "ㅠ", "ㅡ", "ㅢ",
			"ㅣ",
		}
		JONGSEONG = []string{
			"", "ㄱ", "ㄲ", "ㄳ", "ㄴ",
			"ㄵ", "ㄶ", "ㄷ", "ㄹ", "ㄺ",
			"ㄻ", "ㄼ", "ㄽ", "ㄾ", "ㄿ",
			"ㅀ", "ㅁ", "ㅂ", "ㅄ", "ㅅ",
			"ㅆ", "ㅇ", "ㅈ", "ㅊ", "ㅋ",
			"ㅌ", "ㅍ", "ㅎ",
		}
	)

	for _, v := range runes {
		if (option.OnlyCombinedHangul && !IsHangul(string(v), true)) ||
			!IsHangul(string(v)) {
			if option.IncludeOtherLanguage {
				result = append(result, []string{string(v)})
			}
			continue
		}

		uniChar := float64(v)

		runeResult := make([]string, 0)

		choIndex := (uniChar - 0xAC00) / (21 * 28)
		if choIndex >= 0 && choIndex <= 18 {
			runeResult = append(runeResult, CHOSEONG[int(choIndex)])
		}

		jungIndex := ((int(uniChar) - 0xAC00) % (21 * 28)) / 28
		if jungIndex >= 0 && jungIndex <= 20 {
			runeResult = append(runeResult, JUNGSEONG[int(jungIndex)])
		}

		jongIndex := (int(uniChar) - 0xAC00) % 28
		if jongIndex > 0 {
			runeResult = append(runeResult, JONGSEONG[jongIndex])
		}

		result = append(result, runeResult)
	}

	if option.SeparateHangul {
		r := make([][]string, 0)

		for _, v := range result {
			ur := make([]string, 0)
			for _, c := range v {
				switch c {
				case "ㄲ":
					ur = append(ur, []string{"ㄱ", "ㄱ"}...)
				case "ㄸ":
					ur = append(ur, []string{"ㄷ", "ㄷ"}...)
				case "ㅃ":
					ur = append(ur, []string{"ㅂ", "ㅂ"}...)
				case "ㅆ":
					ur = append(ur, []string{"ㅅ", "ㅅ"}...)
				case "ㅉ":
					ur = append(ur, []string{"ㅈ", "ㅈ"}...)
				case "ㅘ":
					ur = append(ur, []string{"ㅗ", "ㅏ"}...)
				case "ㅙ":
					ur = append(ur, []string{"ㅗ", "ㅐ"}...)
				case "ㅚ":
					ur = append(ur, []string{"ㅗ", "ㅣ"}...)
				case "ㅝ":
					ur = append(ur, []string{"ㅜ", "ㅓ"}...)
				case "ㅞ":
					ur = append(ur, []string{"ㅜ", "ㅔ"}...)
				case "ㅟ":
					ur = append(ur, []string{"ㅜ", "ㅣ"}...)
				case "ㅢ":
					ur = append(ur, []string{"ㅡ", "ㅣ"}...)
				case "ㄳ":
					ur = append(ur, []string{"ㄱ", "ㅅ"}...)
				case "ㄵ":
					ur = append(ur, []string{"ㄴ", "ㅈ"}...)
				case "ㄶ":
					ur = append(ur, []string{"ㄴ", "ㅎ"}...)
				case "ㄺ":
					ur = append(ur, []string{"ㄹ", "ㄱ"}...)
				case "ㄻ":
					ur = append(ur, []string{"ㄹ", "ㅁ"}...)
				case "ㄼ":
					ur = append(ur, []string{"ㄹ", "ㅂ"}...)
				case "ㄽ":
					ur = append(ur, []string{"ㄹ", "ㅅ"}...)
				case "ㄾ":
					ur = append(ur, []string{"ㄹ", "ㅌ"}...)
				case "ㄿ":
					ur = append(ur, []string{"ㄹ", "ㅍ"}...)
				case "ㅀ":
					ur = append(ur, []string{"ㄹ", "ㅎ"}...)
				case "ㅄ":
					ur = append(ur, []string{"ㅂ", "ㅅ"}...)
				default:
					ur = append(ur, c)
				}
			}
			r = append(r, ur)
		}

		return r
	}

	return result
}
