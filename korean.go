package korean

type Text string

func (t Text) HasJongSeong() bool                          { return HasJongSeong(string(t)) }
func (t Text) HasJongSeongArray() []bool                   { return HasJongSeongArray(string(t)) }
func (t Text) GetSyllables(opt ...SyllableOption) []string { return GetSyllables(string(t), opt...) }
func (t Text) GetSyllableArray(opt ...SyllableOption) [][]string {
	return GetSyllableArray(string(t), opt...)
}

func (t Text) Josa(josa string) string { return Josa(string(t), josa) }

func (t Text) IsHangul(opt ...bool) bool        { return IsHangul(string(t), opt...) }
func (t Text) IsHangulArray(opt ...bool) []bool { return IsHangulArray(string(t), opt...) }
