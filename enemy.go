package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const enemyPixelWidth, enemyPixelHeight = 103, 84
const enemySpeed = 3

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func createEnemy(renderer *sdl.Renderer, position vector2) *entity {
	enemy := createEntity("enemy")

	enemy.position = position
	enemy.active = true

	spriteRendererComponent := createSpriteRenderer(enemy, renderer, "enemy.png")
	enemy.addComponent(spriteRendererComponent)

	circleColliderComponent := createCircleCollider(enemy, enemy.position, 30)
	enemy.addComponent(circleColliderComponent)

	return enemy
}
