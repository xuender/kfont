package xxt_test

import (
	"fmt"

	"github.com/xuender/kfont/pmzd/xxt"
)

func ExampleFont() {
	font := xxt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// (c) copyright PangMenZhengDao 2021
}
