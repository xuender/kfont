package htft_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/htft"
)

func ExampleFont() {
	font := htft.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.1999
}
