package scoreboard

import (
	"fmt"

	"github.com/gltchitm/pong-go/consts"
	"github.com/gltchitm/pong-go/text"
	"github.com/veandco/go-sdl2/sdl"
)

type Scoreboard struct {
	LeftScore  int
	RightScore int
	renderer   *sdl.Renderer
}

func NewScoreboard(renderer *sdl.Renderer) *Scoreboard {
	return &Scoreboard{
		LeftScore:  0,
		RightScore: 0,
		renderer:   renderer,
	}
}

func (scoreboard *Scoreboard) Draw() error {
	color := sdl.Color{
		R: consts.ForegroundR,
		G: consts.ForegroundG,
		B: consts.ForegroundB,
		A: consts.ForegroundA,
	}

	leftScoreSurface, leftScoreTexture, err := text.Text(
		consts.ScoreboardFontSize,
		color,
		fmt.Sprint(scoreboard.LeftScore),
		scoreboard.renderer,
	)
	if err != nil {
		return err
	}
	defer leftScoreSurface.Free()
	defer leftScoreTexture.Destroy()

	leftScoreRect := sdl.Rect{
		X: (consts.WindowWidth/2 - leftScoreSurface.W) / 2,
		Y: consts.ScoreboardTopPadding,
		W: leftScoreSurface.W,
		H: leftScoreSurface.H,
	}

	err = scoreboard.renderer.Copy(leftScoreTexture, nil, &leftScoreRect)
	if err != nil {
		return err
	}

	rightScoreSurface, rightScoreTexture, err := text.Text(
		consts.ScoreboardFontSize,
		color,
		fmt.Sprint(scoreboard.RightScore),
		scoreboard.renderer,
	)
	if err != nil {
		return err
	}
	defer rightScoreSurface.Free()
	defer rightScoreTexture.Destroy()

	rightScoreRect := sdl.Rect{
		X: consts.WindowWidth/2 + (consts.WindowWidth/2-rightScoreSurface.W)/2,
		Y: consts.ScoreboardTopPadding,
		W: rightScoreSurface.W,
		H: rightScoreSurface.H,
	}

	err = scoreboard.renderer.Copy(rightScoreTexture, nil, &rightScoreRect)
	if err != nil {
		return err
	}

	return nil
}
