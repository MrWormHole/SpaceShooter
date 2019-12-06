package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type timeTrigger struct {
	attachedEntity *entity
	cooldown       time.Duration
	lastTriggered  time.Time

	renderer *spriteRenderer
}

func createTimeTrigger(toAttached *entity, cooldown time.Duration) *timeTrigger {
	return &timeTrigger{
		attachedEntity: toAttached,
		cooldown:       cooldown,
		renderer:       toAttached.getComponent(&spriteRenderer{}).(*spriteRenderer)}
}

func (trigger *timeTrigger) onUpdate() error {

	spawnPosition := trigger.attachedEntity.position

	if trigger.attachedEntity.tag == "player" && trigger.attachedEntity.active {
		keys := sdl.GetKeyboardState()

		if keys[sdl.SCANCODE_SPACE] == 1 {
			if time.Since(trigger.lastTriggered) >= trigger.cooldown {
				trigger.action(spawnPosition.x+float64(trigger.renderer.width/2-4), spawnPosition.y-50, 3*math.Pi/2)
				trigger.lastTriggered = time.Now()
			}
		}
	}

	if trigger.attachedEntity.tag == "enemy" && trigger.attachedEntity.active {
		if time.Since(trigger.lastTriggered) >= trigger.cooldown {
			chance := rand.Int() % 10 // 20% chance
			if chance < 2 {
				trigger.action(spawnPosition.x+float64(trigger.renderer.width/2-4), spawnPosition.y+50, math.Pi/2)
			}
			trigger.lastTriggered = time.Now()
		}
	}

	return nil
}

func (trigger *timeTrigger) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (trigger *timeTrigger) action(x float64, y float64, rotation float64) {
	proj, status := projectileFromProjectilesPool()
	if status {
		proj.active = true
		if rotation == math.Pi/2 {
			proj.tag = "enemyProjectile"
		} else if rotation == 3*math.Pi/2 {
			proj.tag = "playerProjectile"
		}
		proj.position.x = x
		proj.position.y = y
		proj.rotation = rotation
	}
}
