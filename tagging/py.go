package tagging

import (
	"strings"
	"io/ioutil"
	"bytes"
	"fmt"
)

var Words = make(map[string]string, 0)

func InitKVDict(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(bytes), "\n") {
		k, v := KVFromCsvLine(line)
		Words[k] = v
	}
	return nil
}

func KVFromCsvLine(line string) (string, string) {
	kv := strings.Split(string([]rune(line)), ",")
	return kv[0], kv[1]
}

func IteratorLine(line string) {
	var (
		jia      bytes.Buffer
		han      bytes.Buffer
		pair     = make(map[string]string)
		pairKeys = make([]string, 0)
	)
	idx := 0
	words := []rune(line)
	for _, word := range strings.Split(string(words), "") {
		ws, jias := "", ""
		minWords := words[:idx]
		for i := 0; i <= len(minWords); i ++ {
			for j := 0; j < i; j++ {
				minWord := string(minWords[j : i+1])
				if v, ok := Words[minWord]; ok {
					ws = minWord
					jias = v
				} else {
					ws = word
					jias = word
				}
			}
		}

		if v, ok := Words[word]; ok {
			ws = word
			jias = v
		}
		if jias != "" {
			for k, _ := range pair {
				if strings.Contains(ws, k) {
					delete(pair, k)
					pair[ws] = jias
				}
			}
		}
		pair[ws] = jias
		pairKeys = append(pairKeys, ws)
		idx += 1
	}
	for _, k := range pairKeys {
		jiaWidth, wordWidth := 0, 0
		jiaWord := pair[k]
		if jiaWord != "" {
			jiaWidth = len([]rune(jiaWord))
			wordWidth = len([]rune(k))
			if jiaWidth > wordWidth {
				space := nSpace(jiaWidth - wordWidth)
				jia.WriteString(jiaWord)
				han.WriteString(space)
				han.WriteString(k)
			} else if jiaWidth == wordWidth && jiaWord != k{
				jia.WriteString(jiaWord)
				han.WriteString(k)
			} else {
				jia.WriteString("　")
				han.WriteString(k)
			}
			if jiaWord != k {
				jia.WriteString("·")
				han.WriteString(" ")
			}
		}
	}

	fmt.Printf("%s\n", jia.String())
	fmt.Printf("%s\n", han.String())
}

func nSpace(n int) string {
	space := ""
	for i := 0; i < n; i ++ {
		space += fmt.Sprintf("　")
	}
	return space
}
