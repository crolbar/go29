package ui

import "go29/device"

var WheelRangeMarks = []int{40, 90, 180, 270, 360, 450, 540, 720, 900}
var AutoCenterMarks = []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

func (u *Ui) HandleSelectedBarLeft(dev *device.Device) {
	switch u.selectedBar {
	case Range:
		dev.SetRange(u.getSelecedBarPrevMark())
		u.RangeBar.SetValue(dev.GetRange())
	case AutoCenter:
		p := u.getSelecedBarPrevMark()
		dev.SetAutocenter(p)
		u.AutoCenterBar.SetValue(p)
	case ConstEffect:
		v := max(u.ConstEffectBar.GetValue() - 1, -10)
		dev.SetConstantEffect(float32(v) / 10.0)
		u.ConstEffectBar.SetValue(v)
	}
}

func (u *Ui) HandleSelectedBarRight(dev *device.Device) {
	switch u.selectedBar {
	case Range:
		dev.SetRange(u.getSelecedBarNextMark())
		// unnecessary ?
		u.RangeBar.SetValue(dev.GetRange())

	case AutoCenter:
		p := u.getSelecedBarNextMark()
		dev.SetAutocenter(p)
		u.AutoCenterBar.SetValue(p)

	case ConstEffect:
		v := min(u.ConstEffectBar.GetValue() + 1, 10)
		dev.SetConstantEffect(float32(v) / 10.0)
		u.ConstEffectBar.SetValue(v)
	}
}

func (u *Ui) getSelecedBarNextMark() int {
	switch u.selectedBar {
	case Range:
		v := u.RangeBar.GetValue()

		for _, m := range WheelRangeMarks {
			if v < m {
				return m
			}
		}

		return WheelRangeMarks[len(WheelRangeMarks)-1]
	case AutoCenter:
		v := u.AutoCenterBar.GetValue()

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
	case Range:
		v := u.RangeBar.GetValue()

		len := len(WheelRangeMarks) - 1
		for i := range WheelRangeMarks {
			m := WheelRangeMarks[len-i]
			if v > m {
				return m
			}
		}

		return WheelRangeMarks[0]
	case AutoCenter:
		v := u.AutoCenterBar.GetValue()

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
	case Range:
		u.RangeBar.DeSelect()
		u.selectedBar = AutoCenter
		u.AutoCenterBar.Select()

		u.reqRender[RangeBar] = true
		u.reqRender[AutoCenterBar] = true
	case AutoCenter:
		u.AutoCenterBar.DeSelect()
		u.selectedBar = ConstEffect
		u.ConstEffectBar.Select()

		u.reqRender[AutoCenterBar] = true
		u.reqRender[ConstEffectBar] = true
	case ConstEffect:
		u.ConstEffectBar.DeSelect()
		u.selectedBar = Range
		u.RangeBar.Select()

		u.reqRender[RangeBar] = true
		u.reqRender[ConstEffectBar] = true
	}
}

func (u *Ui) SelectPrevBar() {
	switch u.selectedBar {
	case Range:
		u.RangeBar.DeSelect()
		u.selectedBar = ConstEffect
		u.ConstEffectBar.Select()

		u.reqRender[RangeBar] = true
		u.reqRender[ConstEffectBar] = true
	case AutoCenter:
		u.AutoCenterBar.DeSelect()
		u.selectedBar = Range
		u.RangeBar.Select()

		u.reqRender[AutoCenterBar] = true
		u.reqRender[RangeBar] = true
	case ConstEffect:
		u.ConstEffectBar.DeSelect()
		u.selectedBar = AutoCenter
		u.AutoCenterBar.Select()

		u.reqRender[AutoCenterBar] = true
		u.reqRender[ConstEffectBar] = true
	}
}
