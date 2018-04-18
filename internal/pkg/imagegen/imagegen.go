package imagegen

import (
	"image"
	"image/color"
	"image/draw"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// GenerateUUIDImage generates an image with the given uuid written onto it.
func GenerateUUIDImage(uuid uuid.UUID) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 260, 50))
	blue := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	addLabel(img, 5, 25, uuid.String())
	return img
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 255, 255, 255}

	point := fixed.Point26_6{
		X: fixed.Int26_6(x * 64),
		Y: fixed.Int26_6(y * 64),
	}

	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
}
