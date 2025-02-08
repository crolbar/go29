package ui

import (
	ec "go29/event_codes"
	"go29/ui/button"
	pb "go29/ui/progbar"

	"github.com/charmbracelet/lipgloss"
)

type SelectedBar int

const (
	Range SelectedBar = iota
	AutoCenter
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

		selectedBar: Range,

		height: 0,
		width:  0,
	}
}

func (u *Ui) UpdateDimensions(width, height int) {
	u.width = width
	u.height = height
}

var s lipgloss.Style = lipgloss.NewStyle()

var screenStyle lipgloss.Style = s.PaddingTop(2).
	PaddingBottom(2).
	PaddingRight(5).
	PaddingLeft(5)

func (u Ui) generateButtonIndicators() string {
	buttons := ""

	j := 0

	for i := 0; i < 3; i++ {
		lineMaxJ := j + 10
		line := ""

		for ; j < lineMaxJ && j < len(buttonMapKeys); j++ {
			line = lipgloss.JoinHorizontal(lipgloss.Left,
				line,
				(u.Buttons[buttonMapKeys[j]]).View(),
			)
		}

		if i == 2 {
			for di, db := range dpadMapKeys {
				line = lipgloss.JoinHorizontal(lipgloss.Left,
					line,
					((u.Dpad[iff((di&1) > 0, ec.ABS_HAT0X, ec.ABS_HAT0Y)])[db]).View(),
				)
			}
		}

		buttons = lipgloss.JoinVertical(lipgloss.Left,
			buttons,
			line,
		)
	}

	return buttons
}

func (u Ui) Render() string {
	wheelBar := lipgloss.JoinHorizontal(lipgloss.Left,
		u.WheelLeftBar.View(),
		u.WheelRightBar.View(),
	)

	if u.width-lipgloss.Width(wheelBar) < 40 {
		u.RangeBar.SetVertical(true)
		u.AutoCenterBar.SetVertical(true)
	}

	sliderBars := s.MarginLeft(10).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				u.RangeBar.View(),
				u.AutoCenterBar.View(),
			),
		)

	buttons := u.generateButtonIndicators()

	pedals :=
		lipgloss.JoinHorizontal(lipgloss.Left,
			u.ClutchBar.View(),
			u.BreakBar.View(),
			u.ThrottleBar.View(),
		)

	buttonsPedals := s.
		Height(u.height - lipgloss.Height(wheelBar) - 4).
		AlignVertical(lipgloss.Bottom).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				buttons,
				pedals,
			),
		)

	return screenStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left,
				wheelBar,
				buttonsPedals,
			),
			sliderBars,
		),
	)
}

func iff[T int](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
