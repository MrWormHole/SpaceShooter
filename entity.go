package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type vector2 struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
}

type entity struct {
	position   vector2
	rotation   float64
	active     bool
	components []component
}

func createEntity() *entity {
	return &entity{}
}

func (e *entity) draw(renderer *sdl.Renderer) error {
	for _, component := range e.components {
		err := component.onDraw(renderer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *entity) update() error {
	for _, component := range e.components {
		err := component.onUpdate()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *entity) addComponent(new component) {
	for _, existingComponent := range e.components {
		if reflect.TypeOf(new) == reflect.TypeOf(existingComponent) {
			panic(fmt.Sprintf(
				`What the heck were you thinking!?! 
				The component that you are trying to add already exists! 
				Take this out of your component please: %v`,
				reflect.TypeOf(new)))
		}
	}
	e.components = append(e.components, new)
}

func (e *entity) getComponent(withType component) component {
	desiredType := reflect.TypeOf(withType)
	for _, existingComponent := range e.components {
		if reflect.TypeOf(existingComponent) == desiredType {
			return existingComponent
		}
	}
	panic(fmt.Sprintf(
		`What the heck were you thinking!?!
		The component that you are trying to get doesn't even exist
		Search for something except this type %v
		`,
		reflect.TypeOf(withType)))
}
