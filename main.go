package main

import (
	"strings"
	"jp/tagging"
	"os"
	"fmt"
	"image"
	"image/color"
	"log"
	"image/png"
	"github.com/golang/freetype"
	"io/ioutil"
)

const (
	dx       = 100         // 图片的大小 宽度
	dy       = 40          // 图片的大小 高度
	fontFile = "luximr.ttf" // 需要使用的字体文件
	fontSize = 20          // 字体尺寸
	fontDPI  = 72          // 屏幕每英寸的分辨率
)

func main2() {

	// 需要保存的文件
	imgcounter := 123
	imgfile, _ := os.Create(fmt.Sprintf("%03d.png", imgcounter))
	defer imgfile.Close()

	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))

	// 画背景
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色，依次是 RGBA
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		log.Println(err)
		return
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
		return
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	pt := freetype.Pt(10, 10+int(c.PointToFixed(fontSize)>>8)) // 字出现的位置

	_, err = c.DrawString("ABCDE", pt)
	if err != nil {
		log.Println(err)
		return
	}

	// 以PNG格式保存文件
	err = png.Encode(imgfile, img)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	words := []string{
		"食べる,たべる,吃",
		"分かる,わかる,了解",
		"見る,みる,看",
		"寝る,ねる,睡",
		"起きる,おきる,叫醒,发生",
		"考える,かんがえる,考虑",
		"教える,おしえる,教,告知",
		"出る,でる, 出来",
		"いる,いる,存在（有生命）",
		"着る,きる,穿",
		"話す,はなす,说",
		"聞く,きく,问,听",
		"泳ぐ,およぐ,游泳",
		"遊ぶ,あそぶ,玩",
		"待つ,まつ,等待",
		"飲む,のむ,喝",
		"買う,かう,买",
		"ある,ある,存在（无生命）",
		"死ぬ,しぬ, 死",
	}

	for _, word := range words {
		splitd := strings.Split(word, ",")
		if len(splitd) < 3 {
			panic("至少需要输入 单词,假名,中文释义")
		}
		//eatVerb := verb.IdentifyVerbWord(splitd[0], splitd[1], splitd[2])
		//fmt.Printf("%+v", eatVerb)
	}

	err := tagging.InitKVDict("/Users/Bernie/Codes/GoWorkspace/jp/word_csv")
	if err != nil {
		panic(err)
	}

	tagging.IteratorLine("冷雨が降季節も気味若さもささずに")
	//tagging.IteratorLine("冷雨が降季節も気味若さもささずに")

}
