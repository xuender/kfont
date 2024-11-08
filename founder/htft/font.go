// Package htft 黑体繁体.
package htft

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed FangZhengHeiTiFanTi-1.ttf
var Bytes []byte

// Font 方正黑体繁体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
