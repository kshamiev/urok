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
