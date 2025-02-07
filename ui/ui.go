package ui

import (
	"go29/ui/progbar"

	"github.com/charmbracelet/lipgloss"
)

type SelectedBar int

const (
	WheelRange SelectedBar = iota
	AutoCenter
)

type Ui struct {
	WheelLeft   progbar.ProgBar
	WheelRight  progbar.ProgBar
	Throttle    progbar.ProgBar
	WheelRange  progbar.ProgBar
	AutoCenter  progbar.ProgBar
	selectedBar SelectedBar
}

func NewUi(
	wheelLeft progbar.ProgBar,
	wheelRight progbar.ProgBar,
	throttle progbar.ProgBar,
	wheelRange progbar.ProgBar,
	autoCenter progbar.ProgBar,
) Ui {
	return Ui{
		WheelLeft:   wheelLeft,
		WheelRight:  wheelRight,
		Throttle:    throttle,
		WheelRange:  wheelRange,
		AutoCenter: autoCenter,
		selectedBar: WheelRange,
	}
}

func (u Ui) Render() string {
	return lipgloss.JoinVertical(lipgloss.Center,
		lipgloss.JoinHorizontal(lipgloss.Center,
			u.WheelLeft.View(),
			u.WheelRight.View(),
		),
		lipgloss.JoinHorizontal(lipgloss.Center,
			u.Throttle.View(),
			u.WheelRange.View(),
			u.AutoCenter.View(),
		),
	)
}
