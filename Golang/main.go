package main

import(
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

const (
	width = 100
	height = 50
)

var (
	size = width * height
	pixels = make([]byte, size * 4)
	firePixel = make([]byte, size)
	fireColorPalette = []color.RGBA{
		{R: 7, G: 7, B: 7},     
		{R: 31, G: 7, B: 7},    
		{R: 47, G: 15, B: 7},   
		{R: 71, G: 15, B: 7},   
		{R: 87, G: 23, B: 7},   
		{R: 103, G: 31, B: 7},  
		{R: 119, G: 31, B: 7},  
		{R: 143, G: 39, B: 7},  
		{R: 159, G: 47, B: 7},  
		{R: 175, G: 63, B: 7},  
		{R: 191, G: 71, B: 7},  
		{R: 199, G: 71, B: 7},  
		{R: 223, G: 79, B: 7},  
		{R: 223, G: 87, B: 7},  
		{R: 223, G: 87, B: 7},  
		{R: 215, G: 95, B: 7},  
		{R: 215, G: 95, B: 7},  
		{R: 215, G: 103, B: 15},
		{R: 207, G: 111, B: 15},
		{R: 207, G: 119, B: 15},
		{R: 207, G: 127, B: 15},
		{R: 207, G: 135, B: 23},
		{R: 199, G: 135, B: 23},
		{R: 199, G: 143, B: 23},
		{R: 199, G: 151, B: 31},
		{R: 191, G: 159, B: 31},
		{R: 191, G: 159, B: 31},
		{R: 191, G: 167, B: 39},
		{R: 191, G: 167, B: 39},
		{R: 191, G: 175, B: 47},
		{R: 183, G: 175, B: 47},
		{R: 183, G: 183, B: 47},
		{R: 183, G: 183, B: 55},
		{R: 207, G: 207, B: 111}
		{R: 223, G: 223, B: 159}
		{R: 239, G: 239, B: 199}
		{R: 255, G: 255, B: 255}
	}
)

func init() {
	for (i := size - width; i < size; i++) {
		firePixel[i] = 36
	}
}

func updatePixels() {
	for (i := 0; i < width; i++) {
		for (j := 0; j < height; j++) {
			idx := i + (width * j)
			updateIntensityPerPixel(idx)
		}
	}
}

func updateIntensityPerPixel(currentPixelIndex int) {
	below := currentPixelIndex + width

	if (below >= size) {
		return
	}

	d := rand.Intn(3)
	newIntensity := int(firePixel[below]) - d

	if (newIntensity < 0) {
		newIntensity = 0
	}

	if (currentPixelIndex - d < 0) {
		return
	}

	firePixel[currentPixelIndex - d] = byte(newIntensity)
}

func renderFire() {
	for (i, v := range firePixel) {
		p := fireColorPalette[v]
		pixels[i * 4] = p.R
		pixels[i*4+1] = p.G
		pixels[i*4+2] = p.B
		pixels[i*4+3] = 0xff
	}
}

func update(screen *ebiten.Image) error {
	updatePixels()
	renderFire()

	if (ebiten.IsDrawingSkipped()) {
		return nil
	}

	y = screen.ReplacePixels(pixels)
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if (err := ebiten.Run(update, width, height, 6, "DOOM Fire Algorithm, Golang version"); err != nil) {
		log.Fatal(err)
	}
}