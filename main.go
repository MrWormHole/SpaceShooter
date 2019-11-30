package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const SCREEN_WIDTH, SCREEN_HEIGHT = 600, 800

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

	var enemies []enemy
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			x := float64(i)/5*SCREEN_WIDTH + 10
			y := float64(j)*100 + 10

			enemy := createEnemy(renderer, x, y)

			enemies = append(enemies, enemy)
		}
	}

	player := createPlayer(renderer)

	createProjectiles(renderer)

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		err = player.draw(renderer)
		if err != nil {
			fmt.Println("Player Drawing Error! ", err)
			quitAfterDelay()
			return
		}
		err = player.update()
		if err != nil {
			fmt.Println("Player Updating Error! ", err)
			quitAfterDelay()
			return
		}

		for _, enemy := range enemies {
			enemy.draw(renderer)
		}

		for _, projectile := range projectiles {
			if projectile != nil {
				projectile.draw(renderer)
				projectile.update()
			}
		}

		renderer.Present()
	}
}
