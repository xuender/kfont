// Package ssft 书宋繁体.
package ssft

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengShuSongFanTi-1.ttf
var _font []byte

// Font 方正书宋繁体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(_font)

	return font
}
