package ui

import (
	"go29/device"
	"go29/event_codes"
)

func (u *Ui) HandleInputEvent(evt device.InputEvent) {
	var (
		value = int(evt.Value)
		code  = int(evt.Code)
	)

	switch code {
	case event_codes.ABS_X:
		if evt.Value < 32767 {
			u.WheelLeftBar.SetValue(32767 - value)
			u.WheelRightBar.SetValue(0)
		} else {
			u.WheelLeftBar.SetValue(0)
			u.WheelRightBar.SetValue(value - 32767)
		}
	case event_codes.ABS_Z:
		u.ThrottleBar.SetValue(255 - value)
	case event_codes.ABS_RZ:
		u.BreakBar.SetValue(255 - value)
	case event_codes.ABS_Y:
		u.ClutchBar.SetValue(255 - value)
	case event_codes.ABS_RY:
		u.reqRender[Buttons] = true
		u.Buttons[value].Toggle()
	case event_codes.ABS_HAT0X, event_codes.ABS_HAT0Y:
		u.reqRender[Dpad] = true

		if value == 0 {
			u.Dpad[code][-1].Release()
			u.Dpad[code][1].Release()
			break
		}

		u.Dpad[code][value].Toggle()
	}
}
