package consts

const (
	Title        = "Pong Go"
	SpaceToStart = "Press SPACE To Start!"

	FontPath = "./assets/font/font.ttf"

	TitleFontSize            = 64
	PressToStartFontSize     = 24
	TitlePressToStartPadding = 30

	BackgroundR = 0
	BackgroundG = 0
	BackgroundB = 0
	BackgroundA = 0

	ForegroundR = 255
	ForegroundG = 255
	ForegroundB = 255
	ForegroundA = 255

	WindowHeight  = 750
	WindowWidth   = 750
	WindowPadding = 15

	TickDuration    = 24
	TicksUntilStart = 36

	BallRadius  = 13
	BallSpeed   = 23
	ArrowLength = 40

	PaddleHeight = 200
	PaddleWidth  = 15
	PaddleSpeed  = 25

	ScoreboardFontSize   = 48
	ScoreboardTopPadding = 28
)

var ArrowTickRanges = [][]int{
	{6, 12},
	{18, 24},
	{30, 36},
}
