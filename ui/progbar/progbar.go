package progbar

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ProgBar struct {
	title string

	value     int
	min_value int
	max_value int

	height int
	width  int

	reverse       bool
	vertical      bool
	noLeftBorder  bool
	noRightBorder bool
	selected      bool
}

var s lipgloss.Style = lipgloss.NewStyle()

func WithDisabledRightBorder() Opts {
	return func(p *ProgBar) {
		p.noRightBorder = true
	}
}

func WithDisabledLeftBorder() Opts {
	return func(p *ProgBar) {
		p.noLeftBorder = true
	}
}

func WithMaxValue(max int) Opts {
	return func(p *ProgBar) {
		p.max_value = max
	}
}

func WithMinValue(min int) Opts {
	return func(p *ProgBar) {
		p.min_value = min
		if p.value < min {
			p.value = min
		}
	}
}

func WithVertical() Opts {
	return func(p *ProgBar) {
		p.vertical = true
	}
}

func WithReverse() Opts {
	return func(p *ProgBar) {
		p.reverse = true
	}
}

func WithSelected() Opts {
	return func(p *ProgBar) {
		p.selected = true
	}
}

func WithValue(value int) Opts {
	return func(p *ProgBar) {
		p.value = value
	}
}

type Opts func(*ProgBar)

func NewProgBar(
	title string,
	height int,
	width int,
	opts ...Opts,
) ProgBar {
	p := ProgBar{
		title:         title,
		value:         0,
		min_value:     0,
		max_value:     100,
		height:        height,
		width:         width,
		vertical:      false,
		noLeftBorder:  false,
		noRightBorder: false,
		selected:      false,
	}

	for _, opt := range opts {
		opt(&p)
	}

	return p
}

func (p *ProgBar) Select() {
	p.selected = true
}

func (p *ProgBar) DeSelect() {
	p.selected = false
}

func (p *ProgBar) SetTitle(title string) {
	p.title = title
}

func (p *ProgBar) SetValue(value int) {
	p.value = value
}

func (p *ProgBar) GetValue() int {
	return p.value
}

func (p ProgBar) View() string {
	var barbi strings.Builder

	end := iff(p.vertical, p.height, p.width)

	v := float32(p.value)
	min_v := float32(p.min_value)
	max_v := float32(p.max_value)
	perc := (v - min_v) / (max_v - min_v)
	progress := int(perc * float32(end))

	for i := 0; i < end; i++ {
		if iff(p.reverse, end-1-i, i) < progress {
			barbi.WriteString("█")
		} else {
			barbi.WriteString(" ")
		}

		if p.vertical && i != end-1 {
			barbi.WriteString("\n")
		}
	}

	tmp := barbi.String()

	barStr := barbi.String()

	for i := 1; i < iff(p.vertical, p.width, p.height); i++ {
		if p.vertical {
			barStr = lipgloss.JoinHorizontal(lipgloss.Center,
				barStr,
				tmp,
			)
			continue
		}

		barStr = fmt.Sprintf("%s\n%s", barStr, tmp)
	}

	borderColor := lipgloss.Color(iff(p.selected, "57", "15"))

	title := s.Border(lipgloss.Border{
		Left:        "│",
		Right:       "│",
		Bottom:      "─",
		BottomRight: "┤",
		BottomLeft:  "├",
	}).
		BorderTop(false).
		BorderRight(!p.noRightBorder).
		BorderLeft(!p.noLeftBorder).
		BorderForeground(borderColor).
		Width(p.width).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("%s(%d)", p.title, p.value))

	bar := s.Border(lipgloss.NormalBorder()).
		BorderTop(false).
		BorderRight(!p.noRightBorder).
		BorderLeft(!p.noLeftBorder).
		BorderForeground(borderColor).
		Foreground(lipgloss.Color("57")).
		Width(p.width).
		Height(p.height).
		Render(barStr)

	return lipgloss.JoinVertical(lipgloss.Center,
		title,
		bar,
	)
}

func iff[T int | bool | string](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
