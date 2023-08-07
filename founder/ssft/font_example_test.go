package ssft_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/ssft"
)

func ExampleFont() {
	font := ssft.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.1999
}
