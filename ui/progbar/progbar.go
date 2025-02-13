package progbar

import (
	"fmt"
	"strings"

	lb "github.com/crolbar/lipbalm"
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

func (p *ProgBar) SetVertical(b bool) {
	p.vertical = b
	tmp := p.height
	p.height = p.width / 3
	p.width = tmp * 3
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
			barStr = lb.JoinHorizontal(lb.Left,
				barStr,
				tmp,
			)
			continue
		}

		barStr = fmt.Sprintf("%s\n%s", barStr, tmp)
	}

	borderColor := lb.Color(uint8(iff(p.selected, 57, 15)))

	title := lb.Border(lb.BorderType{
		Left:        "│",
		Right:       "│",
		Bottom:      "─",
		BottomRight: "┤",
		BottomLeft:  "├",
		ColorFg:     borderColor,
	}, lb.ExpandHorizontal(
		p.width, lb.Center,
		fmt.Sprintf("%s(%d)", p.title, p.value),
	),
		true,
		p.noRightBorder,
		false,
		p.noLeftBorder)

	bar := lb.Border(lb.NormalBorder(borderColor),
		lb.SetColor(lb.Color(57),
			lb.ExpandHorizontal(p.width, lb.Left,
				lb.ExpandVertical(p.height, lb.Left,
					barStr,
				),
			),
		),
		true,
		p.noRightBorder,
		false,
		p.noLeftBorder)

	return lb.JoinVertical(lb.Center,
		title,
		bar,
	)
}

func iff[T int | bool | string | uint8](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
