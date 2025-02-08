package device

import (
	"bytes"
	"encoding/binary"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"go29/event_codes"
	"syscall"
	"time"
)

type WheelTurnMsg struct{ Value int }
type BreakPedalMsg struct{ Value int }
type ClutchPedalMsg struct{ Value int }
type ThrottlePedalMsg struct{ Value int }
type ButtonMsg struct{ Value int }
type DpadMsg struct {
	Value int
	Code  int
}

func (d *Device) PrintEvents() {
	for true {
		var data []byte = make([]byte, 300)

		fd, err := syscall.Open(d.dev_name, syscall.O_RDONLY, 0644)
		if err != nil {
			fmt.Println("Error opening dev event", err)
		}

		_, err = syscall.Read(fd, data)
		if err != nil {
			fmt.Println("Error reading input_event:", err)
		}

		var event InputEvent
		reader := bytes.NewReader(data)
		err = binary.Read(reader, binary.LittleEndian, &event)
		if err != nil {
			fmt.Println("binary.Read Error:", err)
			return
		}

		// fmt.Println(event)

		switch event.Code {
		case event_codes.ABS_X:
			// fmt.Println(event.Value)
			d.p.Send(func() tea.Msg {
				return WheelTurnMsg{Value: int(event.Value)}
			}())
		case event_codes.ABS_Y:
			d.p.Send(func() tea.Msg {
				return ClutchPedalMsg{Value: int(event.Value)}
			}())
		case event_codes.ABS_Z:
			d.p.Send(func() tea.Msg {
				return ThrottlePedalMsg{Value: int(event.Value)}
			}())
		case event_codes.ABS_RZ:
			d.p.Send(func() tea.Msg {
				return BreakPedalMsg{Value: int(event.Value)}
			}())
		case event_codes.ABS_RY:
			d.p.Send(func() tea.Msg {
				return ButtonMsg{Value: int(event.Value)}
			}())
		case event_codes.ABS_HAT0Y:
			fallthrough
		case event_codes.ABS_HAT0X:
			d.p.Send(func() tea.Msg {
				return DpadMsg{Code: int(event.Code), Value: int(event.Value)}
			}())
		}

		time.Sleep(1 * time.Millisecond)
	}
}
