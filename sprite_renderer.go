package main

import (
	"fmt"

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
	if err != nil {
		fmt.Println("Renderer Query Error! ", err)
		quitAfterDelay()
	}

	return &spriteRenderer{
		attachedEntity: toAttached,
		texture:        texture,
		width:          width,
		height:         height}
}

func (sR *spriteRenderer) onDraw(renderer *sdl.Renderer) error {

	renderer.CopyEx(sR.texture,
		&sdl.Rect{X: 0, Y: 0, W: sR.width, H: sR.height},
		&sdl.Rect{X: int32(sR.attachedEntity.position.x), Y: int32(sR.attachedEntity.position.y), W: sR.width, H: sR.height},
		sR.attachedEntity.rotation,
		&sdl.Point{X: int32(sR.width / 2.0), Y: int32(sR.height / 2.0)},
		sdl.FLIP_NONE)

	return nil
}

func (sR *spriteRenderer) onUpdate() error {
	return nil
}
