package ktjt_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/ktjt"
)

func ExampleFont() {
	font := ktjt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.2000
}
