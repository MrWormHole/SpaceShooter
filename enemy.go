package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const enemyPixelWidth, enemyPixelHeight = 103, 84
const enemySpeed, enemyRoamCooldown = 3, time.Millisecond * 750

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

	aiControllerComponent := createAIController(enemy, enemySpeed, enemyRoamCooldown)
	enemy.addComponent(aiControllerComponent)

	return enemy
}
