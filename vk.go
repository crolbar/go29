package main

import (
	"errors"
	"fmt"
	"go29/virtDev"
)

func (m *model) startVK() error {
	vk, err := virtDev.NewVirtKeyboard()
	if err != nil {
		return errors.New(fmt.Sprintf("Error while creating virtDev: %s", err))
	}

	m.vk = vk

	m.ui.VirtDevButton.Toggle()

	return nil
}

func (m *model) stopVK() {
	m.vk.DestroyDev()
	m.vk = nil

	m.ui.VirtDevButton.Toggle()
}

func (m *model) ToggleVK() error {
	if !m.ui.VirtDevButton.GetState() {
		return m.startVK()
	}
	m.stopVK()
	return nil
}
