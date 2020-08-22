package main

import (
	"image"
	"image/color"
	"image/draw"
	"math"
	"math/rand"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"

	"github.com/status-im/avatars/colours"
	"github.com/status-im/avatars/dimentions"
	"github.com/status-im/avatars/hair"
)

var (
	upLeft   = image.Point{X: 0, Y: 0}
	lowRight = image.Point{X: dimentions.Width, Y: dimentions.Height}
)

func main() {
	exampleSkin()
	exampleHairColour()
	exampleHairStyles()
	exampleEyes()
}

func randomExamples() {
	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight.Mul(8)})

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

			drawAvatar(
				img,
				colours.Rand(colours.Skin),
				colours.Rand(colours.Hair),
				colours.Rand(colours.Eye),
				hair.Styles[rand.Intn(len(hair.Styles))],
			)

			minP := image.Point{X: upLeft.X + (i)*dimentions.Width, Y: upLeft.Y + (j)*dimentions.Height}
			maxP := image.Point{X: lowRight.X + (i)*dimentions.Width, Y: lowRight.Y + (j)*dimentions.Height}
			rect := image.Rectangle{Min: minP, Max: maxP}

			draw.Draw(rgba, rect, img, image.Point{X: 0, Y: 0}, draw.Src)
		}
	}

	draw2dimg.SaveToPngFile("examples/avatars.png", rgba)
}

func exampleSkin() {
	nLowRight := lowRight
	nLowRight.X = nLowRight.X * len(colours.Skin)
	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: nLowRight})

	for i, s := range colours.Skin {
		img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

		drawAvatar(img, s, colours.Hair[0], colours.Black, hair.Styles[1])

		minP := image.Point{X: upLeft.X + (i)*dimentions.Width, Y: upLeft.Y}
		maxP := image.Point{X: lowRight.X + (i)*dimentions.Width, Y: lowRight.Y}
		rect := image.Rectangle{Min: minP, Max: maxP}

		draw.Draw(rgba, rect, img, image.Point{X: 0, Y: 0}, draw.Src)
	}

	draw2dimg.SaveToPngFile("examples/skins.png", rgba)
}

func exampleHairColour() {
	avatarsPerRow := 5

	nLowRight := lowRight
	nLowRight.X = nLowRight.X * len(colours.Hair) / len(colours.Hair) * avatarsPerRow
	nLowRight.Y = nLowRight.Y * int(math.Ceil(float64(len(colours.Hair))/float64(avatarsPerRow)))
	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: nLowRight})

	for i, h := range colours.Hair {
		img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

		drawAvatar(img, colours.Skin[0], h, colours.Black, hair.Styles[1])

		minP := image.Point{X: upLeft.X + (i%5)*dimentions.Width, Y: upLeft.Y + (i%len(colours.Hair)/avatarsPerRow)*dimentions.Height}
		maxP := image.Point{X: lowRight.X + (i%5)*dimentions.Width, Y: lowRight.Y + (i%len(colours.Hair)/avatarsPerRow)*dimentions.Height}
		rect := image.Rectangle{Min: minP, Max: maxP}

		draw.Draw(rgba, rect, img, image.Point{X: 0, Y: 0}, draw.Src)
	}

	draw2dimg.SaveToPngFile("examples/hair.png", rgba)
}

func exampleHairStyles() {
	nLowRight := lowRight
	nLowRight.X = nLowRight.X * len(hair.Styles)
	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: nLowRight})

	for i, s := range hair.Styles {
		img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

		drawAvatar(img, colours.Skin[0], colours.Black, colours.Black, s)

		minP := image.Point{X: upLeft.X + (i)*dimentions.Width, Y: upLeft.Y}
		maxP := image.Point{X: lowRight.X + (i)*dimentions.Width, Y: lowRight.Y}
		rect := image.Rectangle{Min: minP, Max: maxP}

		draw.Draw(rgba, rect, img, image.Point{X: 0, Y: 0}, draw.Src)
	}

	draw2dimg.SaveToPngFile("examples/hair_styles.png", rgba)
}

func exampleEyes() {
	nLowRight := lowRight
	nLowRight.X = nLowRight.X * len(colours.Eye)
	rgba := image.NewRGBA(image.Rectangle{Min: upLeft, Max: nLowRight})

	for i, e := range colours.Eye {
		img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

		drawAvatar(img, colours.Skin[0], colours.Black, e, hair.Styles[1])

		minP := image.Point{X: upLeft.X + (i)*dimentions.Width, Y: upLeft.Y}
		maxP := image.Point{X: lowRight.X + (i)*dimentions.Width, Y: lowRight.Y}
		rect := image.Rectangle{Min: minP, Max: maxP}

		draw.Draw(rgba, rect, img, image.Point{X: 0, Y: 0}, draw.Src)
	}

	draw2dimg.SaveToPngFile("examples/eyes.png", rgba)
}

func drawAvatar(img draw.Image, skinColour, hairColour, eyeColour color.Color, hairStyle hair.StyleHandler) {
	gc := draw2dimg.NewGraphicContext(img)

	drawFace(gc, skinColour)
	drawHair(gc, hairColour, hairStyle)
	drawEyes(gc, eyeColour)
	drawMouth(gc)
}

func drawFace(gc *draw2dimg.GraphicContext, c color.Color) {
	gc.SetFillColor(c)
	gc.SetStrokeColor(c)
	gc.SetLineWidth(1)

	draw2dkit.Circle(gc, float64(dimentions.Width/2), float64(dimentions.Height/2), float64(dimentions.Width)/100*75/2)
	gc.FillStroke()
}

func drawHair(gc *draw2dimg.GraphicContext, c color.Color, s hair.StyleHandler) {
	gc.SetFillColor(c)
	gc.SetStrokeColor(c)
	gc.SetLineWidth(1)

	s(gc)
}

func drawEyes(gc *draw2dimg.GraphicContext, c color.Color) {
	eyeSize := float64(dimentions.Width) / 100 * 11
	irisSize := eyeSize / 100 * 75
	pupilSize := irisSize / 100 * 45

	leftEyePos := [2]float64{float64(dimentions.Width/2 - dimentions.Width/7), float64(dimentions.Height/2 - dimentions.Height/8)}
	rightEyePos := [2]float64{float64(dimentions.Width/2 + dimentions.Width/7), float64(dimentions.Height/2 - dimentions.Height/8)}
	irisOffset := float64(dimentions.Width) / 100 * 1
	pupilOffset := float64(dimentions.Width) / 100 * 2

	// Draw whites
	gc.SetFillColor(colours.White)
	gc.SetStrokeColor(colours.White)
	gc.SetLineWidth(1)

	draw2dkit.Circle(gc, leftEyePos[0], leftEyePos[1], eyeSize)
	gc.FillStroke()
	draw2dkit.Circle(gc, rightEyePos[0], rightEyePos[1], eyeSize)
	gc.FillStroke()

	// Draw Irises
	gc.SetFillColor(c)
	gc.SetStrokeColor(c)

	draw2dkit.Circle(gc, leftEyePos[0]+irisOffset, leftEyePos[1]+irisOffset, irisSize)
	gc.FillStroke()
	draw2dkit.Circle(gc, rightEyePos[0]+irisOffset, rightEyePos[1]+irisOffset, irisSize)
	gc.FillStroke()

	// Draw Pupils
	gc.SetFillColor(colours.Black)
	gc.SetStrokeColor(colours.Black)

	draw2dkit.Circle(gc, leftEyePos[0]+pupilOffset, leftEyePos[1]+pupilOffset, pupilSize)
	gc.FillStroke()
	draw2dkit.Circle(gc, rightEyePos[0]+pupilOffset, rightEyePos[1]+pupilOffset, pupilSize)
	gc.FillStroke()
}

func drawMouth(gc *draw2dimg.GraphicContext) {
	gc.SetFillColor(colours.Black)
	gc.SetStrokeColor(colours.Black)
	gc.SetLineWidth(1)

	gc.ArcTo(128, 128+16+8, 64, 32+16, 180*(math.Pi/180), -180*(math.Pi/180))
	gc.Close()
	gc.FillStroke()

	gc.SetFillColor(colours.White)
	gc.SetStrokeColor(colours.White)

	draw2dkit.Rectangle(gc, 64+4, 128+16+8, 128+64-4, 128+16+8+12)
	gc.FillStroke()
}
