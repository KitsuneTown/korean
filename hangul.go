package korean

func IsHangul(str string, onlyCombined ...bool) bool {
	runes := []rune(str)

	if len(onlyCombined) <= 0 || (len(onlyCombined) > 0 && onlyCombined[0]) {
		return (runes[len(runes)-1] >= 0x1100 && runes[len(runes)-1] <= 0x11FF) ||
			(runes[len(runes)-1] >= 0x3130 && runes[len(runes)-1] <= 0x318F) ||
			(runes[len(runes)-1] >= 0xAC00 && runes[len(runes)-1] <= 0xD7A3)
	}

	return runes[len(runes)-1] >= 0xAC00 && runes[len(runes)-1] <= 0xD7A3
}

func IsHangulArray(str string, onlyCombined ...bool) []bool {
	runes := []rune(str)
	result := make([]bool, 0)

	for _, v := range runes {
		if len(onlyCombined) <= 0 || (len(onlyCombined) > 0 && onlyCombined[0]) {
			result = append(
				result,
				(v >= 0x1100 && v <= 0x11FF) ||
					(v >= 0x3130 && v <= 0x318F) ||
					(v >= 0xAC00 && v <= 0xD7A3),
			)
			continue
		}
		result = append(result, v >= 0xAC00 && v <= 0xD7A3)
	}

	return result
}
