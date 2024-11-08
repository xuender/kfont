// Package ssjt 书宋简体.
package ssjt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengShuSongJianTi-1.ttf
var Bytes []byte

// Font 方正书宋简体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
