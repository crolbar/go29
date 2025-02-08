package button

import "github.com/charmbracelet/lipgloss"

type Button struct {
	title   string
	pressed bool
}

var s lipgloss.Style = lipgloss.NewStyle()

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
	return s.Border(lipgloss.NormalBorder()).
		Background(lipgloss.Color(iff(b.pressed, "57", ""))).
		Padding(1).
		Render(b.title)
}

func iff[T string](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
