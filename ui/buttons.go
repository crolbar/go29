package ui

import (
	ec "go29/event_codes"
	"go29/ui/button"
	b "go29/ui/button"
	"sort"

	"github.com/crolbar/lipbalm"
)

var buttonsMap = map[int]*b.Button{
	ec.BTN_TRIANGLE: b.NewButton("T"),
	ec.BTN_SQUARE:   b.NewButton("S"),
	ec.BTN_CIRCLE:   b.NewButton("C"),
	ec.BTN_X:        b.NewButton("X"),

	ec.BTN_R3: b.NewButton("R3"),
	ec.BTN_L3: b.NewButton("L3"),
	ec.BTN_R2: b.NewButton("R2"),
	ec.BTN_L2: b.NewButton("L2"),

	ec.BTN_RIGHT_PADDLE: b.NewButton("RP"),
	ec.BTN_LEFT_PADDLE:  b.NewButton("LP"),

	ec.BTN_ROTARY_RIGHT:  b.NewButton("RL"),
	ec.BTN_ROTARY_LEFT:   b.NewButton("RR"),
	ec.BTN_ROTARY_BUTTON: b.NewButton("RB"),

	ec.BTN_PLUS:  b.NewButton("Pl"),
	ec.BTN_MINUS: b.NewButton("Mi"),

	ec.BTN_SHIFTER_FIRST:   b.NewButton("1t"),
	ec.BTN_SHIFTER_SECOND:  b.NewButton("2d"),
	ec.BTN_SHIFTER_THIRD:   b.NewButton("3d"),
	ec.BTN_SHIFTER_FOURTH:  b.NewButton("4h"),
	ec.BTN_SHIFTER_FIFTH:   b.NewButton("5h"),
	ec.BTN_SHIFTER_SIXTH:   b.NewButton("6h"),
	ec.BTN_SHIFTER_REVERSE: b.NewButton("R"),

	ec.BTN_OPTIONS: b.NewButton("OP"),
	ec.BTN_PS:      b.NewButton("PS"),
	ec.BTN_SHARE:   b.NewButton("SH"),
}

var buttonMapKeys = make([]int, 0, len(buttonsMap))

func setButtonMapKeys() {
	for k := range buttonsMap {
		buttonMapKeys = append(buttonMapKeys, k)
	}
	sort.Ints(buttonMapKeys)
}

var dpadMap = map[int]map[int]*b.Button{
	ec.ABS_DPADY: {
		ec.DPAD_UP:   b.NewButton("DU"), // 17: -1
		ec.DPAD_DOWN: b.NewButton("DD"), // 17: 1
	},
	ec.ABS_DPADX: {
		ec.DPAD_LEFT:  b.NewButton("DL"), // 16: -1
		ec.DPAD_RIGHT: b.NewButton("DR"), // 16: 1
	},
}

var dpadMapKeys = make([]int, 0, len(dpadMap)*2)

func setDpadMapKeys() {
	for _, o := range dpadMap {
		for k := range o {
			dpadMapKeys = append(dpadMapKeys, k)
		}
	}
	sort.Ints(dpadMapKeys)
}

func (u *Ui) renderButtons() string {
	buttons := ""

	j := 0

	for i := 0; i < 3; i++ {
		lineMaxJ := j + 10
		line := ""

		for ; j < lineMaxJ && j < len(buttonMapKeys); j++ {
			line = lipbalm.JoinHorizontal(lipbalm.Left,
				line,
				(u.Buttons[buttonMapKeys[j]]).View(),
			)
		}

		if i == 2 {
			var dpad string
			if !u.reqRender[Dpad] && u.havePreRender(Dpad) {
				dpad = u.preRenders[Dpad]
			} else {
				dpad = renderDpad(u.Dpad)
				u.preRenders[Dpad] = dpad
				u.reqRender[Dpad] = false
			}

			line = lipbalm.JoinHorizontal(lipbalm.Left,
				line,
				dpad,
			)
		}

		if buttons == "" {
			buttons = line
		} else {
			buttons = lipbalm.JoinVertical(lipbalm.Left,
				buttons,
				line,
			)
		}
	}

	return buttons
}

func renderDpad(Dpad map[int]map[int]*button.Button) string {
	dpad := ""

	for di, db := range dpadMapKeys {
		dpad = lipbalm.JoinHorizontal(lipbalm.Left,
			dpad,
			((Dpad[iff((di&1) > 0, ec.ABS_DPADX, ec.ABS_DPADY)])[db]).View(),
		)
	}

	return dpad
}
