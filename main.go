package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 600, 800
const TARGET_FPS = 60

func quitAfterDelay() {
	sdl.Delay(3000)
	sdl.Quit()
}

var delta float64

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("SDL Initialization Error! ", err)
		quitAfterDelay()
	}

	window, err := sdl.CreateWindow("Space Shooter",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH,
		SCREEN_HEIGHT,
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

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*SCREEN_WIDTH + 10
			y := float64(j)*100 + 10

			enemy := createEnemy(renderer, vector2{x, y})

			entities = append(entities, enemy)
		}
	}

	entities = append(entities, createPlayer(renderer))

	createProjectiles(renderer)

	for {
		frameStartTime := time.Now()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		for _, entity := range entities {
			if entity.active {
				err = entity.draw(renderer)
				if err != nil {
					fmt.Println("Entity Drawing Error! ", err)
					quitAfterDelay()
					return
				}
				err = entity.update()
				if err != nil {
					fmt.Println("Entity Updating Error! ", err)
					quitAfterDelay()
					return
				}
			}
		}
		delta = time.Since(frameStartTime).Seconds() * TARGET_FPS
		renderer.Present()
	}
}
