package fsft_test

import (
	"fmt"

	"github.com/xuender/kfont/founder/fsft"
)

func ExampleFont() {
	font := fsft.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright(c) Founder Corporation.1999
}
