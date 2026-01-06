package shared

import (
	"image"
	_ "image/jpeg"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadAssets() (*ebiten.Image, *ebiten.Image, *ebiten.Image, *ebiten.Image) {
	grassTexture := loadGrassTexture()
	antImage := loadAntAsset()
	homeImage := loadHomeAsset()
	foodSourceImage := loadFoodSourceAsset()

	return grassTexture, antImage, homeImage, foodSourceImage
}

func loadGrassTexture() *ebiten.Image {
	originalGrassTexture := loadImage("../assets/grass_texture.jpeg")
	return rescaleImage(originalGrassTexture, 0.35)
}

func loadAntAsset() *ebiten.Image {
	originalAntImage := loadImage("../assets/ant_sprite.png")
	return rescaleImage(originalAntImage, 0.05)
}

func loadHomeAsset() *ebiten.Image {
	originaHomeImage := loadImage("../assets/colony_home.png")
	return rescaleImage(originaHomeImage, 0.1)
}

func loadFoodSourceAsset() *ebiten.Image {
	originalFoodSourceImage := loadImage("../assets/food_source.png")
	return rescaleImage(originalFoodSourceImage, 0.25)
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
