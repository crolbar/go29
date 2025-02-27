package button

import lb "github.com/crolbar/lipbalm"

type Button struct {
	title   string
	pressed bool
}

func NewButton(title string) *Button {
	return &Button{
		title:   title,
		pressed: false,
	}
}

func (b *Button) Toggle() {
	b.pressed = !b.pressed
}

func (b *Button) Release() {
	b.pressed = false
}

func (b Button) View() string {
	if b.pressed {
		return lb.BorderN(
			lb.SetColor(lb.ColorBg(57),
				lb.Margin(1, b.title),
			),
		)
	}

	return lb.BorderN(lb.Margin(1, b.title))
}
