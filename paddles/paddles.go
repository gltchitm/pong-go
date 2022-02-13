package paddles

import (
	"github.com/gltchitm/pong-go/consts"
	"github.com/veandco/go-sdl2/sdl"
)

type Paddles struct {
	LeftY    int
	RightY   int
	renderer *sdl.Renderer
}

func NewPaddles(renderer *sdl.Renderer) *Paddles {
	paddles := Paddles{
		LeftY:    int(consts.WindowHeight/2 - consts.PaddleHeight/2),
		RightY:   int(consts.WindowHeight/2 - consts.PaddleHeight/2),
		renderer: renderer,
	}

	return &paddles
}

func drawPaddle(x, y int, renderer *sdl.Renderer) error {
	rect := sdl.Rect{X: int32(x), Y: int32(y), W: consts.PaddleWidth, H: consts.PaddleHeight}

	err := renderer.SetDrawColor(
		consts.ForegroundR,
		consts.ForegroundG,
		consts.ForegroundB,
		consts.ForegroundA,
	)
	if err != nil {
		return err
	}

	err = renderer.FillRect(&rect)
	if err != nil {
		return err
	}

	err = renderer.SetDrawColor(
		consts.BackgroundR,
		consts.BackgroundG,
		consts.BackgroundB,
		consts.BackgroundA,
	)
	if err != nil {
		return err
	}

	return nil
}

func (paddles *Paddles) Draw() error {
	err := drawPaddle(consts.WindowPadding, paddles.LeftY, paddles.renderer)
	if err != nil {
		return err
	}

	err = drawPaddle(
		int(consts.WindowWidth-consts.WindowPadding-consts.PaddleWidth),
		paddles.RightY,
		paddles.renderer,
	)

	return err
}
