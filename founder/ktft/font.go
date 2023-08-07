// Package ktft 楷体繁体.
package ktft

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengKaiTiFanTi-1.ttf
var _font []byte

// Font 方正楷体繁体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(_font)

	return font
}
