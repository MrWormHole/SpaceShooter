package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const projectilePixelWidth, projectilePixelHeight = 9, 57
const projectileSpeed = 0.5

type projectile struct {
	tex    *sdl.Texture
	x, y   float64
	angle  float64
	active bool
}

func createProjectile(renderer *sdl.Renderer) (p projectile) {
	p.tex = textureFromPNG(renderer, "projectile.png")

	return p
}

func (p *projectile) draw(renderer *sdl.Renderer) {
	if !p.active {
		return
	}

	renderer.Copy(p.tex,
		&sdl.Rect{X: 0, Y: 0, W: projectilePixelWidth, H: projectilePixelHeight},
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: projectilePixelWidth, H: projectilePixelHeight})
}

func (p *projectile) update() {
	if !p.active {
		return
	}
	/*
		Note for those who don't understand basic trigonometry
		  when angle is 0, it will go right
		  when angle is 90, it will go downward
		  when angle is 180, it will go left
		  when angle is 270, it will go upwards
	*/

	p.x += projectileSpeed * math.Cos(p.angle)
	p.y += projectileSpeed * math.Sin(p.angle)

	if p.x > SCREEN_WIDTH || p.x < 0 || p.y > SCREEN_HEIGHT+projectilePixelHeight || p.y < 0-projectilePixelHeight {
		p.active = false
	}
}

var projectiles []*projectile

func createProjectiles(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		projectile := createProjectile(renderer)
		projectiles = append(projectiles, &projectile)
	}
}

func projectileFromProjectiles() (*projectile, bool) {
	// THERE MIGHT BE BUG HERE PLEASE TEST
	for _, p := range projectiles {
		if !p.active {
			return p, true
		}
	}

	return nil, false
}
