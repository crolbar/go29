package ui

import (
	"go29/ui/button"
	pb "go29/ui/progbar"
)

type SelectedBar int

const (
	Range SelectedBar = iota
	AutoCenter
)

type UiElement int

const (
	Screen UiElement = iota
	WheelBar
	WheelBarLeft  // prev only
	WheelBarRight // prev only
	ThrottleBar
	BreakBar
	ClutchBar
	RangeBar
	AutoCenterBar
	Buttons // includes Dpad
	Dpad    // part of Buttons
)

type Ui struct {
	WheelLeftBar  pb.ProgBar
	WheelRightBar pb.ProgBar

	ThrottleBar pb.ProgBar
	BreakBar    pb.ProgBar
	ClutchBar   pb.ProgBar

	RangeBar      pb.ProgBar
	AutoCenterBar pb.ProgBar

	selectedBar SelectedBar

	Buttons map[int]*button.Button
	Dpad    map[int]map[int]*button.Button

	preRenders map[UiElement]string
	prevValues map[UiElement]int
	reqRender  map[UiElement]bool

	height int
	width  int
}

func NewUi(wRange int) Ui {
	setButtonMapKeys()
	setDpadMapKeys()

	return Ui{
		WheelLeftBar: pb.NewProgBar("left", 3, 40,
			pb.WithMaxValue(32767),
			pb.WithDisabledRightBorder(),
			pb.WithReverse(),
		),
		WheelRightBar: pb.NewProgBar("right", 3, 40,
			pb.WithMaxValue(32767),
			pb.WithDisabledLeftBorder(),
		),
		ThrottleBar: pb.NewProgBar("throttle", 15, 13,
			pb.WithVertical(),
			pb.WithReverse(),
			pb.WithMaxValue(255),
		),
		BreakBar: pb.NewProgBar("break", 15, 13,
			pb.WithVertical(),
			pb.WithReverse(),
			pb.WithMaxValue(255),
		),
		ClutchBar: pb.NewProgBar("clutch", 15, 13,
			pb.WithVertical(),
			pb.WithReverse(),
			pb.WithMaxValue(255),
		),
		RangeBar: pb.NewProgBar("range", 3, 40,
			pb.WithMaxValue(900),
			pb.WithMinValue(30),
			pb.WithValue(wRange),
			pb.WithSelected(),
		),
		AutoCenterBar: pb.NewProgBar("autocenter", 3, 40,
			pb.WithMaxValue(100),
		),

		Buttons: buttonsMap,
		Dpad:    dpadMap,

		preRenders: make(map[UiElement]string),
		prevValues: make(map[UiElement]int),
		reqRender:  make(map[UiElement]bool),

		selectedBar: Range,

		height: 0,
		width:  0,
	}
}

func (u *Ui) UpdateDimensions(width, height int) {
	u.width = width
	u.height = height
}

func (u Ui) Render() string {
	return u.preRenders[Screen]
}

func iff[T int](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
