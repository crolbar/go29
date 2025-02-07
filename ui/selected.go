package ui

import "go29/device"

var WheelRangeMarks = []int{40, 90, 180, 270, 360, 450, 540, 720, 900}
var AutoCenterMarks = []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

func (u *Ui) HandleSelectedBarLeft(dev *device.Device) {
	switch u.selectedBar {
	case WheelRange:
		dev.SetRange(u.getSelecedBarPrevMark())
		u.WheelRange.SetValue(dev.GetRange())
	case AutoCenter:
		p := u.getSelecedBarPrevMark()
		dev.SetAutocenter(p)
		u.AutoCenter.SetValue(p)
	}
}

func (u *Ui) HandleSelectedBarRight(dev *device.Device) {
	switch u.selectedBar {
	case WheelRange:
		dev.SetRange(u.getSelecedBarNextMark())
		// unnecessary ?
		u.WheelRange.SetValue(dev.GetRange())

	case AutoCenter:
		p := u.getSelecedBarNextMark()
		dev.SetAutocenter(p)
		u.AutoCenter.SetValue(p)
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
	case AutoCenter:
		v := u.AutoCenter.GetValue()

		for _, m := range AutoCenterMarks {
			if v < m {
				return m
			}
		}

		return AutoCenterMarks[len(AutoCenterMarks)-1]
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
	case AutoCenter:
		v := u.AutoCenter.GetValue()

		len := len(AutoCenterMarks) - 1
		for i := range AutoCenterMarks {
			m := AutoCenterMarks[len-i]
			if v > m {
				return m
			}
		}

		return AutoCenterMarks[0]
	}

	return 0
}

func (u *Ui) SelectNextBar() {
	switch u.selectedBar {
	case WheelRange:
		u.selectedBar = AutoCenter
		u.WheelRange.DeSelect()
		u.AutoCenter.Select()
	case AutoCenter:
		u.selectedBar = WheelRange
		u.AutoCenter.DeSelect()
		u.WheelRange.Select()
	}
}

func (u *Ui) SelectPrevBar() {
	switch u.selectedBar {
	case WheelRange:
		u.selectedBar = AutoCenter
		u.WheelRange.DeSelect()
		u.AutoCenter.Select()
	case AutoCenter:
		u.selectedBar = WheelRange
		u.AutoCenter.DeSelect()
		u.WheelRange.Select()
	}
}
