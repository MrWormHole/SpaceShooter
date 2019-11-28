package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const WIDTH, HEIGHT = 600, 800

func quitAfterDelay() {
	sdl.Delay(3000)
	sdl.Quit()
}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("SDL Initialization Error! ", err)
		quitAfterDelay()
	}

	window, err := sdl.CreateWindow("Space Shooter",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WIDTH,
		HEIGHT,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Window Initialization Error! ", err)
		quitAfterDelay()
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window,
		-1,
		sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Renderer Initialization Error! ", err)
		quitAfterDelay()
	}
	defer renderer.Destroy()

	player := createPlayer(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		player.draw(renderer)

		renderer.Present()
	}
}
