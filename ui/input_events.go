package ui

import (
	"go29/device"
	ec "go29/event_codes"
)

func (u *Ui) HandleInputEvent(evt device.InputEvent) {
	var (
		_type = evt.Type
		code  = int(evt.Code)
		value = int(evt.Value)
	)

	if _type == ec.EV_KEY {
		u.reqRender[Buttons] = true
		switch value {
		case 1:
			u.Buttons[code].Toggle()
		case 0:
			u.Buttons[code].Release()
		}
		return
	}

	switch code {
	case ec.ABS_WHEEL:
		if evt.Value < 32767 {
			u.WheelLeftBar.SetValue(32767 - value)
			u.WheelRightBar.SetValue(0)
		} else {
			u.WheelLeftBar.SetValue(0)
			u.WheelRightBar.SetValue(value - 32767)
		}
	case ec.ABS_THROTTLE:
		u.ThrottleBar.SetValue(255 - value)
	case ec.ABS_BREAK:
		u.BreakBar.SetValue(255 - value)
	case ec.ABS_CLUTCH:
		u.ClutchBar.SetValue(255 - value)
	case ec.ABS_DPADX, ec.ABS_DPADY:
		u.handleDpadInput(value, code)
	}
}

func (u *Ui) handleDpadInput(value, code int) {
	u.reqRender[Dpad] = true

	if value == 0 {
		u.Dpad[code][-1].Release()
		u.Dpad[code][1].Release()
		return
	}

	u.Dpad[code][value].Toggle()
}
