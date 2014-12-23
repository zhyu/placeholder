package main

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
)

type Placeholder struct {
	width     uint
	height    uint
	bgColor   string
	fontColor string
	format    string
	text      string
	filename  string
}

func (p *Placeholder) Generate() {
	if p.width == 0 {
		p.width = 512
	}
	if p.height == 0 {
		p.height = 512
	}
	if p.bgColor == "" {
		p.bgColor = "#cccccc"
	}
	if p.fontColor == "" {
		p.fontColor = "#3a7ba8"
	}
	if p.format == "" {
		p.format = "png"
	}
	if p.text == "" {
		p.text = fmt.Sprintf("%dx%d", p.width, p.height)
	}
	if p.filename == "" {
		p.filename = p.text + "." + p.format
	}

	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	pw := imagick.NewPixelWand()
	defer pw.Destroy()

	dw := imagick.NewDrawingWand()
	defer dw.Destroy()

	pw.SetColor(p.bgColor)
	mw.NewImage(p.width, p.height, pw)

	pw.SetColor(p.fontColor)
	dw.SetFillColor(pw)
	dw.SetFont("/Library/Fonts/Georgia.ttf")
	dw.SetFontSize(24)
	dw.SetGravity(imagick.GRAVITY_CENTER)

	mw.AnnotateImage(dw, 0, 0, 0, p.text)
	mw.SetImageFormat(p.format)
	mw.WriteImage(p.filename)
}

func main() {
	p := new(Placeholder)
	p.Generate()
}
