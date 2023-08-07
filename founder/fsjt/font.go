// Package fsjt 仿宋简体.
package fsjt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengFangSongJianTi-1.ttf
var _font []byte

// Font 方正仿宋简体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(_font)

	return font
}
