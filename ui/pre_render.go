package ui

import (
	lb "github.com/crolbar/lipbalm"
	lbl "github.com/crolbar/lipbalm/layout"
)

var l lbl.Layout = lbl.DefaultLayout()

func (u *Ui) PreRender() {
	var (
		wheelBar = u.preRenderWheelBar()

		rangeBar       = u.preRenderRangeBar()
		autoCenterBar  = u.preRenderAutoCenterBar()
		constEffectBar = u.preRenderConstEffectBar()
		bars           = lb.JoinVertical(lb.Left,
			rangeBar,
			autoCenterBar,
			constEffectBar)

		buttons     = u.preRenderButtons()
		clutchBar   = u.preRenderClutchBar()
		breakBar    = u.preRenderBreakBar()
		throttleBar = u.preRenderThrottleBar()
		pedals      = lb.JoinHorizontal(lb.Left,
			clutchBar,
			breakBar,
			throttleBar,
		)

		vs = l.Vercital().
			Constrains(
				lbl.NewConstrain(lbl.Length, uint16(lb.GetHeight(wheelBar))),
				lbl.NewConstrain(lbl.Length, uint16(lb.GetHeight(bars))),
				lbl.NewConstrain(lbl.Percent, 60),
			).Split(u.fb.Size())

		hs = l.Horizontal().
			Constrains(
				lbl.NewConstrain(lbl.Length, uint16(lb.GetWidth(pedals))),
				lbl.NewConstrain(lbl.Percent, 60),
			).Split(vs[2])
	)

	u.fb.Clear()

	u.fb.RenderString(wheelBar, vs[0], lb.Left, lb.Top)
	u.fb.RenderString(buttons, hs[1], lb.Left, lb.Bottom)
	u.fb.RenderString(pedals, hs[0], lb.Left, lb.Bottom)
	u.fb.RenderString(bars, vs[1], lb.Right, lb.Center)

	u.preRenders[Screen] = u.fb.View()
}

func (u Ui) havePreRender(elem UiElement) bool {
	return u.preRenders[elem] != ""
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
