package ssjt_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/ssjt"
)

func ExampleFont() {
	font := ssjt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.2000
}
