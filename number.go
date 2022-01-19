package korean

type NumberOption struct {
	// 앞에 "일"을 붙여 읽습니다.
	// 가령, 1000을 "일천"으로,
	// 10000을 "일만"으로 반환합니다.
	MarkNumberOne bool
}

func Number(num int, opt ...NumberOption) []string {
	if num == 0 {
		return []string{"영"}
	}

	var option NumberOption
	if len(opt) > 0 {
		option = opt[0]
	} else {
		option = NumberOption{false}
	}

	result := make([]string, 0)

	slicedNum := make([]int, 0)

	n := num

	for n > 0 {
		slicedNum = append(slicedNum, n%10000)
		n /= 10000
	}

	for i, v := range slicedNum {
		_thousand := v / 1000
		_hundred := (v - (_thousand * 1000)) / 100
		_ten := (v - (_thousand * 1000) - (_hundred * 100)) / 10
		_one := v - (_thousand * 1000) - (_hundred * 100) - (_ten * 10)

		var (
			thousand, hundred, ten, one, digits string
		)

		if r := getKoreanNumber(_thousand); r != "" {
			if !option.MarkNumberOne && r == "일" {
				r = ""
			}
			thousand = r + "천"
		}

		if r := getKoreanNumber(_hundred); r != "" {
			if !option.MarkNumberOne && r == "일" && _thousand <= 0 {
				r = ""
			}
			hundred = r + "백"
		}

		if r := getKoreanNumber(_ten); r != "" {
			if r == "일" {
				r = ""
			}
			ten = r + "십"
		}

		if r := getKoreanNumber(_one); r != "" {
			one = r
		}

		digits = getKoreanNumberDigits(i)
		if thousand == "" && hundred == "" && ten == "" && one == "" {
			digits = ""
		}

		result = append(
			result,
			thousand+hundred+ten+one+digits,
		)
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func getKoreanNumberDigits(i int) string {
	switch i {
	case 1:
		return "만"
	case 2:
		return "억"
	case 3:
		return "조"
	case 4:
		return "경"
	default:
		return ""
	}
}

func getKoreanNumber(n int) string {
	switch n {
	case 1:
		return "일"
	case 2:
		return "이"
	case 3:
		return "삼"
	case 4:
		return "사"
	case 5:
		return "오"
	case 6:
		return "육"
	case 7:
		return "칠"
	case 8:
		return "팔"
	case 9:
		return "구"
	}
	return ""
}
