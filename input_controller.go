package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type inputController struct {
	attachedEntity *entity
	speed          float64

	sR *spriteRenderer
}

func createInputController(toAttached *entity, speed float64) *inputController {
	return &inputController{
		attachedEntity: toAttached,
		speed:          speed,
		sR:             toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (iC *inputController) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (iC *inputController) onUpdate() error {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_LEFT] == 1 && iC.attachedEntity.position.x > 0 {
		iC.attachedEntity.position.x -= iC.speed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 && iC.attachedEntity.position.x+float64(iC.sR.width) < SCREEN_WIDTH {
		iC.attachedEntity.position.x += iC.speed
	}

	return nil
}
