package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const enemyPixelWidth, enemyPixelHeight = 103, 84
const enemySpeed = 0.35

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func createEnemy(renderer *sdl.Renderer, x float64, y float64) (e enemy) {
	e.tex = textureFromPNG(renderer, "enemy.png")
	e.x, e.y = x, y

	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	renderer.CopyEx(e.tex,
		&sdl.Rect{X: 0, Y: 0, W: enemyPixelWidth, H: enemyPixelHeight},
		&sdl.Rect{X: int32(e.x), Y: int32(e.y), W: enemyPixelWidth, H: enemyPixelHeight},
		0,
		&sdl.Point{X: 51, Y: 42},
		sdl.FLIP_NONE)
}
