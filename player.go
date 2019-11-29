package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const playerSpeed = 0.25

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func createPlayer(renderer *sdl.Renderer) (p player) {

	src := sdl.RWFromFile("player.png", "rb")
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
	p.tex, p.x, p.y = texture, WIDTH/2.0-56, HEIGHT-75-25

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 112, H: 75},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: 112, H: 75})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += playerSpeed
	}
}
