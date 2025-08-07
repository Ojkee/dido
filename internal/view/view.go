package view

import "github.com/veandco/go-sdl2/sdl"

type View struct {
	window *sdl.Window

	bgColor     sdl.Color
	bufferColor sdl.Color
}

func NewView() View {
	window, err := sdl.CreateWindow(
		"Dido",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}
	return View{
		window:      window,
		bgColor:     sdl.Color{R: 51, G: 51, B: 51},
		bufferColor: sdl.Color{R: 255, G: 248, B: 231},
	}
}

func (v *View) Draw() error {
	surface, err := v.window.GetSurface()
	if err != nil {
		return err
	}

	v.drawBackground(surface)
	return nil
}

func (v *View) drawBackground(surface *sdl.Surface) {
	bgUint32 := sdl.MapRGB(
		surface.Format,
		v.bgColor.R,
		v.bgColor.G,
		v.bgColor.B,
	)
	_ = surface.FillRect(nil, bgUint32)
}

func (v *View) Update() error {
	err := v.window.UpdateSurface()
	return err
}

func (v *View) Close() error {
	err := v.window.Destroy()
	return err
}
