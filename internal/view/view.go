package view

import (
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	_ "github.com/veandco/go-sdl2/ttf"

	config_api "dido/internal/config"
	"dido/internal/context"
	_ "dido/internal/cursor"
	"dido/internal/textstorage"
	_ "dido/internal/textstorage"
)

func init() {
	runtime.LockOSThread()
}

type View struct {
	renderer    *sdl.Renderer
	config      *config_api.Config
	font        *ttf.Font
	bgColor     sdl.Color
	bufferColor sdl.Color
}

func NewView(config *config_api.Config) View {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("SDL Init: %v", err)
	}
	window, err := sdl.CreateWindow(
		"Dido",
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		800, 600,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		log.Fatalf("SDL Window Init: %v", err)
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("SDL Renderer Init: %v", err)
	}
	if err := ttf.Init(); err != nil {
		log.Fatalf("TTF Init: %v", err)
	}
	fontPath, err := config.FontPath()
	if err != nil {
		log.Fatalf("Font Path Loading: %v", err)
	}
	font, err := ttf.OpenFont(fontPath, 20)
	if err != nil {
		log.Fatalf("Open font: %v", err)
	}
	return View{
		renderer:    renderer,
		config:      config,
		font:        font,
		bgColor:     sdl.Color{R: 51, G: 51, B: 51},
		bufferColor: sdl.Color{R: 255, G: 248, B: 231},
	}
}

func (v *View) Draw(ctx *context.Context) error {
	v.renderer.SetDrawColor(v.bgColor.R, v.bgColor.G, v.bgColor.B, 255)
	v.renderer.Clear()

	err := v.drawText(&ctx.Buffer, 0, 0)

	v.renderer.Present()
	return err
}

func (v *View) drawText(buffer textstorage.TextStorage, xOffset int32, yOffset int32) error {
	for i, line := range *buffer.AsLines() {
		textSurface, err := v.font.RenderUTF8Solid(line, v.bufferColor)
		if err != nil {
			return err
		}
		defer textSurface.Free()

		textTexture, err := v.renderer.CreateTextureFromSurface(textSurface)
		if err != nil {
			return err
		}
		defer textTexture.Destroy()

		dstRect := sdl.Rect{
			X: xOffset,
			Y: yOffset + int32(i)*textSurface.H,
			W: textSurface.W,
			H: textSurface.H,
		}
		if err := v.renderer.Copy(textTexture, nil, &dstRect); err != nil {
			return err
		}
	}

	return nil
}

func (v *View) Close() error {
	v.font.Close()
	ttf.Quit()
	err := v.renderer.Clear()
	sdl.Quit()
	return err
}
