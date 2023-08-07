// Package btt 标题体.
package btt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed PangMenZhengDaoBiaoTiTi-1.ttf
var _font []byte

// Font 旁门正道标题体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(_font)

	return font
}
