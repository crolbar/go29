package ui

import (
	lb "github.com/crolbar/lipbalm"
	lbl "github.com/crolbar/lipbalm/layout"
)

var l lbl.Layout = lbl.DefaultLayout()

func (u *Ui) PreRender() {
	var (
		wheelBar = u.preRenderWheelBar()

		virtDevButton = u.preRenderVirtDevButton()

		rangeBar       = u.preRenderRangeBar()
		autoCenterBar  = u.preRenderAutoCenterBar()
		constEffectBar = u.preRenderConstEffectBar()

		buttons     = u.preRenderButtons()
		clutchBar   = u.preRenderClutchBar()
		breakBar    = u.preRenderBreakBar()
		throttleBar = u.preRenderThrottleBar()
		pedals      = lb.JoinHorizontal(lb.Left,
			clutchBar,
			breakBar,
			throttleBar,
		)

		rangeBarHeight = uint16(lb.GetHeight(rangeBar))
		rangeBarWidth  = uint16(lb.GetWidth(rangeBar))

		vs = l.Vercital().
			Constrains(
				lbl.NewConstrain(lbl.Length, uint16(lb.GetHeight(wheelBar))),
				lbl.NewConstrain(lbl.Length, rangeBarHeight*3),
				lbl.NewConstrain(lbl.Percent, 60),
			).Split(u.fb.Size())

		barss = l.Vercital().
			Constrains(
				lbl.NewConstrain(lbl.Length, rangeBarHeight),
				lbl.NewConstrain(lbl.Length, rangeBarHeight),
				lbl.NewConstrain(lbl.Length, rangeBarHeight),
			).Split(vs[1])

		rangeBarRect = l.Horizontal().
				Constrains(lbl.NewConstrain(lbl.Percent, 100), lbl.NewConstrain(lbl.Length, rangeBarWidth)).
				Split(barss[0])[1]
		autoCenterBarRect = l.Horizontal().
					Constrains(lbl.NewConstrain(lbl.Percent, 100), lbl.NewConstrain(lbl.Length, rangeBarWidth)).
					Split(barss[1])[1]
		constEffectBarRect = l.Horizontal().
					Constrains(lbl.NewConstrain(lbl.Percent, 100), lbl.NewConstrain(lbl.Length, rangeBarWidth)).
					Split(barss[2])[1]

		hs = l.Horizontal().
			Constrains(
				lbl.NewConstrain(lbl.Length, uint16(lb.GetWidth(pedals))),
				lbl.NewConstrain(lbl.Percent, 60),
			).Split(vs[2])

		ths = l.Horizontal().
			Constrains(
				lbl.NewConstrain(lbl.Length, uint16(lb.GetWidth(wheelBar))),
				lbl.NewConstrain(lbl.Percent, 100),
			).Split(vs[0])
	)

	u.fb.Clear()

	u.fb.RenderString(wheelBar, ths[0], lb.Left, lb.Top)
	u.fb.RenderString(virtDevButton, ths[1], lb.Right, lb.Top)
	u.fb.RenderString(buttons, hs[1], lb.Left, lb.Bottom)
	u.fb.RenderString(pedals, hs[0], lb.Left, lb.Bottom)

	u.fb.RenderString(rangeBar, rangeBarRect, lb.Right, lb.Center)
	u.fb.RenderString(autoCenterBar, autoCenterBarRect, lb.Right, lb.Center)
	u.fb.RenderString(constEffectBar, constEffectBarRect, lb.Right, lb.Center)

	u.Rects[RangeBar] = rangeBarRect
	u.Rects[AutoCenterBar] = autoCenterBarRect
	u.Rects[ConstEffectBar] = constEffectBarRect

	u.preRenders[Screen] = u.fb.View()
}

func (u Ui) havePreRender(elem UiElement) bool {
	return u.preRenders[elem] != ""
}

func (u *Ui) preRenderVirtDevButton() string {
	val := 0
	if !u.VirtDevButton.GetState() {
		val = 1
	}

	if u.prevValues[virtDevButton] == val &&
		u.havePreRender(virtDevButton) {
		return u.preRenders[virtDevButton]
	}

	btn := u.VirtDevButton.View()
	u.prevValues[virtDevButton] = val
	u.preRenders[virtDevButton] = btn

	return btn
}

func (u *Ui) preRenderButtons() string {
	if !u.reqRender[Buttons] && !u.reqRender[Dpad] &&
		u.havePreRender(Buttons) {
		return u.preRenders[Buttons]
	}

	buttons := u.renderButtons()
	u.reqRender[Buttons] = false
	u.preRenders[Buttons] = buttons

	return buttons
}

func (u *Ui) preRenderThrottleBar() string {
	val := u.ThrottleBar.GetValue()

	if u.prevValues[ThrottleBar] == val &&
		u.havePreRender(ThrottleBar) {
		return u.preRenders[ThrottleBar]
	}

	throttleBar := u.ThrottleBar.View()

	u.preRenders[ThrottleBar] = throttleBar
	u.prevValues[ThrottleBar] = val

	return throttleBar
}

func (u *Ui) preRenderBreakBar() string {
	val := u.BreakBar.GetValue()

	if u.prevValues[BreakBar] == val &&
		u.havePreRender(BreakBar) {
		return u.preRenders[BreakBar]
	}

	breakBar := u.BreakBar.View()

	u.preRenders[BreakBar] = breakBar
	u.prevValues[BreakBar] = val

	return breakBar
}

func (u *Ui) preRenderClutchBar() string {
	val := u.ClutchBar.GetValue()

	if u.prevValues[ClutchBar] == val &&
		u.havePreRender(ClutchBar) {
		return u.preRenders[ClutchBar]
	}

	clutchBar := u.ClutchBar.View()

	u.preRenders[ClutchBar] = clutchBar
	u.prevValues[ClutchBar] = val

	return clutchBar
}

func (u *Ui) preRenderAutoCenterBar() string {
	val := u.AutoCenterBar.GetValue()

	if !u.reqRender[AutoCenterBar] &&
		u.prevValues[AutoCenterBar] == val &&
		u.havePreRender(AutoCenterBar) {
		return u.preRenders[AutoCenterBar]
	}

	autoCenterBar := u.AutoCenterBar.View()

	u.preRenders[AutoCenterBar] = autoCenterBar
	u.prevValues[AutoCenterBar] = val
	u.reqRender[AutoCenterBar] = false

	return autoCenterBar
}

func (u *Ui) preRenderConstEffectBar() string {
	val := u.ConstEffectBar.GetValue()

	if !u.reqRender[ConstEffectBar] &&
		u.prevValues[ConstEffectBar] == val &&
		u.havePreRender(ConstEffectBar) {
		return u.preRenders[ConstEffectBar]
	}

	constEffectBar := u.ConstEffectBar.View()

	u.preRenders[ConstEffectBar] = constEffectBar
	u.prevValues[ConstEffectBar] = val
	u.reqRender[ConstEffectBar] = false

	return constEffectBar
}

func (u *Ui) preRenderRangeBar() string {
	val := u.RangeBar.GetValue()

	if !u.reqRender[RangeBar] &&
		u.prevValues[RangeBar] == val &&
		u.havePreRender(RangeBar) {
		return u.preRenders[RangeBar]
	}

	rangeBar := u.RangeBar.View()

	u.preRenders[RangeBar] = rangeBar
	u.prevValues[RangeBar] = val
	u.reqRender[RangeBar] = false

	return rangeBar
}

func (u *Ui) preRenderWheelBar() string {
	leftVal := u.WheelLeftBar.GetValue()
	rightVal := u.WheelRightBar.GetValue()

	if leftVal == u.prevValues[WheelBarLeft] &&
		rightVal == u.prevValues[WheelBarRight] &&
		u.havePreRender(WheelBar) {
		return u.preRenders[WheelBar]
	}

	wheelBar := lb.JoinHorizontal(lb.Left,
		u.WheelLeftBar.View(),
		u.WheelRightBar.View(),
	)

	u.preRenders[WheelBar] = wheelBar
	u.prevValues[WheelBarLeft] = leftVal
	u.prevValues[WheelBarRight] = rightVal

	return wheelBar
}
