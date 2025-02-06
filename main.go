package main

import (
	"fmt"

	"go29/device"
	"go29/ui"
	"go29/ui/progbar"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dev device.Device
	ui  ui.Ui
}

func newModel() model {
	return model{
		ui: ui.NewUi(
			progbar.NewProgBar("left", 3, 40),
			progbar.NewProgBar("right", 3, 40),
		),
		dev: device.NewDevice(),
	}
}

func main() {
	m := newModel()

	go m.dev.PrintEvents()

	// m.dev.SetAutocenter(5000)
	// m.dev.SetRange(500)

	// m.ui.Progbar.SetMaxValue(65535)
	m.ui.WheelLeft.SetMaxValue(32767)
	m.ui.WheelRight.SetMaxValue(32767)

	// m.ui.WheelLeft.SetMaxValue(200)
	// m.ui.WheelRight.SetMaxValue(200)

	m.ui.WheelLeft.Reverse(true)
	m.ui.WheelLeft.DisableRightBorder(true)
	m.ui.WheelRight.DisableLeftBorder(true)

	p := tea.NewProgram(m)

	m.dev.SetProgram(p)
	if _, err := p.Run(); err != nil {
		fmt.Println("Exited with Error: ", err)
		return
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "u":
			m.ui.WheelLeft.SetValue(20)
		case "i":
			m.ui.WheelLeft.SetValue(50)
		case "d":
			m.ui.WheelRight.SetValue(70)
		case "h":
			m.ui.WheelLeft.SetValue(200)
		case "+":
			m.ui.WheelLeft.SetValue(m.ui.WheelLeft.GetValue() + 1)
		case "-":
			m.ui.WheelLeft.SetValue(m.ui.WheelLeft.GetValue() - 1)
		}

	case device.Send:
		if msg.Value < 32767 {
			m.ui.WheelLeft.SetValue(32767 - msg.Value)
			m.ui.WheelRight.SetValue(0)
		} else {
			m.ui.WheelLeft.SetValue(0)
			m.ui.WheelRight.SetValue(msg.Value - 32767)
		}
	}

	return m, cmd
}

func (m model) View() string {
	return m.ui.Render()
}
