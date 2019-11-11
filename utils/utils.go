package utils

// a, i, u, e, o
var KEYS = [][]string{
	[]string{"あ", "い", "う", "え", "お"},
	[]string{"か", "き", "く", "け", "こ"},
	[]string{"さ", "し", "す", "せ", "そ"},
	[]string{"た", "ち", "つ", "て", "と"},
	[]string{"な", "に", "ぬ", "ね", "の"},
	[]string{"は", "ひ", "ふ", "へ", "ほ"},
	[]string{"ま", "み", "む", "め", "も"},
	[]string{"や", "い", "ゆ", "え", "よ"},
	[]string{"ら", "り", "る", "れ", "ろ"},
	[]string{"わ", "い", "う", "え", "を"},
	[]string{"が", "ぎ", "ぐ", "げ", "ご"},
	[]string{"ざ", "じ", "ぜ", "ず", "ぞ"},
	[]string{"だ", "ぢ", "ず", "ぜ", "ぞ"},
	[]string{"ば", "び", "ぶ", "べ", "ぼ"},
	[]string{"ぱ", "ぴ", "ぷ", "ぺ", "ぽ"},
}

type Tone struct {
	Current string
	row     int
	column  int
}

func NewTone(tone string) *Tone {
	tt := &Tone{}
	for ri, row := range KEYS {
		for ci, word := range row {
			if word == tone {
				tt.Current = tone
				tt.row = ri
				tt.column = ci
			}
		}
	}
	return tt
}

//判断此假名是否在 i, e 段
func (t *Tone) IsIE() bool {
	return t.column == 1 || t.column == 3
}

func (t *Tone) String() string {
	return t.Current
}

//取到假名的同行i段假名
func (tone *Tone) Part(i int) Tone {
	c := KEYS[tone.row][i]
	return Tone{
		Current: c,
		row:     tone.row,
		column:  i,
	}
}
