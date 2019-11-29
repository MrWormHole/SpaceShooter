package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const playerPixelWidth, playerPixelHeight = 112, 75
const playerSpeed, playerShootCooldown = 0.25, time.Millisecond * 250

type player struct {
	tex  *sdl.Texture
	x, y float64

	lastShootTime time.Time
}

func createPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = textureFromPNG(renderer, "player.png")
	p.x, p.y = SCREEN_WIDTH/2.0-playerPixelWidth/2, SCREEN_HEIGHT-playerPixelHeight-25

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: playerPixelWidth, H: playerPixelHeight},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: playerPixelWidth, H: playerPixelHeight})
}

func (p *player) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && p.x > 0 {
		p.x -= playerSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && p.x+playerPixelWidth < SCREEN_WIDTH {
		p.x += playerSpeed
	}

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(p.lastShootTime) >= playerShootCooldown {
			p.shoot()
		}
	}
}

func (p *player) shoot() {
	proj, status := projectileFromProjectiles()
	if status {
		proj.active = true
		proj.x = p.x + playerPixelWidth/2 - projectilePixelWidth/2
		proj.y = p.y - 75
		proj.angle = math.Pi * 3 / 2

		p.lastShootTime = time.Now()
	}
}
