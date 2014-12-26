package main

import (
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

func NewPlaceholder() *Placeholder {
	return &Placeholder{
		width:     512,
		height:    512,
		bgColor:   "#cbcbcb",
		fontColor: "#999999",
		format:    "png",
		text:      "TEXT",
		filename:  "OUT.png",
	}
}

func (p *Placeholder) Generate() {
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
	app.Version = "0.1.0"
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
		cli.StringFlag{
			Name:  "background",
			Value: "#cbcbcb",
			Usage: "Background color in hex (#RRGGBB) format",
		},
		cli.StringFlag{
			Name:  "foreground",
			Value: "#999999",
			Usage: "Foreground color in hex (#RRGGBB) format",
		},
		cli.StringFlag{
			Name:  "format",
			Value: "png",
			Usage: "File format of placeholder",
		},
		cli.StringFlag{
			Name:  "text, t",
			Value: "TEXT",
			Usage: "Text of placeholder",
		},
		cli.StringFlag{
			Name:  "output, o",
			Value: "OUT",
			Usage: "Output file",
		},
	}
	app.Action = func(c *cli.Context) {
		p := NewPlaceholder()
		if width := c.Int("width"); width > 0 {
			p.width = uint(width)
		}
		if height := c.Int("height"); height > 0 {
			p.height = uint(height)
		}
		if bgColor := c.String("background"); bgColor != "" {
			p.bgColor = bgColor
		}
		if fontColor := c.String("foreground"); fontColor != "" {
			p.fontColor = fontColor
		}
		if format := c.String("format"); format != "" {
			p.format = format
		}
		if text := c.String("text"); text != "" {
			p.text = text
		}
		if output := c.String("output"); output != "" {
			p.filename = output + "." + p.format
		}
		p.Generate()
	}
	app.Run(os.Args)
}
