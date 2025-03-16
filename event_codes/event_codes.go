package event_codes

// EVERYTHING IS SPECIFIC FOR THE Logitech G29 PS

/*
#include <linux/uinput.h>
*/
import "C"

// TYPES
const (
	EV_SYN = 0x00 // sync
	EV_KEY = 0x01 // send on buttons
	EV_MSC = 0x04 // send on buttons (unneeded)
)

// CODES
const (
	ABS_WHEEL    = C.ABS_X  // Steering wheel
	ABS_CLUTCH   = C.ABS_Y  // Clutch
	ABS_THROTTLE = C.ABS_Z  // Throttle
	ABS_BREAK    = C.ABS_RZ // Break

	ABS_DPADX = C.ABS_HAT0X // D-pad X
	ABS_DPADY = C.ABS_HAT0Y // D-pad Y
)

// BUTTON CODES
const (
	BTN_X        = C.BTN_TRIGGER
	BTN_CIRCLE   = C.BTN_THUMB2
	BTN_SQUARE   = C.BTN_THUMB
	BTN_TRIANGLE = C.BTN_TOP

	BTN_R2 = C.BTN_BASE
	BTN_R3 = C.BTN_BASE5

	BTN_L2 = C.BTN_BASE2
	BTN_L3 = C.BTN_BASE6

	BTN_ROTARY_BUTTON = C.BTN_TRIGGER_HAPPY8
	BTN_ROTARY_RIGHT  = C.BTN_TRIGGER_HAPPY6
	BTN_ROTARY_LEFT   = C.BTN_TRIGGER_HAPPY7

	BTN_SHARE   = C.BTN_BASE3
	BTN_OPTIONS = C.BTN_BASE4
	BTN_PS      = C.BTN_TRIGGER_HAPPY9

	BTN_MINUS = C.BTN_TRIGGER_HAPPY5
	BTN_PLUS  = C.BTN_TRIGGER_HAPPY4

	BTN_RIGHT_PADDLE = C.BTN_TOP2
	BTN_LEFT_PADDLE  = C.BTN_PINKIE

	BTN_SHIFTER_FIRST   = 0x12C
	BTN_SHIFTER_SECOND  = 0x12D
	BTN_SHIFTER_THIRD   = 0x12E
	BTN_SHIFTER_FOURTH  = C.BTN_DEAD
	BTN_SHIFTER_FIFTH   = C.BTN_TRIGGER_HAPPY1
	BTN_SHIFTER_SIXTH   = C.BTN_TRIGGER_HAPPY2
	BTN_SHIFTER_REVERSE = C.BTN_TRIGGER_HAPPY3
)

// D-pad (code 16-17)
const (
	// code 17
	DPAD_UP   = -1
	DPAD_DOWN = 1

	// code 16
	DPAD_LEFT  = -1
	DPAD_RIGHT = 1

	DPAD_RELEASE = 0
)

// KEYBOARD CODES
const (
	KEY_ESC        = 1
	KEY_1          = 2
	KEY_2          = 3
	KEY_3          = 4
	KEY_4          = 5
	KEY_5          = 6
	KEY_6          = 7
	KEY_7          = 8
	KEY_8          = 9
	KEY_9          = 10
	KEY_0          = 11
	KEY_MINUS      = 12
	KEY_EQUAL      = 13
	KEY_BACKSPACE  = 14
	KEY_TAB        = 15
	KEY_Q          = 16
	KEY_W          = 17
	KEY_E          = 18
	KEY_R          = 19
	KEY_T          = 20
	KEY_Y          = 21
	KEY_U          = 22
	KEY_I          = 23
	KEY_O          = 24
	KEY_P          = 25
	KEY_LEFTBRACE  = 26
	KEY_RIGHTBRACE = 27
	KEY_ENTER      = 28
	KEY_LEFTCTRL   = 29
	KEY_A          = 30
	KEY_S          = 31
	KEY_D          = 32
	KEY_F          = 33
	KEY_G          = 34
	KEY_H          = 35
	KEY_J          = 36
	KEY_K          = 37
	KEY_L          = 38
	KEY_SEMICOLON  = 39
	KEY_APOSTROPHE = 40
	KEY_GRAVE      = 41
	KEY_LEFTSHIFT  = 42
	KEY_BACKSLASH  = 43
	KEY_Z          = 44
	KEY_X          = 45
	KEY_C          = 46
	KEY_V          = 47
	KEY_B          = 48
	KEY_N          = 49
	KEY_M          = 50
	KEY_COMMA      = 51
	KEY_DOT        = 52
	KEY_SLASH      = 53
	KEY_RIGHTSHIFT = 54
	KEY_KPASTERISK = 55
	KEY_LEFTALT    = 56
	KEY_SPACE      = 57
	KEY_CAPSLOCK   = 58
	KEY_F1         = 59
	KEY_F2         = 60
	KEY_F3         = 61
	KEY_F4         = 62
	KEY_F5         = 63
	KEY_F6         = 64
	KEY_F7         = 65
	KEY_F8         = 66
	KEY_F9         = 67
	KEY_F10        = 68
	KEY_NUMLOCK    = 69
	KEY_SCROLLLOCK = 70
	KEY_KP7        = 71
	KEY_KP8        = 72
	KEY_KP9        = 73
	KEY_KPMINUS    = 74
	KEY_KP4        = 75
	KEY_KP5        = 76
	KEY_KP6        = 77
	KEY_KPPLUS     = 78
	KEY_KP1        = 79
	KEY_KP2        = 80
	KEY_KP3        = 81
	KEY_KP0        = 82
	KEY_KPDOT      = 83

	KEY_ZENKAKUHANKAKU   = 85
	KEY_102ND            = 86
	KEY_F11              = 87
	KEY_F12              = 88
	KEY_RO               = 89
	KEY_KATAKANA         = 90
	KEY_HIRAGANA         = 91
	KEY_HENKAN           = 92
	KEY_KATAKANAHIRAGANA = 93
	KEY_MUHENKAN         = 94
	KEY_KPJPCOMMA        = 95
	KEY_KPENTER          = 96
	KEY_RIGHTCTRL        = 97
	KEY_KPSLASH          = 98
	KEY_SYSRQ            = 99
	KEY_RIGHTALT         = 100
	KEY_LINEFEED         = 101
	KEY_HOME             = 102
	KEY_UP               = 103
	KEY_PAGEUP           = 104
	KEY_LEFT             = 105
	KEY_RIGHT            = 106
	KEY_END              = 107
	KEY_DOWN             = 108
	KEY_PAGEDOWN         = 109
	KEY_INSERT           = 110
	KEY_DELETE           = 111
	KEY_MACRO            = 112
	KEY_MUTE             = 113
	KEY_VOLUMEDOWN       = 114
	KEY_VOLUMEUP         = 115
	KEY_PAUSE            = 119

	KEY_KPCOMMA   = 121
	KEY_HANGEUL   = 122
	KEY_HANGUEL   = KEY_HANGEUL
	KEY_HANJA     = 123
	KEY_YEN       = 124
	KEY_LEFTMETA  = 125
	KEY_RIGHTMETA = 126
	KEY_COMPOSE   = 127

	KEY_F13 = 183
	KEY_F14 = 184
	KEY_F15 = 185
	KEY_F16 = 186
	KEY_F17 = 187
	KEY_F18 = 188
	KEY_F19 = 189
	KEY_F20 = 190
	KEY_F21 = 191
	KEY_F22 = 192
	KEY_F23 = 193
	KEY_F24 = 194

	KEY_MICMUTE = 248
)
