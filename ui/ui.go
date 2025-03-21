package ui

import (
	"github.com/crolbar/lipbalm/framebuffer"
	lbl "github.com/crolbar/lipbalm/layout"
	"go29/ui/button"
	pb "go29/ui/progbar"
)

type SelectedBar int

const (
	Range SelectedBar = iota
	AutoCenter
	ConstEffect
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
	ConstEffectBar
	Buttons // includes Dpad
	Dpad    // part of Buttons
	virtDevButton
)

type Ui struct {
	WheelLeftBar  pb.ProgBar
	WheelRightBar pb.ProgBar

	ThrottleBar pb.ProgBar
	BreakBar    pb.ProgBar
	ClutchBar   pb.ProgBar

	RangeBar       pb.ProgBar
	AutoCenterBar  pb.ProgBar
	ConstEffectBar pb.ProgBar

	selectedBar SelectedBar

	Buttons map[int]*button.Button
	Dpad    map[int]map[int]*button.Button

	VirtDevButton button.Button

	preRenders map[UiElement]string
	prevValues map[UiElement]int
	reqRender  map[UiElement]bool

	fb framebuffer.FrameBuffer

	Rects map[UiElement]lbl.Rect

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
		ConstEffectBar: pb.NewProgBar("const_effect", 3, 40,
			pb.WithMinValue(-10),
			pb.WithMaxValue(10),
		),

		Buttons: buttonsMap,
		Dpad:    dpadMap,

		VirtDevButton: *button.NewButton(
			"  'space'/'v'   \n to start/stop  \nvirtual keyboard",
			button.WithMargin(0),
		),

		preRenders: make(map[UiElement]string),
		prevValues: make(map[UiElement]int),
		reqRender:  make(map[UiElement]bool),

		selectedBar: Range,

		fb: framebuffer.NewFrameBuffer(0, 0),


		Rects: make(map[UiElement]lbl.Rect),

		height: 0,
		width:  0,
	}
}

func (u *Ui) UpdateDimensions(width, height int) {
	u.width = width
	u.height = height
	u.fb.Resize(width, height)
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
