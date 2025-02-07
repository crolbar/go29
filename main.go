package main

import (
	"fmt"

	"go29/device"
	"go29/ui"
	pb "go29/ui/progbar"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dev device.Device
	ui  ui.Ui
}

func newModel() model {
	d := device.NewDevice()

	return model{
		ui: ui.NewUi(
			pb.NewProgBar("left", 3, 40,
				pb.WithMaxValue(32767),
				pb.WithDisabledRightBorder(),
				pb.WithReverse(),
			),
			pb.NewProgBar("right", 3, 40,
				pb.WithMaxValue(32767),
				pb.WithDisabledLeftBorder(),
			),
			pb.NewProgBar("throttle", 15, 13,
				pb.WithVertical(),
				pb.WithReverse(),
				pb.WithMaxValue(255),
			),
			pb.NewProgBar("range", 3, 40,
				pb.WithMaxValue(900),
				pb.WithMinValue(30),
				pb.WithValue(d.GetRange()),
				pb.WithSelected(),
			),
			pb.NewProgBar("autocenter", 3, 40,
				pb.WithMaxValue(100),
			),
		),
		dev: d,
	}
}

func main() {
	m := newModel()
	p := tea.NewProgram(m)

	m.dev.SetProgram(p)
	go m.dev.PrintEvents()

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
		case "h":
			m.ui.HandleSelectedBarLeft(&m.dev)
		case "l":
			m.ui.HandleSelectedBarRight(&m.dev)

		case "tab":
			m.ui.SelectNextBar()
		case "shift+tab":
			m.ui.SelectPrevBar()
		}

	case device.Send:
		if msg.Value < 32767 {
			m.ui.WheelLeftBar.SetValue(32767 - msg.Value)
			m.ui.WheelRightBar.SetValue(0)
		} else {
			m.ui.WheelLeftBar.SetValue(0)
			m.ui.WheelRightBar.SetValue(msg.Value - 32767)
		}
	case device.SendThrottle:
		m.ui.ThrottleBar.SetValue(255 - msg.Value)
	}

	return m, cmd
}

func (m model) View() string {
	return m.ui.Render()
}
