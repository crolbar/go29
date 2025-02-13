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
	return lb.Border(
		lb.NormalBorder(),
		lb.SetColor(
			lb.ColorBg(iff(b.pressed, uint8(57), 0)),
			lb.Margin(1,
				b.title,
			),
		),
	)
}

func iff[T string | uint8](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
