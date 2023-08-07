package yyt_test

import (
	"fmt"

	"github.com/xuender/kfont/tsanger/yyt"
)

func ExampleFont() {
	font := yyt.Font()
	name, _ := font.Name(nil, 0)

	fmt.Println(name)

	// Output:
	// Copyright © Beijing Tsanger Character Technology Co., Ltd.. All rights reserved.
}
