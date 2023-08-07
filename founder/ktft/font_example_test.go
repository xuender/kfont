package ktft_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/ktft"
)

func ExampleFont() {
	font := ktft.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.1999
}
