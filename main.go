package main

import (
	"fmt"
	"go29/device"
	"go29/ui"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dev device.Device
	ui  ui.Ui
}

func newModel() (*model, error) {
	d, err := device.NewDevice()
	if err != nil {
		return nil, err
	}

	return &model{
		ui:  ui.NewUi(d.GetRange()),
		dev: *d,
	}, nil
}

func main() {
	m, err := newModel()
	if err != nil {
		fmt.Println("Error while creating model: ", err)
		return
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	m.dev.SetProgram(p)
	go m.dev.PrintEvents()

	if _, err := p.Run(); err != nil {
		fmt.Println("Exited with Error: ", err)
		return
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "h":
			m.ui.HandleSelectedBarLeft(&m.dev)
		case "l":
			m.ui.HandleSelectedBarRight(&m.dev)

		case "tab":
			m.ui.SelectNextBar()
		case "shift+tab":
			m.ui.SelectPrevBar()
		}
	case tea.WindowSizeMsg:
		m.ui.UpdateDimensions(msg.Width, msg.Height)

	case device.WheelTurnMsg:
		if msg.Value < 32767 {
			m.ui.WheelLeftBar.SetValue(32767 - msg.Value)
			m.ui.WheelRightBar.SetValue(0)
		} else {
			m.ui.WheelLeftBar.SetValue(0)
			m.ui.WheelRightBar.SetValue(msg.Value - 32767)
		}
	case device.ThrottlePedalMsg:
		m.ui.ThrottleBar.SetValue(255 - msg.Value)
	case device.BreakPedalMsg:
		m.ui.BreakBar.SetValue(255 - msg.Value)
	case device.ClutchPedalMsg:
		m.ui.ClutchBar.SetValue(255 - msg.Value)
	case device.ButtonMsg:
		m.ui.Buttons[msg.Value].Toggle()
	case device.DpadMsg:
		if msg.Value == 0 {
			m.ui.Dpad[msg.Code][-1].Release()
			m.ui.Dpad[msg.Code][1].Release()
			break
		}

		m.ui.Dpad[msg.Code][msg.Value].Toggle()
	}

	return m, cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.ui.Render()
}
