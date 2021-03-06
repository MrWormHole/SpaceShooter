package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const projectilePixelWidth, projectilePixelHeight = 9, 57
const projectileSpeed = 10

func createProjectile(renderer *sdl.Renderer) *entity {
	projectile := createEntity("projectile")

	spriteRendererComponent := createSpriteRenderer(projectile, renderer, "projectile.png")
	projectile.addComponent(spriteRendererComponent)

	projectileControllerComponent := createProjectileController(projectile, projectileSpeed)
	projectile.addComponent(projectileControllerComponent)

	projectile.active = false

	circleColliderComponent := createCircleCollider(projectile, projectile.position, 10)
	projectile.addComponent(circleColliderComponent)

	return projectile
}

var projectilesPool []*entity

func createProjectiles(renderer *sdl.Renderer) {
	for i := 0; i < 30; i++ {
		projectile := createProjectile(renderer)
		entities = append(entities, projectile)
		projectilesPool = append(projectilesPool, projectile)
	}
}

func projectileFromProjectilesPool() (*entity, bool) {
	for _, projectile := range projectilesPool {
		if !projectile.active {
			return projectile, true
		}
	}

	return nil, false
}
