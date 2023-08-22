package text

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"

	_ "embed"
)

//go:embed font/font.ttf
var fontData []byte

func Text(
	fontSize int,
	color sdl.Color,
	message string,
	renderer *sdl.Renderer,
) (*sdl.Surface, *sdl.Texture, error) {
	rw, err := sdl.RWFromMem(fontData)
	if err != nil {
		return nil, nil, err
	}

	font, err := ttf.OpenFontRW(rw, 1, fontSize)
	if err != nil {
		return nil, nil, err
	}
	defer font.Close()

	surface, err := font.RenderUTF8Solid(message, color)
	if err != nil {
		return nil, nil, err
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return nil, nil, err
	}

	return surface, texture, nil
}
