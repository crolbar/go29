package ui

import (
	ec "go29/event_codes"
	"go29/ui/button"
	b "go29/ui/button"
	"sort"

	"github.com/crolbar/lipbalm"
)

var buttonsMap = map[int]*b.Button{
	ec.BUTTON_OPTIONS: b.NewButton("OP"),
	ec.BUTTON_PS:      b.NewButton("PS"),
	ec.BUTTON_SHARE:   b.NewButton("SH"),

	ec.BUTTON_R3:      b.NewButton("R3"),
	ec.BUTTON_L3:      b.NewButton("L3"),
	ec.BUTTON_R2:      b.NewButton("R2"),
	ec.BUTTON_L2:      b.NewButton("L2"),
	ec.BUTTON_RPaddle: b.NewButton("RP"),
	ec.BUTTON_LPaddle: b.NewButton("LP"),

	ec.BUTTON_RotaryL: b.NewButton("RL"),
	ec.BUTTON_RotaryR: b.NewButton("RR"),
	ec.BUTTON_Rotary:  b.NewButton("RB"),

	ec.BUTTON_T: b.NewButton("T"),
	ec.BUTTON_S: b.NewButton("S"),
	ec.BUTTON_C: b.NewButton("C"),
	ec.BUTTON_X: b.NewButton("X"),

	ec.BUTTON_Plus:  b.NewButton("Pl"),
	ec.BUTTON_Minus: b.NewButton("Mi"),

	ec.SHIFTER_FIRST:   b.NewButton("1t"),
	ec.SHIFTER_SECOND:  b.NewButton("2d"),
	ec.SHIFTER_THIRD:   b.NewButton("3d"),
	ec.SHIFTER_FOURTH:  b.NewButton("4h"),
	ec.SHIFTER_FIFTH:   b.NewButton("5h"),
	ec.SHIFTER_SIXTH:   b.NewButton("6h"),
	ec.SHIFTER_REVERSE: b.NewButton("R"),
}

var buttonMapKeys = make([]int, 0, len(buttonsMap))

func setButtonMapKeys() {
	for k := range buttonsMap {
		buttonMapKeys = append(buttonMapKeys, k)
	}
	sort.Ints(buttonMapKeys)
}

var dpadMap = map[int]map[int]*b.Button{
	ec.ABS_HAT0Y: {
		ec.DPAD_UP:   b.NewButton("DU"), // 17: -1
		ec.DPAD_DOWN: b.NewButton("DD"), // 17: 1
	},
	ec.ABS_HAT0X: {
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
			((Dpad[iff((di&1) > 0, ec.ABS_HAT0X, ec.ABS_HAT0Y)])[db]).View(),
		)
	}

	return dpad
}
