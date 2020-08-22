package hair

import (
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	"github.com/status-im/avatars/dimentions"
)

type StyleHandler func(gc *draw2dimg.GraphicContext)

func Bald(gc *draw2dimg.GraphicContext) {
	return
}

func Mohawk(gc *draw2dimg.GraphicContext){
	hairPos := [4]float64{
		float64(dimentions.Width)/2 - float64(dimentions.Width)/100*10,
		float64(dimentions.Height) / 100 * 5,
		float64(dimentions.Width)/2 + float64(dimentions.Width)/100*10,
		float64(dimentions.Height) / 100 * 20,
	}

	draw2dkit.Rectangle(gc, hairPos[0], hairPos[1], hairPos[2], hairPos[3])
	gc.FillStroke()
}

var (
	Styles = []StyleHandler{
		Bald,
		Mohawk,
	}
)


