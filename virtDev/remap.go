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

type Remap struct {
	from WheelKey
	to   []KBKey

	// press & release every button in order
	click bool

	// press & hold the first button and click the others
	modified bool
}

var fromMap = map[string]WheelKey{
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

var toMap = map[string]KBKey{
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
	"KEY_POWER":            KEY_POWER,
	"KEY_KPEQUAL":          KEY_KPEQUAL,
	"KEY_KPPLUSMINUS":      KEY_KPPLUSMINUS,
	"KEY_PAUSE":            KEY_PAUSE,
	"KEY_SCALE":            KEY_SCALE,
	"KEY_KPCOMMA":          KEY_KPCOMMA,
	"KEY_HANGEUL":          KEY_HANGEUL,
	"KEY_HANGUEL":          KEY_HANGUEL,
	"KEY_HANJA":            KEY_HANJA,
	"KEY_YEN":              KEY_YEN,
	"KEY_LEFTMETA":         KEY_LEFTMETA,
	"KEY_RIGHTMETA":        KEY_RIGHTMETA,
	"KEY_COMPOSE":          KEY_COMPOSE,
	"KEY_STOP":             KEY_STOP,
	"KEY_AGAIN":            KEY_AGAIN,
	"KEY_PROPS":            KEY_PROPS,
	"KEY_UNDO":             KEY_UNDO,
	"KEY_FRONT":            KEY_FRONT,
	"KEY_COPY":             KEY_COPY,
	"KEY_OPEN":             KEY_OPEN,
	"KEY_PASTE":            KEY_PASTE,
	"KEY_FIND":             KEY_FIND,
	"KEY_CUT":              KEY_CUT,
	"KEY_HELP":             KEY_HELP,
	"KEY_MENU":             KEY_MENU,
	"KEY_CALC":             KEY_CALC,
	"KEY_SETUP":            KEY_SETUP,
	"KEY_SLEEP":            KEY_SLEEP,
	"KEY_WAKEUP":           KEY_WAKEUP,
	"KEY_FILE":             KEY_FILE,
	"KEY_SENDFILE":         KEY_SENDFILE,
	"KEY_DELETEFILE":       KEY_DELETEFILE,
	"KEY_XFER":             KEY_XFER,
	"KEY_PROG1":            KEY_PROG1,
	"KEY_PROG2":            KEY_PROG2,
	"KEY_WWW":              KEY_WWW,
	"KEY_MSDOS":            KEY_MSDOS,
	"KEY_COFFEE":           KEY_COFFEE,
	"KEY_SCREENLOCK":       KEY_SCREENLOCK,
	"KEY_ROTATE_DISPLAY":   KEY_ROTATE_DISPLAY,
	"KEY_DIRECTION":        KEY_DIRECTION,
	"KEY_CYCLEWINDOWS":     KEY_CYCLEWINDOWS,
	"KEY_MAIL":             KEY_MAIL,
	"KEY_BOOKMARKS":        KEY_BOOKMARKS,
	"KEY_COMPUTER":         KEY_COMPUTER,
	"KEY_BACK":             KEY_BACK,
	"KEY_FORWARD":          KEY_FORWARD,
	"KEY_CLOSECD":          KEY_CLOSECD,
	"KEY_EJECTCD":          KEY_EJECTCD,
	"KEY_EJECTCLOSECD":     KEY_EJECTCLOSECD,
	"KEY_NEXTSONG":         KEY_NEXTSONG,
	"KEY_PLAYPAUSE":        KEY_PLAYPAUSE,
	"KEY_PREVIOUSSONG":     KEY_PREVIOUSSONG,
	"KEY_STOPCD":           KEY_STOPCD,
	"KEY_RECORD":           KEY_RECORD,
	"KEY_REWIND":           KEY_REWIND,
	"KEY_PHONE":            KEY_PHONE,
	"KEY_ISO":              KEY_ISO,
	"KEY_CONFIG":           KEY_CONFIG,
	"KEY_HOMEPAGE":         KEY_HOMEPAGE,
	"KEY_REFRESH":          KEY_REFRESH,
	"KEY_EXIT":             KEY_EXIT,
	"KEY_MOVE":             KEY_MOVE,
	"KEY_EDIT":             KEY_EDIT,
	"KEY_SCROLLUP":         KEY_SCROLLUP,
	"KEY_SCROLLDOWN":       KEY_SCROLLDOWN,
	"KEY_KPLEFTPAREN":      KEY_KPLEFTPAREN,
	"KEY_KPRIGHTPAREN":     KEY_KPRIGHTPAREN,
	"KEY_NEW":              KEY_NEW,
	"KEY_REDO":             KEY_REDO,
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

func ParseRemapConfig() ([]Remap, error) {
	confConts, err := getConfigConts()
	if err != nil {
		return []Remap{}, nil
	}

	confConts = strings.ReplaceAll(confConts, " ", "")
	confConts = removeComments(confConts)

	remaps := make([]Remap, 0)

	for _, remap := range strings.Split(confConts, ";") {
		if len(remap) == 0 {
			continue
		}

		var (
			fromTo = strings.Split(remap, "=")

			from  = fromMap[fromTo[0]]
			toStr = fromTo[len(fromTo)-1]

			to = make([]KBKey, 0)
		)

		for _, key := range strings.Split(toStr, ",") {
			to = append(to, toMap[key])
		}

		remaps = append(remaps, Remap{
			from:     from,
			to:       to,
			modified: len(fromTo) > 3,
			click:    len(fromTo) > 2,
		})
	}

	return remaps, nil
}
