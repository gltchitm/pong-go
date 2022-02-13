package main

import (
	"time"

	"github.com/gltchitm/pong-go/ball"
	"github.com/gltchitm/pong-go/consts"
	"github.com/gltchitm/pong-go/paddles"
	"github.com/gltchitm/pong-go/scoreboard"
	"github.com/gltchitm/pong-go/welcome"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	keyboardStateActive = iota
	keyboardStateBackground
	keyboardStateInactive
)

type keyboardState struct {
	w    int
	s    int
	up   int
	down int
}

func clamp(value, min, max int) int {
	if value > max {
		return max
	} else if value < min {
		return min
	}

	return value
}
func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}

	defer sdl.Quit()

	window, err := sdl.CreateWindow(
		consts.Title,
		sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED,
		consts.WindowHeight,
		consts.WindowWidth,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	err = ttf.Init()
	if err != nil {
		panic(err)
	}

	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	keyboardState := keyboardState{
		w:    keyboardStateInactive,
		s:    keyboardStateInactive,
		up:   keyboardStateInactive,
		down: keyboardStateInactive,
	}

	started := false
	ticksUntilStart := consts.TicksUntilStart

	ball, err := ball.NewBall(renderer)
	if err != nil {
		panic(err)
	}

	paddles := paddles.NewPaddles(renderer)
	scoreboard := scoreboard.NewScoreboard(renderer)

game:
	for {
		tickStart := time.Now()

		err = renderer.Clear()
		if err != nil {
			panic(err)
		}

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				break game
			case *sdl.KeyboardEvent:
				if event.Type == sdl.KEYDOWN {
					switch event.Keysym.Scancode {
					case sdl.SCANCODE_SPACE:
						started = true
					case sdl.SCANCODE_W:
						if keyboardState.s == keyboardStateActive {
							keyboardState.s = keyboardStateBackground
						}
						keyboardState.w = keyboardStateActive
					case sdl.SCANCODE_S:
						if keyboardState.w == keyboardStateActive {
							keyboardState.w = keyboardStateBackground
						}
						keyboardState.s = keyboardStateActive
					case sdl.SCANCODE_UP:
						if keyboardState.down == keyboardStateActive {
							keyboardState.down = keyboardStateBackground
						}
						keyboardState.up = keyboardStateActive
					case sdl.SCANCODE_DOWN:
						if keyboardState.up == keyboardStateActive {
							keyboardState.up = keyboardStateBackground
						}
						keyboardState.down = keyboardStateActive
					}
				} else if event.Type == sdl.KEYUP {
					switch event.Keysym.Scancode {
					case sdl.SCANCODE_W:
						if keyboardState.s == keyboardStateBackground {
							keyboardState.s = keyboardStateActive
						}
						keyboardState.w = keyboardStateInactive
					case sdl.SCANCODE_S:
						if keyboardState.w == keyboardStateBackground {
							keyboardState.w = keyboardStateActive
						}
						keyboardState.s = keyboardStateInactive
					case sdl.SCANCODE_UP:
						if keyboardState.down == keyboardStateBackground {
							keyboardState.down = keyboardStateActive
						}
						keyboardState.up = keyboardStateInactive
					case sdl.SCANCODE_DOWN:
						if keyboardState.up == keyboardStateBackground {
							keyboardState.up = keyboardStateActive
						}
						keyboardState.down = keyboardStateInactive
					}
				}
			}
		}

		if started {
			if keyboardState.w == keyboardStateActive {
				paddles.LeftY = clamp(
					paddles.LeftY-consts.PaddleSpeed,
					consts.WindowPadding,
					consts.WindowHeight-consts.PaddleHeight-consts.WindowPadding,
				)
			}
			if keyboardState.s == keyboardStateActive {
				paddles.LeftY = clamp(
					paddles.LeftY+consts.PaddleSpeed,
					consts.WindowPadding,
					consts.WindowHeight-consts.PaddleHeight-consts.WindowPadding,
				)
			}

			if keyboardState.up == keyboardStateActive {
				paddles.RightY = clamp(
					paddles.RightY-consts.PaddleSpeed,
					consts.WindowPadding,
					consts.WindowHeight-consts.PaddleHeight-consts.WindowPadding,
				)
			}
			if keyboardState.down == keyboardStateActive {
				paddles.RightY = clamp(
					paddles.RightY+consts.PaddleSpeed,
					consts.WindowPadding,
					consts.WindowHeight-consts.PaddleHeight-consts.WindowPadding,
				)
			}

			if ticksUntilStart > 0 {
				for _, arrowTickRange := range consts.ArrowTickRanges {
					if ticksUntilStart > arrowTickRange[0] && ticksUntilStart < arrowTickRange[1] {
						err = ball.DrawArrow()
						if err != nil {
							panic(err)
						}
					}
				}

				ticksUntilStart--
			} else {
				ball.X += ball.XVelocity
				ball.Y += ball.YVelocity
			}

			if ball.X < consts.WindowPadding+consts.BallRadius {
				scoreboard.RightScore++
				ball.Center()
				err = ball.RandomizeVelocity()
				if err != nil {
					panic(err)
				}
				ticksUntilStart = consts.TicksUntilStart
			} else if ball.X > consts.WindowWidth-consts.WindowPadding-consts.BallRadius {
				scoreboard.LeftScore++
				ball.Center()
				err = ball.RandomizeVelocity()
				if err != nil {
					panic(err)
				}
				ticksUntilStart = consts.TicksUntilStart
			}

			if ball.Y < consts.WindowPadding+consts.BallRadius {
				ball.Y = consts.WindowPadding + consts.BallRadius
				ball.YVelocity *= -1
			} else if ball.Y > consts.WindowHeight-consts.WindowPadding-consts.BallRadius {
				ball.Y = consts.WindowHeight - consts.WindowPadding - consts.BallRadius
				ball.YVelocity *= -1
			}

			if ball.X < consts.WindowPadding*2+consts.PaddleWidth+consts.BallRadius &&
				ball.Y+consts.BallRadius > paddles.LeftY &&
				ball.Y-consts.BallRadius < paddles.LeftY+consts.PaddleHeight {
				ball.X = consts.WindowPadding*2 + consts.PaddleWidth + consts.BallRadius
				ball.XVelocity *= -1
			}
			if ball.X > consts.WindowWidth-consts.WindowPadding*2-consts.PaddleWidth-consts.BallRadius &&
				ball.Y+consts.BallRadius > paddles.RightY &&
				ball.Y-consts.BallRadius < paddles.RightY+consts.PaddleHeight {
				ball.X = consts.WindowWidth - consts.WindowPadding*2 - consts.PaddleWidth - consts.BallRadius
				ball.XVelocity *= -1
			}

			err = renderer.SetDrawColor(
				consts.ForegroundR,
				consts.ForegroundG,
				consts.ForegroundB,
				consts.ForegroundA,
			)
			if err != nil {
				panic(err)
			}

			err = renderer.DrawLine(consts.WindowWidth/2, 0, consts.WindowWidth/2, consts.WindowHeight)
			if err != nil {
				panic(err)
			}

			err = renderer.SetDrawColor(
				consts.BackgroundR,
				consts.BackgroundG,
				consts.BackgroundB,
				consts.BackgroundA,
			)
			if err != nil {
				panic(err)
			}

			err = ball.Draw()
			if err != nil {
				panic(err)
			}

			err = paddles.Draw()
			if err != nil {
				panic(err)
			}

			err = scoreboard.Draw()
			if err != nil {
				panic(err)
			}
		} else {
			err = welcome.DrawWelcome(renderer)
			if err != nil {
				panic(err)
			}
		}

		renderer.Present()

		timeUntilNextTick := consts.TickDuration - time.Since(tickStart).Milliseconds()
		if timeUntilNextTick < 0 {
			continue
		}

		sdl.Delay(uint32(timeUntilNextTick))
	}
}
