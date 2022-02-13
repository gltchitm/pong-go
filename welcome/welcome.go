package welcome

import (
	"github.com/gltchitm/pong-go/consts"
	"github.com/gltchitm/pong-go/text"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawWelcome(renderer *sdl.Renderer) error {
	color := sdl.Color{
		R: consts.ForegroundR,
		G: consts.ForegroundG,
		B: consts.ForegroundB,
		A: consts.ForegroundA,
	}

	titleSurface, titleTexture, err := text.Text(
		consts.TitleFontSize,
		color,
		consts.Title,
		renderer,
	)
	if err != nil {
		return err
	}
	defer titleSurface.Free()
	defer titleTexture.Destroy()

	startSurface, startTexture, err := text.Text(
		consts.PressToStartFontSize,
		color,
		consts.SpaceToStart,
		renderer,
	)
	if err != nil {
		return err
	}
	defer startSurface.Free()
	defer startTexture.Destroy()

	titleRect := sdl.Rect{
		X: consts.WindowWidth/2 - titleSurface.W/2,
		Y: consts.WindowHeight/2 - titleSurface.H/2 - startSurface.H/2,
		W: titleSurface.W,
		H: titleSurface.H,
	}
	startRect := sdl.Rect{
		X: consts.WindowWidth/2 - startSurface.W/2,
		Y: consts.WindowHeight/2 - startSurface.H/2 + titleSurface.H/2 + consts.TitlePressToStartPadding,
		W: startSurface.W,
		H: startSurface.H,
	}

	err = renderer.Copy(titleTexture, nil, &titleRect)
	if err != nil {
		return err
	}

	err = renderer.Copy(startTexture, nil, &startRect)

	return err
}
