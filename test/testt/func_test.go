package testt

import (
	"fmt"
	"regexp"
	"testing"
)

var haystack2 = `type CarrierMoneyRatio struct { // ИД`

var pattern = regexp.MustCompile("type (.+?) struct {")

func Test1(t1 *testing.T) {
	data := pattern.FindStringSubmatch(haystack2)
	if data == nil {
		fmt.Println(data)
	}
	for i := range data {
		fmt.Println(data[i])
	}
}

var pattern2 = regexp.MustCompile(`[[:graph:]]|\s|[а-яА-ЯёЁ№]+`)

func Test2(t *testing.T) {
	haystack := "aa  ˜  adfd Ü tuy ₽ thgfh6*,.~^&$%#Ёё!№ .ю, э ' ~ ` _ абырвалк � Ж ::;-+=*/˚"
	flag := pattern2.MatchString(haystack)
	t.Log(flag)
	str := pattern2.ReplaceAllString(haystack, "")
	t.Log("-" + str + "-")
}

var pattern3 = regexp.MustCompile(`json:"(.+?)"`)

func TestReq(t *testing.T) {
	haystack := `hfghgfhfg gfhfgh fhfg json:"status" dfdsf json:"status" nfghfghfg`
	str := pattern3.ReplaceAllString(haystack, "")
	t.Log("-" + str + "-")
}
