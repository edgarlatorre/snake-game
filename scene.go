package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg    *sdl.Texture
	snake *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "resources/images/Background.png")

	if err != nil {
		return nil, fmt.Errorf("Could not load background: %v", err)
	}

	return &scene{bg: bg}, nil
}

func (s *scene) draw(r *sdl.Renderer, snake *Snake) error {
	r.Clear()

	err := r.Copy(s.bg, nil, nil)

	if err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	err = snake.Draw(r)

	if err != nil {
		return fmt.Errorf("could not draw snake: %v", err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
