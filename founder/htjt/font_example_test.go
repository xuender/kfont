package htjt_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/htjt"
)

func ExampleFont() {
	font := htjt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.2000
}
