package world

import (
	"github.com/Tnze/go-mc/bot/maths"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestWorld_PathFind(t *testing.T) {
	tS := time.Now()
	w := NewWorld()
	point1, point2 := maths.Vec3d{X: 3644, Y: 4, Z: 354}, maths.Vec3d{X: 2205, Y: 255, Z: 3378}
	path := w.PathFind(point1, point2)
	DrawPath(path)
	t.Log("Path found in", time.Since(tS))
}

func DrawPath(path []maths.Vec3d) {
	const (
		gridSize  = 4096
		pointSize = 8
	)
	// Create an image
	img := image.NewRGBA(image.Rect(0, 0, gridSize, gridSize))

	// Draw the grid on the image
	for x := 0; x < gridSize; x += 16 {
		for y := 0; y < gridSize; y += 16 {
			for i := 0; i < 16; i++ {
				img.Set(x+i, y, color.RGBA{0, 0, 0, 255})
				img.Set(x, y+i, color.RGBA{0, 0, 0, 255})
			}
		}
	}

	// Draw the path on the image.
	// The color of the y position is from red to green depending on the height.
	for _, point := range path {
		for x := 0; x < pointSize; x++ {
			for y := 0; y < pointSize; y++ {
				img.Set(int(point.X)+x, int(point.Z)+y, color.RGBA{R: uint8(point.Y), G: 255 - uint8(point.Y), A: 255})
			}
		}
	}

	// Save the image to a file
	f, _ := os.Create("path.png")
	defer f.Close()
	png.Encode(f, img)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
