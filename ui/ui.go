package ui

import (
	"go29/ui/progbar"

	"github.com/charmbracelet/lipgloss"
)

type Ui struct {
	WheelLeft  progbar.ProgBar
	WheelRight progbar.ProgBar
}

func NewUi(
	wheelLeft progbar.ProgBar,
	wheelRight progbar.ProgBar,
) Ui {
	return Ui{
		WheelLeft:  wheelLeft,
		WheelRight: wheelRight,
	}
}

func (u Ui) Render() string {
	return lipgloss.JoinHorizontal(lipgloss.Center,
		u.WheelLeft.View(),
		u.WheelRight.View(),
	)
}
