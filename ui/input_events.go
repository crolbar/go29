package ui

import (
	"go29/device"
	ec "go29/event_codes"
)

func (u *Ui) HandleInputEvent(
	evt device.InputEvent,
	dev *device.Device,
) {
	var (
		value = int(evt.Value)
		code  = int(evt.Code)
	)

	switch code {
	case ec.ABS_X:
		if evt.Value < 32767 {
			u.WheelLeftBar.SetValue(32767 - value)
			u.WheelRightBar.SetValue(0)
		} else {
			u.WheelLeftBar.SetValue(0)
			u.WheelRightBar.SetValue(value - 32767)
		}
	case ec.ABS_Z:
		u.ThrottleBar.SetValue(255 - value)
	case ec.ABS_RZ:
		u.BreakBar.SetValue(255 - value)
	case ec.ABS_Y:
		u.ClutchBar.SetValue(255 - value)
	case ec.ABS_RY:
		u.reqRender[Buttons] = true
		u.Buttons[value].Toggle()
	case ec.ABS_HAT0X, ec.ABS_HAT0Y:
		u.handleDpadInput(value, code, dev)
	}
}

func (u *Ui) handleDpadInput(
	value, code int,
	dev *device.Device,
) {
	u.reqRender[Dpad] = true

	if value == 0 {
		u.Dpad[code][-1].Release()
		u.Dpad[code][1].Release()
		return
	}

	switch code {
	case ec.ABS_HAT0Y:
		switch value {
		case ec.DPAD_UP:
			u.SelectPrevBar()
		case ec.DPAD_DOWN:
			u.SelectNextBar()
		}
	case ec.ABS_HAT0X:
		switch value {
		case ec.DPAD_RIGHT:
			u.HandleSelectedBarRight(dev)
		case ec.DPAD_LEFT:
			u.HandleSelectedBarLeft(dev)
		}
	}

	u.Dpad[code][value].Toggle()
}
