package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Snake struct
type Snake struct {
	X, Y, Velx, Vely int32
	Pos              [][]int32
	Tail             int
	action           string
}

// Update snake
func (snake *Snake) Update() {
	snake.X += snake.Velx
	snake.Y += snake.Vely

	if snake.X > WIDTH_IN_TILE {
		snake.X = 0
	}

	if snake.X < 0 {
		snake.X = WIDTH_IN_TILE
	}

	if snake.Y > WIDTH_IN_TILE {
		snake.Y = 0
	}

	if snake.Y < 0 {
		snake.Y = WIDTH_IN_TILE
	}

	snake.Pos = append(snake.Pos, []int32{snake.X, snake.Y})

	for len(snake.Pos) > snake.Tail {
		fmt.Println("Slice", snake.Pos[1:])
		snake.Pos = snake.Pos[1:]

	}
}

// Draw snake
func (snake *Snake) Draw(r *sdl.Renderer) error {
	for _, Pos := range snake.Pos {
		rect := &sdl.Rect{X: Pos[0] * TILE, Y: Pos[1] * TILE, H: TILE - 1, W: TILE - 1}

		err := r.DrawRect(rect)

		if err != nil {
			return fmt.Errorf("could not draw snake: %v", err)
		}

		err = r.SetDrawColor(18, 88, 27, 1)

		if err != nil {
			return fmt.Errorf("could not set rect color: %v", err)
		}

		err = r.FillRect(rect)

		if err != nil {
			return fmt.Errorf("could not fill snake color: %v", err)
		}
	}

	return nil
}
