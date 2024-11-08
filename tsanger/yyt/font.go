// Package yyt 渔阳体W02.
package yyt

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed CangErYuYangTiW02-2.ttf
var Bytes []byte

// Font 仓耳道渔阳体W02.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
