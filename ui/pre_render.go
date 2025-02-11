package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/crolbar/lipbalm"
)

var s lipgloss.Style = lipgloss.NewStyle()

var screenStyle lipgloss.Style = s.PaddingTop(2).
	PaddingBottom(2).
	PaddingRight(5).
	PaddingLeft(5)

func (u *Ui) PreRender() {
	var (
		wheelBar = u.preRenderWheelBar()
	)

	var (
		rangeBar      = u.preRenderRangeBar()
		autoCenterBar = u.preRenderAutoCenterBar()
		sliderBars = s.MarginLeft(10).
				Render(
				lipbalm.JoinVertical(lipbalm.Left,
					rangeBar,
					autoCenterBar,
				),
			)
	)

	var (
		buttons     = u.preRenderButtons()
		clutchBar   = u.preRenderClutchBar()
		breakBar    = u.preRenderBreakBar()
		throttleBar = u.preRenderThrottleBar()
		pedals      = lipbalm.JoinHorizontal(lipbalm.Left,
			clutchBar,
			breakBar,
			throttleBar,
		)

		buttonsPedals = s.
				Height(u.height - lipgloss.Height(wheelBar) - 4).
				AlignVertical(lipgloss.Bottom).
				Render(
				lipbalm.JoinVertical(lipbalm.Left,
					buttons,
					pedals,
				),
			)
	)

	u.preRenders[Screen] = screenStyle.Render(
		lipbalm.JoinHorizontal(lipbalm.Top,
			lipbalm.JoinVertical(lipbalm.Left,
				wheelBar,
				buttonsPedals,
			),
			sliderBars,
		),
	)
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

	// add selected
	if u.prevValues[AutoCenterBar] == val &&
		u.havePreRender(AutoCenterBar) {
		return u.preRenders[AutoCenterBar]
	}

	autoCenterBar := u.AutoCenterBar.View()

	u.preRenders[AutoCenterBar] = autoCenterBar
	u.prevValues[AutoCenterBar] = val

	return autoCenterBar
}

func (u *Ui) preRenderRangeBar() string {
	val := u.RangeBar.GetValue()

	// add selected
	if u.prevValues[RangeBar] == val &&
		u.havePreRender(RangeBar) {
		return u.preRenders[RangeBar]
	}

	rangeBar := u.RangeBar.View()

	u.preRenders[RangeBar] = rangeBar
	u.prevValues[RangeBar] = val

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

	wheelBar := lipbalm.JoinHorizontal(lipbalm.Left,
		u.WheelLeftBar.View(),
		u.WheelRightBar.View(),
	)

	u.preRenders[WheelBar] = wheelBar
	u.prevValues[WheelBarLeft] = leftVal
	u.prevValues[WheelBarRight] = rightVal

	return wheelBar
}
