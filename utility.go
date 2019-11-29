package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func textureFromPNG(renderer *sdl.Renderer, filename string) *sdl.Texture {
	src := sdl.RWFromFile(filename, "rb")
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
	return texture
}
