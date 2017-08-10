package main

import (
	"testing"
)

func TestReset(t *testing.T) {
	fruit := &Fruit{X: 10, Y: 10}

	oldX := fruit.X
	oldY := fruit.Y

	fruit.Reset()

	if fruit.X == oldX {
		t.Errorf("Didn't reset X value")
	}

	if fruit.Y == oldY {
		t.Errorf("Didn't reset Y value")
	}
}
