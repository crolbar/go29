package button

import lb "github.com/crolbar/lipbalm"

type Button struct {
	title   string
	pressed bool
	margin  int
}

type opts func(b *Button)

func WithMargin(m int) opts {
	return func(b *Button) {
		b.margin = m
	}
}

func NewButton(title string, opts ...opts) *Button {
	b := &Button{
		title:   title,
		pressed: false,
		margin:  1,
	}

	for _, o := range opts {
		o(b)
	}

	return b
}

func (b *Button) Toggle() {
	b.pressed = !b.pressed
}

func (b *Button) Release() {
	b.pressed = false
}

func (b *Button) GetState() bool {
	return b.pressed
}

func (b Button) View() string {
	if b.pressed {
		return lb.BorderN(
			lb.SetColor(lb.ColorBg(57),
				lb.Margin(b.margin, b.title),
			),
		)
	}

	return lb.BorderN(lb.Margin(b.margin, b.title))
}
