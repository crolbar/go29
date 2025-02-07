package ui

import "go29/device"

func (u *Ui) SelectRangeBar() {
	u.selectedBar = WheelRange
}

var WheelRangeMarks = []int{40, 90, 180, 270, 360, 450, 540, 720, 900}

func (u *Ui) HandleSelectedBarLeft(dev *device.Device) {
	switch u.selectedBar {
	case WheelRange:
		dev.SetRange(u.getSelecedBarPrevMark())
		u.WheelRange.SetValue(dev.GetRange())
	}
}

func (u *Ui) HandleSelectedBarRight(dev *device.Device) {
	switch u.selectedBar {
	case WheelRange:
		dev.SetRange(u.getSelecedBarNextMark())
		// unnecessary ?
		u.WheelRange.SetValue(dev.GetRange())
	}
}

func (u *Ui) getSelecedBarNextMark() int {
	switch u.selectedBar {
	case WheelRange:
		v := u.WheelRange.GetValue()

		for _, m := range WheelRangeMarks {
			if v < m {
				return m
			}
		}

		return WheelRangeMarks[len(WheelRangeMarks)-1]
	}

	return 0
}

func (u *Ui) getSelecedBarPrevMark() int {
	switch u.selectedBar {
	case WheelRange:
		v := u.WheelRange.GetValue()

		len := len(WheelRangeMarks) - 1
		for i := range WheelRangeMarks {
			m := WheelRangeMarks[len-i]
			if v > m {
				return m
			}
		}

		return WheelRangeMarks[0]
	}

	return 0
}
