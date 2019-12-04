package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type aiController struct {
	attachedEntity  *entity
	speed           float64
	cooldownRoaming time.Duration
	lastDirected    time.Time
	direction       float64
	runOnce         bool
	renderer        *spriteRenderer
}

func createAIController(toAttached *entity, speed float64, cooldownWalk time.Duration) *aiController {
	return &aiController{
		attachedEntity:  toAttached,
		speed:           speed,
		direction:       -1,
		cooldownRoaming: cooldownWalk,
		runOnce:         true,
		renderer:        toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (controller *aiController) onUpdate() error {
	controller.attachedEntity.position.x += controller.speed * delta * controller.direction

	if time.Since(controller.lastDirected) > controller.cooldownRoaming {
		controller.direction = -controller.direction
		controller.lastDirected = time.Now()

		// i hate myself for writing this but it just works and can not produce better design pattern
		if controller.runOnce && controller.direction == -1 {
			controller.cooldownRoaming *= 2
			controller.runOnce = false
		}
	}

	return nil
}

func (controller *aiController) onDraw(renderer *sdl.Renderer) error {
	return nil
}
