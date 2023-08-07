package ym_test

import (
	"fmt"

	"github.com/xuender/kfont/tsanger/ym"
)

func ExampleFont() {
	font := ym.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright © Beijing Tsanger Character Technology Co., Ltd.. All rights reserved.
}
