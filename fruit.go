package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Fruit struct
type Fruit struct {
	X, Y int32
}

// Draw fruit
func (fruit *Fruit) Draw(r *sdl.Renderer) error {
	fruitImage, err := img.LoadTexture(r, "resources/images/fruit.png")

	if err != nil {
		return fmt.Errorf("Could not load fruit: %v", err)
	}

	rect := &sdl.Rect{X: fruit.X * TILE, Y: fruit.Y * TILE, H: TILE - 1, W: TILE - 1}

	err = r.Copy(fruitImage, nil, rect)

	if err != nil {
		return fmt.Errorf("could not copy fruit: %v", err)
	}

	return nil
}
