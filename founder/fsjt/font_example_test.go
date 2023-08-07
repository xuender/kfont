package fsjt_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/fsjt"
)

func ExampleFont() {
	font := fsjt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.2000
}
