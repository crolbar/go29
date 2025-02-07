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
}

var s lipgloss.Style = lipgloss.NewStyle()

func NewProgBar(
	title string,
	height int,
	width int,
) ProgBar {
	return ProgBar{
		title:         title,
		value:         0,
		min_value:     0,
		max_value:     100,
		height:        height,
		width:         width,
		vertical:      false,
		noLeftBorder:  false,
		noRightBorder: false,
	}
}

func (p *ProgBar) Reverse(b bool) {
	p.reverse = b
}

func (p *ProgBar) DisableRightBorder(b bool) {
	p.noRightBorder = b
}

func (p *ProgBar) DisableLeftBorder(b bool) {
	p.noLeftBorder = b
}

func (p *ProgBar) SetMaxValue(max int) {
	p.max_value = max
}

func (p *ProgBar) SetMinValue(min int) {
	p.min_value = min
	if p.value < min {
		p.value = min
	}
}

func (p *ProgBar) SetVertical(vertical bool) {
	p.vertical = vertical
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
	perc := (v-min_v)/(max_v-min_v)
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
		Width(p.width).
		Align(lipgloss.Center).
		Render(fmt.Sprintf("%s(%d)", p.title, p.value))

	bar := s.Border(lipgloss.NormalBorder()).
		BorderTop(false).
		BorderRight(!p.noRightBorder).
		BorderLeft(!p.noLeftBorder).
		Foreground(lipgloss.Color("57")).
		Width(p.width).
		Height(p.height).
		Render(barStr)

	return lipgloss.JoinVertical(lipgloss.Center,
		title,
		bar,
	)
}

func iff[T int | bool](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
