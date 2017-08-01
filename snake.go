package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Snake struct
type Snake struct {
	x, y, velx, vely int32
	pos              [][]int32
	tail             int
}

// Update snake
func (snake *Snake) Update() {
	snake.x += snake.velx
	snake.y += snake.vely

	if snake.x > WIDTH_IN_TILE {
		snake.x = 0
	}

	if snake.x < 0 {
		snake.x = WIDTH_IN_TILE
	}

	if snake.y > WIDTH_IN_TILE {
		snake.y = 0
	}

	if snake.y < 0 {
		snake.y = WIDTH_IN_TILE
	}

	snake.pos = append(snake.pos, []int32{snake.x, snake.y})

	for len(snake.pos) > snake.tail {
		fmt.Println("Slice", snake.pos[1:])
		snake.pos = snake.pos[1:]

	}
}

// Draw snake
func (snake *Snake) Draw(r *sdl.Renderer) error {
	for _, pos := range snake.pos {
		rect := &sdl.Rect{X: pos[0] * TILE, Y: pos[1] * TILE, H: TILE - 1, W: TILE - 1}

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
