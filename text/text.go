package text

import (
	"github.com/gltchitm/pong-go/consts"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func Text(
	fontSize int,
	color sdl.Color,
	message string,
	renderer *sdl.Renderer,
) (*sdl.Surface, *sdl.Texture, error) {
	font, err := ttf.OpenFont(consts.FontPath, fontSize)
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
