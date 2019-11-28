package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func createPlayer(renderer *sdl.Renderer) (p player) {

	src := sdl.RWFromFile("test.png", "rb")
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
	p.tex = texture

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: 112, H: 75},
		&sdl.Rect{X: 0, Y: 0, W: 112, H: 75},
	)
}
