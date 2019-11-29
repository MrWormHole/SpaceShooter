package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const enemySpeed = 0.35

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func createEnemy(renderer *sdl.Renderer, x float64, y float64) (e enemy) {

	src := sdl.RWFromFile("enemy.png", "rb")
	image, err := img.LoadPNGRW(src)
	if err != nil {
		fmt.Println("Image Loading Error! ", err)
		quitAfterDelay()
	}
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Println("Texture Loading Error! ", err)
		quitAfterDelay()
	}
	e.tex, e.x, e.y = texture, x, y

	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	renderer.CopyEx(e.tex,
		&sdl.Rect{X: 0, Y: 0, W: 103, H: 84},
		&sdl.Rect{X: int32(e.x), Y: int32(e.y), W: 103, H: 84},
		0,
		&sdl.Point{X: 51, Y: 42},
		sdl.FLIP_NONE)
}
