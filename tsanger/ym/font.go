// Package ym 与墨W02.
package ym

import (
	_ "embed"

	"golang.org/x/image/font/opentype"
)

//go:embed CangErYuMoW02-2.ttf
var Bytes []byte

// Font 仓耳道与墨W02.
func Font() *opentype.Font {
	font, _ := opentype.Parse(Bytes)

	return font
}
