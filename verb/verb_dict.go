package verb

import (
	"jp/utils"
	"strings"
	"fmt"
)

const (
	I_VERB  = 1 //一段动词
	U_VERB  = 2 //五段动词
	KA_VERB = 3
	SA_VERB = 4
	A       = 0
	I       = 1
	U       = 2
	E       = 3
	O       = 4
)

var (
	I_LAST_KEYS = []string{"く", "ぐ", "す", "つ", "ぬ", "ぶ", "む", "る", "う"}
)

func IdentifyVerbWord(word, jia, ch string) *VerbWord {
	keys := []rune(word)
	last := keys[len(keys)-1]

	verb := &VerbWord{
		WordChars: keys,
		Word:      string(keys),
		last:      string(last),
		Jia:       jia,
		Ch:        ch,
	}

	verb.verbClassifer()
	return verb
}

type VerbDeformation interface {
}

//动词
type VerbWord struct {
	WordChars []rune
	Word      string
	last      string
	Comment   string //解释说明
	Jia       string //假名
	Ch        string //中文释义
	Tpe       int    // 分类
}

func (verb *VerbWord) String() string {
	return fmt.Sprintf("%s[%s]\t: %-15s\tます:%-10s\tない:%s\n", verb.Word,
		verb.Jia,
		verb.Ch,
		verb.Masu().Word,
		verb.Nai().Word,
	)
}

//分类:
//五段（一类动词）
//一段（二类动词）
//カ变、サ变（三类动词）
func (verb *VerbWord) verbClassifer() {
	//TODO 增加那些特殊的五段动词
	//一段动词
	// 结尾为 る 并且 前一个假名在 i 段或者 e 段
	if verb.last == "る" {
		//前一个是汉字时, 取其假名做判断
		jias := strings.Split(verb.Jia, "")
		ie := jias[len(jias)-2]
		if utils.NewTone(ie).IsIE() {
			verb.Tpe = I_VERB
		} else {
			verb.Tpe = U_VERB
		}
	} else {
		//不以 る 结尾的 + 非一段的特殊几个
		verb.Tpe = U_VERB
	}
}

func (verb VerbWord) SplitLastChar() (string, string) {
	last := verb.WordChars[len(verb.WordChars)-1]
	pre := verb.WordChars[0 : len(verb.WordChars)-1]
	return string(pre), string(last)
}

//ます
func (verb VerbWord) Masu() VerbWord {
	masu := "ます"
	pre, last := verb.SplitLastChar()
	if verb.Tpe == U_VERB {
		//将结尾ウ段假名变成同行的イ段假名+ます
		tail := utils.NewTone(last).Part(I).String() + masu
		word := pre + tail
		return *IdentifyVerbWord(word, "", "")
	}

	if verb.Tpe == I_VERB {
		//去掉る＋ます
		word := pre + masu
		return *IdentifyVerbWord(word, "", "")
	}
	return VerbWord{}
}

//ない
func (verb VerbWord) Nai() VerbWord {
	nai := "ない"
	pre, last := verb.SplitLastChar()
	if verb.Tpe == U_VERB {
		//将结尾ウ段假名变成同行的ア段假名+ない
		tail := utils.NewTone(last).Part(A).String() + nai
		word := pre + tail
		return *IdentifyVerbWord(word, "", "")
	}
	if verb.Tpe == I_VERB {
		//去掉る＋ます
		word := pre + nai
		return *IdentifyVerbWord(word, "", "")
	}
	return VerbWord{}
}
