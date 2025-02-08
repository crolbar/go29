package ui

import (
	"go29/ui/button"
	"go29/ui/progbar"

	"github.com/charmbracelet/lipgloss"
)

type SelectedBar int

const (
	Range SelectedBar = iota
	AutoCenter
)

type Ui struct {
	WheelLeftBar  progbar.ProgBar
	WheelRightBar progbar.ProgBar
	ThrottleBar   progbar.ProgBar
	RangeBar      progbar.ProgBar
	AutoCenterBar progbar.ProgBar

	selectedBar SelectedBar

	Button button.Button

	height int
	width  int
}

func NewUi(
	wheelLeftBar progbar.ProgBar,
	wheelRightBar progbar.ProgBar,
	throttleBar progbar.ProgBar,
	wheelRangeBar progbar.ProgBar,
	autoCenterBar progbar.ProgBar,
) Ui {
	return Ui{
		WheelLeftBar:  wheelLeftBar,
		WheelRightBar: wheelRightBar,
		ThrottleBar:   throttleBar,
		RangeBar:      wheelRangeBar,
		AutoCenterBar: autoCenterBar,

		Button: button.NewButton("test"),

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

func (u Ui) Render() string {
	screenStyle := s.PaddingTop(2).
		PaddingBottom(2).
		PaddingRight(5).
		PaddingLeft(5).
		Height(u.height).
		Width(u.width)

	wheelBar := lipgloss.JoinHorizontal(lipgloss.Left,
		u.WheelLeftBar.View(),
		u.WheelRightBar.View(),
	)

	sliderBars := s.
		MarginLeft(10).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				u.RangeBar.View(),
				u.AutoCenterBar.View(),
			),
		)

	var buttons string
	{
		for i := 0; i < 2; i++ {
			line := ""

			for j := 0; j < 10; j++ {
				line = lipgloss.JoinHorizontal(lipgloss.Left,
					line,
					u.Button.View(),
				)
			}

			buttons = lipgloss.JoinVertical(lipgloss.Left,
				buttons,
				line,
			)
		}
	}

	throttle := u.ThrottleBar.View()
	pedals :=
		lipgloss.JoinHorizontal(lipgloss.Left,
			throttle,
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
