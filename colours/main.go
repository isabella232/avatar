package colours

import (
	"image/color"
	"math/rand"
	"time"
)

var (
	// #000000 black
	Black = color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	// #ffffff white
	White = color.RGBA{R: 255, G: 255, B: 255, A: 0xff}

	// SKIN COLOURS
	//#FFC83D
	SkinEmojiYellow = color.RGBA{R: 255, G: 200, B: 61, A: 0xff}
	//#F7D7C4
	SkinLight = color.RGBA{R: 247, G: 215, B: 196, A: 0xff}
	//#D8B094
	SkinMediumLight = color.RGBA{R: 216, G: 176, B: 148, A: 0xff}
	//#BB9167
	SkinMedium = color.RGBA{R: 187, G: 145, B: 103, A: 0xff}
	//#8E562E
	SkinMediumDark = color.RGBA{R: 142, G: 86, B: 46, A: 0xff}
	//#613D30
	SkinDark = color.RGBA{R: 97, G: 61, B: 48, A: 0xff}

	Skin = []color.Color{
		SkinEmojiYellow,
		SkinLight,
		SkinMediumLight,
		SkinMedium,
		SkinMediumDark,
		SkinDark,
	}

	// HAIR COLOURS
	// NATURAL
	// #14191C jet black
	// #2E2D29 coffee brown
	// #35221B chestnut brown
	// #553B2A cappuccino
	// #6E523D golden brown
	// #9E704E mousey brown
	// #A56036 warm light brown

	// SIMPLE
	// #000116 black
	hairBlack = color.RGBA{R: 0, G: 1, B: 22, A: 0xff}
	// #582700 dark brown
	hairDarkBrown = color.RGBA{R: 88, G: 39, B: 0, A: 0xff}
	// #DF9A5F light brown
	hairLightBrown = color.RGBA{R: 223, G: 154, B: 95, A: 0xff}
	// #F25E1A auburn
	hairAuburn = color.RGBA{R: 242, G: 94, B: 26, A: 0xff}
	// #AB6B25 ginger
	hairGinger = color.RGBA{R: 171, G: 107, B: 37, A: 0xff}
	// #FCCB83 strawberry blonde
	hairStrawberryBlonde = color.RGBA{R: 252, G: 203, B: 131, A: 0xff}
	// #FFECBA blonde
	hairBlonde = color.RGBA{R: 255, G: 236, B: 186, A: 0xff}
	// #E0DBC6 platinum blonde
	hairPlatinumBlonde = color.RGBA{R: 224, G: 219, B: 198, A: 0xff}
	// #E3E0D6 silver
	hairSilver = color.RGBA{R: 227, G: 224, B: 214, A: 0xff}
	// #EEEFEB white
	hairWhite = color.RGBA{R: 238, G: 239, B: 235, A: 0xff}

	// FUN
	// #DF2624 red
	hairRed = color.RGBA{R: 223, G: 38, B: 36, A: 0xff}
	// #D619D7 fuscia
	hairFucia = color.RGBA{R: 214, G: 25, B: 215, A: 0xff}
	// #651ED1 purple
	hairPurple = color.RGBA{R: 101, G: 30, B: 209, A: 0xff}
	// #263ECF blue
	hairBlue = color.RGBA{R: 38, G: 62, B: 207, A: 0xff}
	// #248FC2 teal
	hairTeal = color.RGBA{R: 36, G: 143, B: 194, A: 0xff}
	// #1DA97E marine
	hairMarine = color.RGBA{R: 29, G: 169, B: 126, A: 0xff}
	// #45B122 green
	hairGreen = color.RGBA{R: 69, G: 177, B: 34, A: 0xff}
	// #96AD21 lime
	hairLime = color.RGBA{R: 150, G: 173, B: 33, A: 0xff}
	// #DBC43A yellow
	hairYellow = color.RGBA{R: 219, G: 196, B: 58, A: 0xff}
	// #D17724 orange
	hairOrange = color.RGBA{R: 209, G: 119, B: 36, A: 0xff}

	Hair = []color.Color{
		hairBlack,
		hairDarkBrown,
		hairLightBrown,
		hairAuburn,
		hairGinger,
		hairStrawberryBlonde,
		hairBlonde,
		hairPlatinumBlonde,
		hairSilver,
		hairWhite,

		hairRed,
		hairFucia,
		hairPurple,
		hairBlue,
		hairTeal,
		hairMarine,
		hairGreen,
		hairLime,
		hairYellow,
		hairOrange,
	}

	// EYE COLOUR
	// #874098 amethyst
	eyeAmethyst = color.RGBA{R: 135, G: 64, B: 152, A: 0xff}
	// #587BA2 blue
	eyeBlue = color.RGBA{R: 88, G: 123, B: 162, A: 0xff}
	// #5F937F green
	eyeGreen = color.RGBA{R: 95, G: 147, B: 127, A: 0xff}
	// #5E4E40 brown
	eyeBrown = color.RGBA{R: 94, G: 78, B: 64, A: 0xff}

	Eye = []color.Color{
		eyeAmethyst,
		eyeBlue,
		eyeGreen,
		eyeBrown,
		Black,
	}
)

func Rand(cs []color.Color) color.Color {
	rand.Seed(time.Now().UnixNano())
	return cs[rand.Intn(len(cs))]
}
