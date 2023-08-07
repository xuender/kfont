// Package htjt 黑体简体.
package htjt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengHeiTiJianTi-1.ttf
var _font []byte

// Font 方正黑体简体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(_font)

	return font
}
