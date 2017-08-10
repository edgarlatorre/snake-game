package main

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	WIDTH         = 800
	SPEED         = 100
	HEIGHT        = 800
	TILE          = 20
	WIDTH_IN_TILE = WIDTH / TILE
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	var event sdl.Event

	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}

	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Count not initialize ttf: %v", err)
	}

	window, r, err := sdl.CreateWindowAndRenderer(WIDTH, HEIGHT, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("Could not create window: %v", err)
	}

	defer window.Destroy()

	err = startScreen(r)

	if err != nil {
		return fmt.Errorf("Could not draw title %v", err)
	}

	time.Sleep(time.Second * 3)

	s, err := newScene(r)

	if err != nil {
		return fmt.Errorf("Could not create a new scene %v", err)
	}

	defer s.destroy()

	snake := &Snake{X: 10, Y: 10, Velx: 0, Vely: 0, Tail: 3, action: "right", Pos: make([][]int32, 0)}
	fruit := &Fruit{X: 10, Y: 10}

	running := true

	for range time.Tick(100 * time.Millisecond) {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_UP {
					snake.Up()
				}

				if t.Keysym.Sym == sdl.K_DOWN {
					snake.Down()
				}

				if t.Keysym.Sym == sdl.K_LEFT {
					snake.Left()
				}

				if t.Keysym.Sym == sdl.K_RIGHT {
					snake.Right()
				}
			}
		}

		switch snake.action {
		case "up":
			snake.Y--
		case "down":
			snake.Y++
		case "left":
			snake.X--
		case "right":
			snake.X++
		}

		snake.Update()

		err = s.draw(r, snake, fruit)

		if err != nil {
			return fmt.Errorf("Could not draw a scene %v", err)
		}

		if !running {
			break
		}
	}

	return nil
}

func background(r *sdl.Renderer) error {
	r.Clear()
	t, err := img.LoadTexture(r, "resources/images/Background.png")

	if err != nil {
		return fmt.Errorf("Could not load backgrund image: %v", err)
	}

	defer t.Destroy()

	err = r.Copy(t, nil, nil)

	if err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}

func startScreen(r *sdl.Renderer) error {
	r.Clear()
	t, err := img.LoadTexture(r, "resources/images/StartScreen.png")

	if err != nil {
		return fmt.Errorf("Could not load backgrund image: %v", err)
	}

	err = r.Copy(t, nil, nil)

	if err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}
