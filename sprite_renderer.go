package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct {
	attachedEntity *entity
	texture        *sdl.Texture

	width, height int32
}

func createSpriteRenderer(toAttached *entity, renderer *sdl.Renderer, filename string) *spriteRenderer {
	texture := textureFromPNG(renderer, filename)
	_, _, width, height, err := texture.Query()
	checkError("Renderer Query Error! ", err)

	return &spriteRenderer{
		attachedEntity: toAttached,
		texture:        texture,
		width:          width,
		height:         height}
}

func (spriteRenderer *spriteRenderer) onDraw(renderer *sdl.Renderer) error {

	renderer.CopyEx(spriteRenderer.texture,
		&sdl.Rect{X: 0, Y: 0, W: spriteRenderer.width, H: spriteRenderer.height},
		&sdl.Rect{X: int32(spriteRenderer.attachedEntity.position.x),
			Y: int32(spriteRenderer.attachedEntity.position.y),
			W: spriteRenderer.width,
			H: spriteRenderer.height},
		0,
		&sdl.Point{X: int32(spriteRenderer.width / 2.0), Y: int32(spriteRenderer.height / 2.0)},
		sdl.FLIP_NONE)

	return nil
}

func (spriteRenderer *spriteRenderer) onUpdate() error {
	return nil
}
