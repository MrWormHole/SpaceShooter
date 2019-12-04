package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type inputController struct {
	attachedEntity *entity
	speed          float64

	renderer *spriteRenderer
}

func createInputController(toAttached *entity, speed float64) *inputController {
	return &inputController{
		attachedEntity: toAttached,
		speed:          speed,
		renderer:       toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (controller *inputController) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (controller *inputController) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && controller.attachedEntity.position.x > 0 {
		controller.attachedEntity.position.x -= controller.speed * delta
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && controller.attachedEntity.position.x+float64(controller.renderer.width) < SCREEN_WIDTH {
		controller.attachedEntity.position.x += controller.speed * delta
	}

	return nil
}
