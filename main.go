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
	vk  virtDev.VirtKeyboard
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

	vk, err := virtDev.NewVirtKeyboard()
	if err != nil {
		fmt.Println("Error while creating virtDev: ", err)
		return
	}

	p := tea.NewProgram(
		model{
			ui:      ui.NewUi(d.GetRange()),
			dev:     *d,
			vk:      *vk,
			pressed: false,
		}, tea.WithAltScreen())

	d.SpawnEventListenerThread(p)

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
		}
	case tea.WindowSizeMsg:
		m.ui.UpdateDimensions(msg.Width, msg.Height)
	case device.InputEvents:
		m.handleInputEvents(msg)
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
