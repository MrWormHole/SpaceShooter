package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const playerPixelWidth, playerPixelHeight = 112, 75
const playerSpeed, playerShootCooldown = 0.25, time.Millisecond * 250

func createPlayer(renderer *sdl.Renderer) *entity {
	player := createEntity()

	player.position = vector2{
		x: SCREEN_WIDTH/2.0 - playerPixelWidth/2,
		y: SCREEN_HEIGHT - playerPixelHeight - 25}

	player.active = true

	spriteRendererComponent := createSpriteRenderer(player, renderer, "player.png")
	player.addComponent(spriteRendererComponent)

	inputControllerComponent := createInputController(player, playerSpeed)
	player.addComponent(inputControllerComponent)

	timeTriggerComponent := createTimeTrigger(player, playerShootCooldown)
	player.addComponent(timeTriggerComponent)

	circleColliderComponent := createCircleCollider(player, player.position, 50)
	player.addComponent(circleColliderComponent)

	return player
}
