// Package ktjt 楷体简体.
package ktjt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengKaiTiJianTi-1.ttf
var Bytes []byte

// Font 方正楷体简体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
