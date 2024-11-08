// Package fsjt 仿宋简体.
package fsjt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengFangSongJianTi-1.ttf
var Bytes []byte

// Font 方正仿宋简体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
