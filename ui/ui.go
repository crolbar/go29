package ui

import (
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
	selectedBar   SelectedBar

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
		selectedBar:   Range,

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
		Margin(1).
		Height(u.height - 5).
		Width(u.width - 5)

	wheelBar := lipgloss.JoinHorizontal(lipgloss.Left,
		u.WheelLeftBar.View(),
		u.WheelRightBar.View(),
	)

	sliderBars := s.MarginLeft(10).
		Render(
			lipgloss.JoinVertical(lipgloss.Left,
				u.RangeBar.View(),
				u.AutoCenterBar.View(),
			),
		)

	throttle := u.ThrottleBar.View()
	pedals := s.Height(u.height - lipgloss.Height(wheelBar) - lipgloss.Height(throttle)/2).
		AlignVertical(lipgloss.Bottom).
		Render(
			lipgloss.JoinHorizontal(lipgloss.Left,
				throttle,
			),
		)

	return screenStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left,
				wheelBar,
				pedals,
			),
			sliderBars,
		),
	)
}
