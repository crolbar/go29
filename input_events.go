package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go29/device"
	ec "go29/event_codes"
)

func (m *model) handleInputEvents(events device.InputEvents) {
	data := events.Data
	n := events.N

	var evt device.InputEvent

	reader := bytes.NewReader(data[0:24])

	for l, r := 0, 24; r < n; l, r = r, r+24 {
		reader.Reset(data[l:r])

		err := binary.Read(reader, binary.LittleEndian, &evt)
		if err != nil {
			fmt.Println("binary.Read Error:", err)
		}

		// skip SYN & unused KEY events
		if evt.Type == ec.EV_SYN || evt.Type == ec.EV_KEY {
			continue
		}

		m.ui.HandleInputEvent(evt)
	}
}
