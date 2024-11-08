// Package fsft 仿宋繁体.
package fsft

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengFangSongFanTi-1.ttf
var Bytes []byte

// Font 方正仿宋繁体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
