package view

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	_ "github.com/veandco/go-sdl2/ttf"

	_ "dido/internal/cursor"
	_ "dido/internal/textstorage"
)

func init() {
	runtime.LockOSThread()
}

type View struct {
	window *sdl.Window

	bgColor     sdl.Color
	bufferColor sdl.Color
}

func NewView() View {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("SDL Init: %v", err)
	}
	if err := ttf.Init(); err != nil {
		log.Fatalf("TTF Init: %v", err)
	}
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
	defer surface.Free()

	v.drawBackground(surface)

	// TEXT
	// bufferUint32 := sdl.MapRGB(
	// 	surface.Format,
	// 	v.bufferColor.R,
	// 	v.bufferColor.G,
	// 	v.bufferColor.B,
	// )
	// _ = surface.FillRect(nil, bufferUint32)
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
	sdl.Quit()
	ttf.Quit()
	return err
}
