package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func textureFromPNG(renderer *sdl.Renderer, filename string) *sdl.Texture {
	src := sdl.RWFromFile(filename, "rb")
	image, err := img.LoadPNGRW(src)
	checkError("Image Loading Error! ", err)
	defer image.Free()

	texture, err := renderer.CreateTextureFromSurface(image)
	checkError("Texture Loading Error! ", err)
	return texture
}

func checkError(errorMessage string, err error) {
	if err != nil {
		fmt.Println(errorMessage, err)
		sdl.Delay(3000)
		sdl.Quit()
	}
}
