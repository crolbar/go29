package virtDev

import (
	"errors"
	ec "go29/event_codes"
	"os"
	"path/filepath"
	"strings"
)

type wheelKey int
type key uint8

type kbKey struct {
	key      key
	click    bool
	modifier bool
}

var fromMap = map[string]wheelKey{
	"BTN_X":               ec.BTN_X,
	"BTN_CIRCLE":          ec.BTN_CIRCLE,
	"BTN_SQUARE":          ec.BTN_SQUARE,
	"BTN_TRIANGLE":        ec.BTN_TRIANGLE,
	"BTN_R2":              ec.BTN_R2,
	"BTN_R3":              ec.BTN_R3,
	"BTN_L2":              ec.BTN_L2,
	"BTN_L3":              ec.BTN_L3,
	"BTN_ROTARY_BUTTON":   ec.BTN_ROTARY_BUTTON,
	"BTN_ROTARY_RIGHT":    ec.BTN_ROTARY_RIGHT,
	"BTN_ROTARY_LEFT":     ec.BTN_ROTARY_LEFT,
	"BTN_SHARE":           ec.BTN_SHARE,
	"BTN_OPTIONS":         ec.BTN_OPTIONS,
	"BTN_PS":              ec.BTN_PS,
	"BTN_MINUS":           ec.BTN_MINUS,
	"BTN_PLUS":            ec.BTN_PLUS,
	"BTN_RIGHT_PADDLE":    ec.BTN_RIGHT_PADDLE,
	"BTN_LEFT_PADDLE":     ec.BTN_LEFT_PADDLE,
	"BTN_SHIFTER_FIRST":   ec.BTN_SHIFTER_FIRST,
	"BTN_SHIFTER_SECOND":  ec.BTN_SHIFTER_SECOND,
	"BTN_SHIFTER_THIRD":   ec.BTN_SHIFTER_THIRD,
	"BTN_SHIFTER_FOURTH":  ec.BTN_SHIFTER_FOURTH,
	"BTN_SHIFTER_FIFT":    ec.BTN_SHIFTER_FIFTH,
	"BTN_SHIFTER_SIXTH":   ec.BTN_SHIFTER_SIXTH,
	"BTN_SHIFTER_REVERSE": ec.BTN_SHIFTER_REVERSE,
	"ABS_WHEEL":           ec.ABS_WHEEL,
	"ABS_CLUTCH":          ec.ABS_CLUTCH,
	"ABS_THROTTLE":        ec.ABS_THROTTLE,
	"ABS_BREAK":           ec.ABS_BREAK,
	"ABS_DPADX":           ec.ABS_DPADX,
	"ABS_DPADY":           ec.ABS_DPADY,
}

var toMap = map[string]key{
	"KEY_ESC":              ec.KEY_ESC,
	"KEY_1":                ec.KEY_1,
	"KEY_2":                ec.KEY_2,
	"KEY_3":                ec.KEY_3,
	"KEY_4":                ec.KEY_4,
	"KEY_5":                ec.KEY_5,
	"KEY_6":                ec.KEY_6,
	"KEY_7":                ec.KEY_7,
	"KEY_8":                ec.KEY_8,
	"KEY_9":                ec.KEY_9,
	"KEY_0":                ec.KEY_0,
	"KEY_MINUS":            ec.KEY_MINUS,
	"KEY_EQUAL":            ec.KEY_EQUAL,
	"KEY_BACKSPACE":        ec.KEY_BACKSPACE,
	"KEY_TAB":              ec.KEY_TAB,
	"KEY_Q":                ec.KEY_Q,
	"KEY_W":                ec.KEY_W,
	"KEY_E":                ec.KEY_E,
	"KEY_R":                ec.KEY_R,
	"KEY_T":                ec.KEY_T,
	"KEY_Y":                ec.KEY_Y,
	"KEY_U":                ec.KEY_U,
	"KEY_I":                ec.KEY_I,
	"KEY_O":                ec.KEY_O,
	"KEY_P":                ec.KEY_P,
	"KEY_LEFTBRACE":        ec.KEY_LEFTBRACE,
	"KEY_RIGHTBRACE":       ec.KEY_RIGHTBRACE,
	"KEY_ENTER":            ec.KEY_ENTER,
	"KEY_LEFTCTRL":         ec.KEY_LEFTCTRL,
	"KEY_A":                ec.KEY_A,
	"KEY_S":                ec.KEY_S,
	"KEY_D":                ec.KEY_D,
	"KEY_F":                ec.KEY_F,
	"KEY_G":                ec.KEY_G,
	"KEY_H":                ec.KEY_H,
	"KEY_J":                ec.KEY_J,
	"KEY_K":                ec.KEY_K,
	"KEY_L":                ec.KEY_L,
	"KEY_SEMICOLON":        ec.KEY_SEMICOLON,
	"KEY_APOSTROPHE":       ec.KEY_APOSTROPHE,
	"KEY_GRAVE":            ec.KEY_GRAVE,
	"KEY_LEFTSHIFT":        ec.KEY_LEFTSHIFT,
	"KEY_BACKSLASH":        ec.KEY_BACKSLASH,
	"KEY_Z":                ec.KEY_Z,
	"KEY_X":                ec.KEY_X,
	"KEY_C":                ec.KEY_C,
	"KEY_V":                ec.KEY_V,
	"KEY_B":                ec.KEY_B,
	"KEY_N":                ec.KEY_N,
	"KEY_M":                ec.KEY_M,
	"KEY_COMMA":            ec.KEY_COMMA,
	"KEY_DOT":              ec.KEY_DOT,
	"KEY_SLASH":            ec.KEY_SLASH,
	"KEY_RIGHTSHIFT":       ec.KEY_RIGHTSHIFT,
	"KEY_KPASTERISK":       ec.KEY_KPASTERISK,
	"KEY_LEFTALT":          ec.KEY_LEFTALT,
	"KEY_SPACE":            ec.KEY_SPACE,
	"KEY_CAPSLOCK":         ec.KEY_CAPSLOCK,
	"KEY_F1":               ec.KEY_F1,
	"KEY_F2":               ec.KEY_F2,
	"KEY_F3":               ec.KEY_F3,
	"KEY_F4":               ec.KEY_F4,
	"KEY_F5":               ec.KEY_F5,
	"KEY_F6":               ec.KEY_F6,
	"KEY_F7":               ec.KEY_F7,
	"KEY_F8":               ec.KEY_F8,
	"KEY_F9":               ec.KEY_F9,
	"KEY_F10":              ec.KEY_F10,
	"KEY_NUMLOCK":          ec.KEY_NUMLOCK,
	"KEY_SCROLLLOCK":       ec.KEY_SCROLLLOCK,
	"KEY_KP7":              ec.KEY_KP7,
	"KEY_KP8":              ec.KEY_KP8,
	"KEY_KP9":              ec.KEY_KP9,
	"KEY_KPMINUS":          ec.KEY_KPMINUS,
	"KEY_KP4":              ec.KEY_KP4,
	"KEY_KP5":              ec.KEY_KP5,
	"KEY_KP6":              ec.KEY_KP6,
	"KEY_KPPLUS":           ec.KEY_KPPLUS,
	"KEY_KP1":              ec.KEY_KP1,
	"KEY_KP2":              ec.KEY_KP2,
	"KEY_KP3":              ec.KEY_KP3,
	"KEY_KP0":              ec.KEY_KP0,
	"KEY_KPDOT":            ec.KEY_KPDOT,
	"KEY_ZENKAKUHANKAKU":   ec.KEY_ZENKAKUHANKAKU,
	"KEY_102ND":            ec.KEY_102ND,
	"KEY_F11":              ec.KEY_F11,
	"KEY_F12":              ec.KEY_F12,
	"KEY_RO":               ec.KEY_RO,
	"KEY_KATAKANA":         ec.KEY_KATAKANA,
	"KEY_HIRAGANA":         ec.KEY_HIRAGANA,
	"KEY_HENKAN":           ec.KEY_HENKAN,
	"KEY_KATAKANAHIRAGANA": ec.KEY_KATAKANAHIRAGANA,
	"KEY_MUHENKAN":         ec.KEY_MUHENKAN,
	"KEY_KPJPCOMMA":        ec.KEY_KPJPCOMMA,
	"KEY_KPENTER":          ec.KEY_KPENTER,
	"KEY_RIGHTCTRL":        ec.KEY_RIGHTCTRL,
	"KEY_KPSLASH":          ec.KEY_KPSLASH,
	"KEY_SYSRQ":            ec.KEY_SYSRQ,
	"KEY_RIGHTALT":         ec.KEY_RIGHTALT,
	"KEY_LINEFEED":         ec.KEY_LINEFEED,
	"KEY_HOME":             ec.KEY_HOME,
	"KEY_UP":               ec.KEY_UP,
	"KEY_PAGEUP":           ec.KEY_PAGEUP,
	"KEY_LEFT":             ec.KEY_LEFT,
	"KEY_RIGHT":            ec.KEY_RIGHT,
	"KEY_END":              ec.KEY_END,
	"KEY_DOWN":             ec.KEY_DOWN,
	"KEY_PAGEDOWN":         ec.KEY_PAGEDOWN,
	"KEY_INSERT":           ec.KEY_INSERT,
	"KEY_DELETE":           ec.KEY_DELETE,
	"KEY_MACRO":            ec.KEY_MACRO,
	"KEY_MUTE":             ec.KEY_MUTE,
	"KEY_VOLUMEDOWN":       ec.KEY_VOLUMEDOWN,
	"KEY_VOLUMEUP":         ec.KEY_VOLUMEUP,
	"KEY_PAUSE":            ec.KEY_PAUSE,
	"KEY_KPCOMMA":          ec.KEY_KPCOMMA,
	"KEY_HANGEUL":          ec.KEY_HANGEUL,
	"KEY_HANGUEL":          ec.KEY_HANGUEL,
	"KEY_HANJA":            ec.KEY_HANJA,
	"KEY_YEN":              ec.KEY_YEN,
	"KEY_LEFTMETA":         ec.KEY_LEFTMETA,
	"KEY_RIGHTMETA":        ec.KEY_RIGHTMETA,
	"KEY_COMPOSE":          ec.KEY_COMPOSE,
	"KEY_F13":              ec.KEY_F13,
	"KEY_F14":              ec.KEY_F14,
	"KEY_F15":              ec.KEY_F15,
	"KEY_F16":              ec.KEY_F16,
	"KEY_F17":              ec.KEY_F17,
	"KEY_F18":              ec.KEY_F18,
	"KEY_F19":              ec.KEY_F19,
	"KEY_F20":              ec.KEY_F20,
	"KEY_F21":              ec.KEY_F21,
	"KEY_F22":              ec.KEY_F22,
	"KEY_F23":              ec.KEY_F23,
	"KEY_F24":              ec.KEY_F24,
	"KEY_MICMUTE":          ec.KEY_MICMUTE,
}

func checkExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func getConfigConts() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := ""

	{
		homeConfPath := filepath.Join(home, "go29")
		confPath := filepath.Join(home, ".config", "go29")

		if checkExists(homeConfPath) {
			path = homeConfPath
		} else if checkExists(confPath) {
			path = confPath
		}
	}

	if path == "" {
		return "", errors.New("no config file ~/go29 or ~/.config/go29")
	}

	conts, err := os.ReadFile(path)
	return string(conts), err
}

// removes lines starting with '#'
// also puts ';' after lines that don't end with '=' or ',' (end of mapping)
func removeComments(s string) string {
	var out string
	for _, l := range strings.Split(s, "\n") {
		if len(l) == 0 {
			continue
		}

		if !strings.HasPrefix(l, "#") {
			out += l
			if !strings.HasSuffix(l, "=") && !strings.HasSuffix(l, ",") {
				out += ";"
			}
		}
	}
	return out
}

func ParseRemapConfig() (map[wheelKey][]kbKey, error) {
	confConts, err := getConfigConts()
	if err != nil {
		return nil, err
	}

	confConts = strings.ReplaceAll(confConts, " ", "")
	confConts = removeComments(confConts)

	remaps := make(map[wheelKey][]kbKey, 0)

	for _, remapStr := range strings.Split(confConts, ";") {
		if len(remapStr) == 0 {
			continue
		}
		fromTo := strings.Split(remapStr, "=")
		if len(fromTo) <= 1 {
			return nil, errors.New("config syntax error, near: " + remapStr)
		}

		from, exists := fromMap[fromTo[0]]
		if !exists {
			return nil,
				errors.New("Config Parsing Error: \"from\" key: '" + fromTo[0] + "' not found. " +
					"see /go29/event_codes for all key types")
		}

		var (
			toStr = fromTo[1]

			to = make([]kbKey, 0)
		)

		for _, keyStr := range strings.Split(toStr, ",") {
			var (
				click    = true
				modifier = false
			)

			// press
			if strings.HasPrefix(keyStr, "(") && strings.HasSuffix(keyStr, ")") {
				keyStr = keyStr[1 : len(keyStr)-1]
				click = false
			}

			// mod
			if strings.HasPrefix(keyStr, "{") && strings.HasSuffix(keyStr, "}") {
				keyStr = keyStr[1 : len(keyStr)-1]
				click = false
				modifier = true
			}

			key, exists := toMap[keyStr]
			if !exists {
				return nil,
					errors.New("Config Parsing Error: \"to\" key: '" + keyStr + "' not found. " +
						"see /go29/event_codes for all key types")
			}

			to = append(to, kbKey{
				key:      key,
				click:    click,
				modifier: modifier,
			})
		}

		remaps[from] = to
	}

	return remaps, nil
}
