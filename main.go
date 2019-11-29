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

	var enemies []enemy
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*WIDTH + 10
			y := float64(j)*100 + 10

			enemy := createEnemy(renderer, x, y)

			enemies = append(enemies, enemy)
		}
	}

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

		player.update()
		player.draw(renderer)

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		renderer.Present()
	}
}
