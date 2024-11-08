// Package xxt 细线体.
package xxt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed PangMenZhengDaoXiXianTi-2.ttf
var Bytes []byte

// Font 旁门正道细线体.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
