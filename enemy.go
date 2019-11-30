package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const enemyPixelWidth, enemyPixelHeight = 103, 84
const enemySpeed = 0.35

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func createEnemy(renderer *sdl.Renderer, position vector2) *entity {
	enemy := createEntity()

	enemy.position = position
	enemy.active = true

	spriteRendererComponent := createSpriteRenderer(enemy, renderer, "enemy.png")
	enemy.addComponent(spriteRendererComponent)

	return enemy
}
