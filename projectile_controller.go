package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type projectileController struct {
	attachedEntity *entity
	speed          float64
}

func createProjectileController(toAttached *entity, speed float64) *projectileController {
	return &projectileController{
		attachedEntity: toAttached,
		speed:          speed}
}

func (controller *projectileController) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (controller *projectileController) onUpdate() error {
	/*
		Note for those who don't understand basic trigonometry
		  when angle is 0, it will go right
		  when angle is 90, it will go downward
		  when angle is 180, it will go left
		  when angle is 270, it will go upwards
	*/
	entity := controller.attachedEntity

	entity.position.x += projectileSpeed * math.Cos(entity.rotation) * delta
	entity.position.y += projectileSpeed * math.Sin(entity.rotation) * delta

	if entity.position.x > SCREEN_WIDTH || entity.position.x < 0 ||
		entity.position.y > SCREEN_HEIGHT+projectilePixelHeight || entity.position.y < 0-projectilePixelHeight {
		entity.active = false
	}

	return nil
}
