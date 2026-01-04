package shared

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadAssets() (*ebiten.Image, *ebiten.Image) {
	antImage := loadAntAsset()
	homeImage := loadHomeAsset()

	return antImage, homeImage
}

func loadAntAsset() *ebiten.Image {
	originalAntImage := loadImage("../assets/ant_sprite.png")
	return rescaleImage(originalAntImage, 0.05)
}

func loadHomeAsset() *ebiten.Image {
	originalAntImage := loadImage("../assets/colony_home.png")
	return rescaleImage(originalAntImage, 0.1)
}

func loadImage(path string) *ebiten.Image {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func rescaleImage(image *ebiten.Image, scale float64) *ebiten.Image {
	w, h := image.Bounds().Dx(), image.Bounds().Dy()

	smallAntImage := ebiten.NewImage(int(float64(w)*scale), int(float64(h)*scale))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	smallAntImage.DrawImage(image, op)

	return smallAntImage
}
