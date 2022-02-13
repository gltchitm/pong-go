package ball

import (
	"crypto/rand"
	"math/big"

	"github.com/gltchitm/pong-go/consts"
	"github.com/veandco/go-sdl2/sdl"
)

type Ball struct {
	X         int
	Y         int
	XVelocity int
	YVelocity int
	renderer  *sdl.Renderer
}

func NewBall(renderer *sdl.Renderer) *Ball {
	ball := Ball{
		X:         0,
		Y:         0,
		XVelocity: 0,
		YVelocity: 0,
		renderer:  renderer,
	}

	ball.Center()
	ball.RandomizeVelocity()

	return &ball
}

func (ball *Ball) Center() {
	ball.X = int(consts.WindowWidth / 2)
	ball.Y = int(consts.WindowHeight / 2)
}

func (ball *Ball) RandomizeVelocity() error {
	xRand, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return err
	}

	if xRand.Int64() == 0 {
		ball.XVelocity = consts.BallSpeed
	} else if xRand.Int64() == 1 {
		ball.XVelocity = -consts.BallSpeed
	}

	yRand, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		return err
	}

	if yRand.Int64() == 0 {
		ball.YVelocity = consts.BallSpeed
	} else if yRand.Int64() == 1 {
		ball.YVelocity = -consts.BallSpeed
	}

	return nil
}

func (ball *Ball) Draw() {
	ball.renderer.SetDrawColor(
		consts.ForegroundR,
		consts.ForegroundG,
		consts.ForegroundB,
		consts.ForegroundA,
	)

	for w := 0; w < consts.BallRadius*2; w++ {
		for h := 0; h < consts.BallRadius*2; h++ {
			dx := consts.BallRadius - w
			dy := consts.BallRadius - h

			if dx*dx+dy*dy <= consts.BallRadius*consts.BallRadius {
				ball.renderer.DrawPoint(int32(ball.X+dx), int32(ball.Y+dy))
			}
		}
	}

	ball.renderer.SetDrawColor(
		consts.BackgroundR,
		consts.BackgroundG,
		consts.BackgroundB,
		consts.BackgroundA,
	)
}

func (ball *Ball) DrawArrow() error {
	ball.renderer.SetDrawColor(
		consts.ForegroundR,
		consts.ForegroundG,
		consts.ForegroundB,
		consts.ForegroundA,
	)

	if ball.XVelocity > 0 && ball.YVelocity < 0 {
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
	} else if ball.XVelocity > 0 && ball.YVelocity > 0 {
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
	} else if ball.XVelocity < 0 && ball.YVelocity > 0 {
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2+consts.BallRadius+consts.ArrowLength,
		)
	} else if ball.XVelocity < 0 && ball.YVelocity < 0 {
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
		ball.renderer.DrawLine(
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
			consts.WindowWidth/2-consts.BallRadius/2-consts.BallRadius-consts.ArrowLength,
		)
	}

	ball.renderer.SetDrawColor(
		consts.BackgroundR,
		consts.BackgroundG,
		consts.BackgroundB,
		consts.BackgroundA,
	)

	return nil
}
