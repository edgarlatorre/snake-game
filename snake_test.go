package main

import (
	"testing"
)

func TestUpdateAppendPos(t *testing.T) {
	snake := &Snake{X: 10, Y: 10, Velx: 0, Vely: 0, Tail: 3, Pos: make([][]int32, 0)}

	snake.Update()

	if len(snake.Pos) != 1 {
		t.Errorf("Update did not set position")
	}

	if snake.Pos[0][0] != snake.X {
		t.Errorf("Update did not append X position")
	}

	if snake.Pos[0][1] != snake.Y {
		t.Errorf("Update did not append Y position")
	}

	snake.Update()

	if len(snake.Pos) != 2 {
		t.Errorf("Update did not append position")
	}
}

func TestUpdateResetX(t *testing.T) {
	snake := &Snake{X: 10 + WIDTH_IN_TILE, Y: 10, Velx: 0, Vely: 0, Tail: 3, Pos: make([][]int32, 0)}
	snake.Update()

	if snake.X != 0 {
		t.Errorf("Update did not reset X")
	}

	snake.X = -1
	snake.Update()

	if snake.X != WIDTH_IN_TILE {
		t.Errorf("Did not update X = WIDTH_IN_TILE if less than 0")
	}
}

func TestUpdateResetY(t *testing.T) {
	snake := &Snake{X: 10, Y: 10 + WIDTH_IN_TILE, Velx: 0, Vely: 0, Tail: 3, Pos: make([][]int32, 0)}
	snake.Update()

	if snake.Y != 0 {
		t.Errorf("Update did not reset Y if greater than WIDTH_IN_TILE")
	}

	snake.Y = -1
	snake.Update()

	if snake.Y != WIDTH_IN_TILE {
		t.Errorf("Did not update Y = WIDTH_IN_TILE if less than 0")
	}
}
