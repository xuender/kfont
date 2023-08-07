package btt_test

import (
	"fmt"

	"github.com/xuender/kfont/pmzd/btt"
)

func ExampleFont() {
	font := btt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// (c) PangMenZhengDao HuXiaoBo 2016
}
