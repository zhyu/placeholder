package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/gographics/imagick/imagick"
	"os"
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
	if p.bgColor == "" {
		p.bgColor = "#cbcbcb"
	}
	if p.fontColor == "" {
		p.fontColor = "#999999"
	}
	if p.format == "" {
		p.format = "png"
	}
	if p.text == "" {
		p.text = fmt.Sprintf("%d x %d", p.width, p.height)
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
	dw.SetFontSize(80)
	dw.SetGravity(imagick.GRAVITY_CENTER)

	mw.AnnotateImage(dw, 0, 0, 0, p.text)
	mw.SetImageFormat(p.format)
	mw.WriteImage(p.filename)
}

func main() {
	app := cli.NewApp()
	app.Name = "placeholder"
	app.Usage = "generate placeholder images"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "width",
			Value: 512,
			Usage: "Width of placeholder",
		},
		cli.IntFlag{
			Name:  "height",
			Value: 512,
			Usage: "Height of placeholder",
		},
	}
	app.Action = func(c *cli.Context) {
		p := new(Placeholder)
		if width := c.Int("width"); width > 0 {
			p.width = uint(width)
		}
		if height := c.Int("height"); height > 0 {
			p.height = uint(height)
		}
		p.Generate()
	}
	app.Run(os.Args)
}
