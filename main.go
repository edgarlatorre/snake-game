package main

import (
	"fmt"
	"os"

	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}

	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Count not initialize ttf: %v", err)
	}

	window, r, err := sdl.CreateWindowAndRenderer(800, 800, sdl.WINDOW_SHOWN)

	if err != nil {
		return fmt.Errorf("Could not create window: %v", err)
	}

	defer window.Destroy()

	err = startScreen(r)

	if err != nil {
		return fmt.Errorf("Could not draw title %v", err)
	}

	time.Sleep(time.Second * 3)

	err = background(r)

	if err != nil {
		return fmt.Errorf("Could not draw background %v", err)
	}

	time.Sleep(time.Second * 3)

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

func title(r *sdl.Renderer) error {
	r.Clear()

	f, err := ttf.OpenFont("resources/fonts/pac-font.ttf", 20)

	if err != nil {
		return fmt.Errorf("Could not open font: %v", err)
	}

	defer f.Close()

	s, err := f.RenderUTF8_Solid("Snake Game", sdl.Color{R: 255, G: 100, B: 0, A: 255})

	if err != nil {
		return fmt.Errorf("Could not render title: %v", err)
	}

	defer s.Free()

	t, err := r.CreateTextureFromSurface(s)

	if err != nil {
		return fmt.Errorf("Could not create texture: %v", err)
	}

	defer t.Destroy()

	err = r.Copy(t, nil, nil)

	if err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	r.Present()

	return nil
}
