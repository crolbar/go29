package virtDev

/*
#include <linux/input-event-codes.h>
*/
import "C"
import (
	"errors"
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

type remap struct {
	from wheelKey
	to   []kbKey
}

var fromMap = map[string]wheelKey{
	"BTN_X":               BTN_X,
	"BTN_CIRCLE":          BTN_CIRCLE,
	"BTN_SQUARE":          BTN_SQUARE,
	"BTN_TRIANGLE":        BTN_TRIANGLE,
	"BTN_R2":              BTN_R2,
	"BTN_R3":              BTN_R3,
	"BTN_L2":              BTN_L2,
	"BTN_L3":              BTN_L3,
	"BTN_ROTARY_BUTTON":   BTN_ROTARY_BUTTON,
	"BTN_ROTARY_RIGHT":    BTN_ROTARY_RIGHT,
	"BTN_ROTARY_LEFT":     BTN_ROTARY_LEFT,
	"BTN_SHARE":           BTN_SHARE,
	"BTN_OPTIONS":         BTN_OPTIONS,
	"BTN_PS":              BTN_PS,
	"BTN_MINUS":           BTN_MINUS,
	"BTN_PLUS":            BTN_PLUS,
	"BTN_RIGHT_PADDLE":    BTN_RIGHT_PADDLE,
	"BTN_LEFT_PADDLE":     BTN_LEFT_PADDLE,
	"BTN_SHIFTER_FIRST":   BTN_SHIFTER_FIRST,
	"BTN_SHIFTER_SECOND":  BTN_SHIFTER_SECOND,
	"BTN_SHIFTER_THIRD":   BTN_SHIFTER_THIRD,
	"BTN_SHIFTER_FOURTH":  BTN_SHIFTER_FOURTH,
	"BTN_SHIFTER_FIFT":    BTN_SHIFTER_FIFT,
	"BTN_SHIFTER_SIXTH":   BTN_SHIFTER_SIXTH,
	"BTN_SHIFTER_REVERSE": BTN_SHIFTER_REVERSE,
	"ABS_WHEEL":           ABS_WHEEL,
	"ABS_CLUTCH":          ABS_CLUTCH,
	"ABS_THROTTLE":        ABS_THROTTLE,
	"ABS_BREAK":           ABS_BREAK,
	"ABS_DPADX":           ABS_DPADX,
	"ABS_DPADY":           ABS_DPADY,
}

var toMap = map[string]key{
	"KEY_ESC":              KEY_ESC,
	"KEY_1":                KEY_1,
	"KEY_2":                KEY_2,
	"KEY_3":                KEY_3,
	"KEY_4":                KEY_4,
	"KEY_5":                KEY_5,
	"KEY_6":                KEY_6,
	"KEY_7":                KEY_7,
	"KEY_8":                KEY_8,
	"KEY_9":                KEY_9,
	"KEY_0":                KEY_0,
	"KEY_MINUS":            KEY_MINUS,
	"KEY_EQUAL":            KEY_EQUAL,
	"KEY_BACKSPACE":        KEY_BACKSPACE,
	"KEY_TAB":              KEY_TAB,
	"KEY_Q":                KEY_Q,
	"KEY_W":                KEY_W,
	"KEY_E":                KEY_E,
	"KEY_R":                KEY_R,
	"KEY_T":                KEY_T,
	"KEY_Y":                KEY_Y,
	"KEY_U":                KEY_U,
	"KEY_I":                KEY_I,
	"KEY_O":                KEY_O,
	"KEY_P":                KEY_P,
	"KEY_LEFTBRACE":        KEY_LEFTBRACE,
	"KEY_RIGHTBRACE":       KEY_RIGHTBRACE,
	"KEY_ENTER":            KEY_ENTER,
	"KEY_LEFTCTRL":         KEY_LEFTCTRL,
	"KEY_A":                KEY_A,
	"KEY_S":                KEY_S,
	"KEY_D":                KEY_D,
	"KEY_F":                KEY_F,
	"KEY_G":                KEY_G,
	"KEY_H":                KEY_H,
	"KEY_J":                KEY_J,
	"KEY_K":                KEY_K,
	"KEY_L":                KEY_L,
	"KEY_SEMICOLON":        KEY_SEMICOLON,
	"KEY_APOSTROPHE":       KEY_APOSTROPHE,
	"KEY_GRAVE":            KEY_GRAVE,
	"KEY_LEFTSHIFT":        KEY_LEFTSHIFT,
	"KEY_BACKSLASH":        KEY_BACKSLASH,
	"KEY_Z":                KEY_Z,
	"KEY_X":                KEY_X,
	"KEY_C":                KEY_C,
	"KEY_V":                KEY_V,
	"KEY_B":                KEY_B,
	"KEY_N":                KEY_N,
	"KEY_M":                KEY_M,
	"KEY_COMMA":            KEY_COMMA,
	"KEY_DOT":              KEY_DOT,
	"KEY_SLASH":            KEY_SLASH,
	"KEY_RIGHTSHIFT":       KEY_RIGHTSHIFT,
	"KEY_KPASTERISK":       KEY_KPASTERISK,
	"KEY_LEFTALT":          KEY_LEFTALT,
	"KEY_SPACE":            KEY_SPACE,
	"KEY_CAPSLOCK":         KEY_CAPSLOCK,
	"KEY_F1":               KEY_F1,
	"KEY_F2":               KEY_F2,
	"KEY_F3":               KEY_F3,
	"KEY_F4":               KEY_F4,
	"KEY_F5":               KEY_F5,
	"KEY_F6":               KEY_F6,
	"KEY_F7":               KEY_F7,
	"KEY_F8":               KEY_F8,
	"KEY_F9":               KEY_F9,
	"KEY_F10":              KEY_F10,
	"KEY_NUMLOCK":          KEY_NUMLOCK,
	"KEY_SCROLLLOCK":       KEY_SCROLLLOCK,
	"KEY_KP7":              KEY_KP7,
	"KEY_KP8":              KEY_KP8,
	"KEY_KP9":              KEY_KP9,
	"KEY_KPMINUS":          KEY_KPMINUS,
	"KEY_KP4":              KEY_KP4,
	"KEY_KP5":              KEY_KP5,
	"KEY_KP6":              KEY_KP6,
	"KEY_KPPLUS":           KEY_KPPLUS,
	"KEY_KP1":              KEY_KP1,
	"KEY_KP2":              KEY_KP2,
	"KEY_KP3":              KEY_KP3,
	"KEY_KP0":              KEY_KP0,
	"KEY_KPDOT":            KEY_KPDOT,
	"KEY_ZENKAKUHANKAKU":   KEY_ZENKAKUHANKAKU,
	"KEY_102ND":            KEY_102ND,
	"KEY_F11":              KEY_F11,
	"KEY_F12":              KEY_F12,
	"KEY_RO":               KEY_RO,
	"KEY_KATAKANA":         KEY_KATAKANA,
	"KEY_HIRAGANA":         KEY_HIRAGANA,
	"KEY_HENKAN":           KEY_HENKAN,
	"KEY_KATAKANAHIRAGANA": KEY_KATAKANAHIRAGANA,
	"KEY_MUHENKAN":         KEY_MUHENKAN,
	"KEY_KPJPCOMMA":        KEY_KPJPCOMMA,
	"KEY_KPENTER":          KEY_KPENTER,
	"KEY_RIGHTCTRL":        KEY_RIGHTCTRL,
	"KEY_KPSLASH":          KEY_KPSLASH,
	"KEY_SYSRQ":            KEY_SYSRQ,
	"KEY_RIGHTALT":         KEY_RIGHTALT,
	"KEY_LINEFEED":         KEY_LINEFEED,
	"KEY_HOME":             KEY_HOME,
	"KEY_UP":               KEY_UP,
	"KEY_PAGEUP":           KEY_PAGEUP,
	"KEY_LEFT":             KEY_LEFT,
	"KEY_RIGHT":            KEY_RIGHT,
	"KEY_END":              KEY_END,
	"KEY_DOWN":             KEY_DOWN,
	"KEY_PAGEDOWN":         KEY_PAGEDOWN,
	"KEY_INSERT":           KEY_INSERT,
	"KEY_DELETE":           KEY_DELETE,
	"KEY_MACRO":            KEY_MACRO,
	"KEY_MUTE":             KEY_MUTE,
	"KEY_VOLUMEDOWN":       KEY_VOLUMEDOWN,
	"KEY_VOLUMEUP":         KEY_VOLUMEUP,
	"KEY_PAUSE":            KEY_PAUSE,
	"KEY_KPCOMMA":          KEY_KPCOMMA,
	"KEY_HANGEUL":          KEY_HANGEUL,
	"KEY_HANGUEL":          KEY_HANGUEL,
	"KEY_HANJA":            KEY_HANJA,
	"KEY_YEN":              KEY_YEN,
	"KEY_LEFTMETA":         KEY_LEFTMETA,
	"KEY_RIGHTMETA":        KEY_RIGHTMETA,
	"KEY_COMPOSE":          KEY_COMPOSE,
	"KEY_F13":              KEY_F13,
	"KEY_F14":              KEY_F14,
	"KEY_F15":              KEY_F15,
	"KEY_F16":              KEY_F16,
	"KEY_F17":              KEY_F17,
	"KEY_F18":              KEY_F18,
	"KEY_F19":              KEY_F19,
	"KEY_F20":              KEY_F20,
	"KEY_F21":              KEY_F21,
	"KEY_F22":              KEY_F22,
	"KEY_F23":              KEY_F23,
	"KEY_F24":              KEY_F24,
	"KEY_MICMUTE":          KEY_MICMUTE,
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

func ParseRemapConfig() ([]remap, error) {
	confConts, err := getConfigConts()
	if err != nil {
		return []remap{}, err
	}

	confConts = strings.ReplaceAll(confConts, " ", "")
	confConts = removeComments(confConts)

	remaps := make([]remap, 0)

	for _, remapStr := range strings.Split(confConts, ";") {
		if len(remapStr) == 0 {
			continue
		}

		var (
			fromTo = strings.Split(remapStr, "=")

			from  = fromMap[fromTo[0]]
			toStr = fromTo[1]

			to = make([]kbKey, 0)
		)

		for _, key := range strings.Split(toStr, ",") {
			var (
				click    = true
				modifier = false
			)

			// press
			if strings.HasPrefix(key, "(") && strings.HasSuffix(key, ")") {
				key = key[1 : len(key)-1]
				click = false
			}

			// mod
			if strings.HasPrefix(key, "{") && strings.HasSuffix(key, "}") {
				key = key[1 : len(key)-1]
				click = false
				modifier = true
			}

			to = append(to, kbKey{
				key:      toMap[key],
				click:    click,
				modifier: modifier,
			})
		}

		remaps = append(remaps, remap{
			from: from,
			to:   to,
		})
	}

	return remaps, nil
}
