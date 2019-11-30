package main

import (
	"math"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type timeTrigger struct {
	attachedEntity *entity
	cooldown       time.Duration
	lastTriggered  time.Time

	sR *spriteRenderer
}

func createTimeTrigger(toAttached *entity, cooldown time.Duration) *timeTrigger {
	return &timeTrigger{
		attachedEntity: toAttached,
		cooldown:       cooldown,
		sR:             toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (trigger *timeTrigger) onUpdate() error {
	keys := sdl.GetKeyboardState()

	spawnPosition := trigger.attachedEntity.position

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(trigger.lastTriggered) >= trigger.cooldown {
			trigger.action(spawnPosition.x+float64(trigger.sR.width/2-4), spawnPosition.y)
		}
	}

	return nil
}

func (trigger *timeTrigger) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (trigger *timeTrigger) action(x float64, y float64) {
	proj, status := projectileFromProjectiles()
	if status {
		proj.active = true
		proj.x = x
		proj.y = y - 75
		proj.angle = 3 * math.Pi / 2

		trigger.lastTriggered = time.Now()
	}
}
