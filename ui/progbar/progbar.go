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
	var (
		sb  strings.Builder
		end = iff(p.vertical, p.height, p.width)

		v        = float32(p.value)
		min_v    = float32(p.min_value)
		max_v    = float32(p.max_value)
		perc     = (v - min_v) / (max_v - min_v)
		progress = int(perc * float32(end))

		fullLine  = strings.Repeat("█", iff(p.vertical, p.width, progress))
		emptyline = strings.Repeat(" ", iff(p.vertical, p.width, end-progress))

		endIdx = p.height - 1

		getLineAt = func(i int) string {
			if p.vertical {
				if iff(p.reverse, end-1-i, i) < progress {
					return fullLine
				}

				return emptyline
			}

			if p.reverse {
				return emptyline + fullLine
			}

			return fullLine + emptyline
		}
	)

	for i := range end {
		sb.WriteString(getLineAt(i))

		if i == endIdx {
			break
		}

		sb.WriteByte('\n')
	}

	return lb.Border(lb.NormalBorder(
		lb.WithFgColor(uint8(iff(p.selected, 57, 15))),
		lb.WithTextTop(fmt.Sprintf("%s(%d)", p.title, p.value), lb.Center),
	), lb.SetColor(lb.Color(57),
		sb.String()),
		false,
		p.noRightBorder,
		false,
		p.noLeftBorder)
}

func iff[T int | bool | string | uint8](b bool, f, s T) T {
	if b {
		return f
	}
	return s
}
