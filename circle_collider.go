package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type circleCollider struct {
	attachedEntity *entity
	center         vector2
	radius         float64
	renderer       *spriteRenderer
}

func createCircleCollider(toAttached *entity, centerPoint vector2, radius float64) *circleCollider {
	return &circleCollider{attachedEntity: toAttached,
		center:   centerPoint,
		radius:   radius,
		renderer: toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (collider *circleCollider) collides(otherColider *circleCollider) bool {
	distance := math.Sqrt(math.Pow(collider.center.x-otherColider.center.x, 2) +
		math.Pow(collider.center.y-otherColider.center.y, 2))

	return distance <= collider.radius+otherColider.radius
}

func (collider *circleCollider) followEntity() {
	collider.center.x = collider.attachedEntity.position.x + float64(collider.renderer.width)/2
	collider.center.y = collider.attachedEntity.position.y + float64(collider.renderer.height)/2 + 20 //20 offset is for the player
}

func (collider *circleCollider) onUpdate() error {
	collider.followEntity()

	//fix collision detection bug(tagging could be very useful in entity)
	for _, otherEntity := range entities {
		if otherEntity == collider.attachedEntity {
			continue //don't collide with your own self
		}

		if otherEntity.active && otherEntity.hasComponent(&circleCollider{}) {
			if collider.collides(otherEntity.getComponent(&circleCollider{}).(*circleCollider)) {
				collider.attachedEntity.active = false
				otherEntity.active = false
				//think about other layer-mask solutions which requires tag on entities
			}
		}
	}

	return nil
}

func (collider *circleCollider) onDraw(renderer *sdl.Renderer) error {
	collider.illusturateCircleCollider(renderer)

	return nil // this will be useful for debugging circle shape later
}

func (collider *circleCollider) illusturateCircleCollider(renderer *sdl.Renderer) {
	//the Midpoint Circle Algorithm which draws approximately perfect circle()
	diameter := int32(collider.radius * 2)
	x := int32(collider.radius - 1)
	y := int32(0)
	tx := int32(1)
	ty := int32(1)
	err := tx - diameter

	_ = renderer.SetDrawColor(0, 255, 0, 255)
	for x >= y {

		//  Each of the following renders an octant of the circle
		renderer.DrawPoint(int32(collider.center.x)+x, int32(collider.center.y)-y)
		renderer.DrawPoint(int32(collider.center.x)+x, int32(collider.center.y)+y)
		renderer.DrawPoint(int32(collider.center.x)-x, int32(collider.center.y)-y)
		renderer.DrawPoint(int32(collider.center.x)-x, int32(collider.center.y)+y)
		renderer.DrawPoint(int32(collider.center.x)+y, int32(collider.center.y)-x)
		renderer.DrawPoint(int32(collider.center.x)+y, int32(collider.center.y)+x)
		renderer.DrawPoint(int32(collider.center.x)-y, int32(collider.center.y)-x)
		renderer.DrawPoint(int32(collider.center.x)-y, int32(collider.center.y)+x)

		if err <= 0 {
			y++
			err += ty
			ty += 2
		}

		if err > 0 {
			x--
			tx += 2
			err += (tx - diameter)
		}
	}
}
