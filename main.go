package main

import (
	"fmt"
	"go29/device"
	"go29/ui"
	"go29/virtDev"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	dev device.Device
	vk  *virtDev.VirtKeyboard
	ui  ui.Ui

	pressed bool
}

func main() {
	d, err := device.NewDevice()
	if err != nil {
		fmt.Println("Error while creating device: ", err)
		return
	}
	defer d.CloseFD()

	p := tea.NewProgram(
		model{
			ui:      ui.NewUi(d.GetRange()),
			dev:     *d,
			vk:      nil,
			pressed: false,
		}, tea.WithAltScreen(), tea.WithMouseCellMotion())

	d.SpawnEventListenerThread(p)

	if _, err := p.Run(); err != nil {
		fmt.Println("Exited with Error: ", err)
		return
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var err error

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "left", "h":
			m.ui.HandleSelectedBarLeft(&m.dev)
		case "right", "l":
			m.ui.HandleSelectedBarRight(&m.dev)
		case "tab", "j", "down":
			m.ui.SelectNextBar()
		case "shift+tab", "k", "up":
			m.ui.SelectPrevBar()
		case "r":
			if m.vk != nil {
				err = m.vk.ReloadConfig()
			}
		case "v", " ":
			err = m.ToggleVK()
		}
	case tea.MouseMsg:
		m.handleMouseEvent(tea.MouseEvent(msg))
	case tea.WindowSizeMsg:
		m.ui.UpdateDimensions(msg.Width, msg.Height)
	case device.InputEvents:
		m.handleInputEvents(msg)
	}

	if err != nil {
		tea.ErrInterrupted = err
		return m, tea.Interrupt
	}

	m.ui.PreRender()

	return m, cmd
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	return m.ui.Render()
}
