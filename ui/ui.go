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
	}
}

func (u Ui) Render() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center,
			u.WheelLeftBar.View(),
			u.WheelRightBar.View(),
		),
		lipgloss.JoinHorizontal(lipgloss.Center,
			u.ThrottleBar.View(),
			u.RangeBar.View(),
			u.AutoCenterBar.View(),
		),
	)
}
