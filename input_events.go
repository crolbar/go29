package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	lbl "github.com/crolbar/lipbalm/layout"
	"go29/device"
	ec "go29/event_codes"
	"go29/ui"
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

		// skip SYN
		if evt.Type == ec.EV_SYN {
			continue
		}

		if m.vk != nil {
			m.vk.HandleInputEvent(evt)
		}
		m.ui.HandleInputEvent(evt)
	}
}

func (m *model) handleMouseEvent(mv tea.MouseEvent) {
	var (
		x = uint16(mv.X)
		y = uint16(mv.Y)

		rangeRect       = m.ui.Rects[ui.RangeBar]
		autoCenterRect  = m.ui.Rects[ui.AutoCenterBar]
		constEffectRect = m.ui.Rects[ui.ConstEffectBar]
	)

	if isInRect(rangeRect, x, y) {
		var (
			w      = float32(rangeRect.Width) - 1 - 2
			xpos   = float32(clamp(int(x-rangeRect.X), 1, int(w+1))) - 1
			perc   = xpos / w
			maxVal = uint16(m.ui.RangeBar.GetMaxValue())
			minVal = uint16(m.ui.RangeBar.GetMinValue())
			newVal = minVal + uint16(perc*float32(maxVal-minVal))
		)

		m.ui.RangeBar.SetValue(int(newVal))
		m.dev.SetRange(int(newVal))

	} else if isInRect(autoCenterRect, x, y) {
		var (
			w      = float32(autoCenterRect.Width) - 1 - 2
			xpos   = float32(clamp(int(x-autoCenterRect.X), 1, int(w+1))) - 1
			perc   = xpos / w
			maxVal = uint16(m.ui.AutoCenterBar.GetMaxValue())
			minVal = uint16(m.ui.AutoCenterBar.GetMinValue())
			newVal = minVal + uint16(perc*float32(maxVal-minVal))
		)

		m.ui.AutoCenterBar.SetValue(int(newVal))
		m.dev.SetAutocenter(int(newVal))
	} else if isInRect(constEffectRect, x, y) {
		var (
			w      = float32(constEffectRect.Width) - 1 - 2
			xpos   = float32(clamp(int(x-constEffectRect.X), 1, int(w+1))) - 1
			perc   = xpos / w
			maxVal = m.ui.ConstEffectBar.GetMaxValue()
			minVal = m.ui.ConstEffectBar.GetMinValue()
			newVal = minVal + int(perc*float32(maxVal-minVal))
		)

		m.ui.ConstEffectBar.SetValue(newVal)
		m.dev.SetConstantEffect(float32(newVal) / 10)
	}
}

func clamp(v, _min, _max int) int {
	return min(_max, max(_min, v))
}

func isInRect(r lbl.Rect, x, y uint16) bool {
	return x >= r.X &&
		x < r.X+r.Width &&
		y >= r.Y &&
		y < r.Y+r.Height
}
