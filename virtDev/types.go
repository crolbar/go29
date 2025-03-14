package virtDev

/*
#include <linux/uinput.h>
*/
import "C"

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
	BTN_SHIFTER_FIFT    = C.BTN_TRIGGER_HAPPY1
	BTN_SHIFTER_SIXTH   = C.BTN_TRIGGER_HAPPY2
	BTN_SHIFTER_REVERSE = C.BTN_TRIGGER_HAPPY3

	ABS_WHEEL    = C.ABS_X  // Steering wheel
	ABS_CLUTCH   = C.ABS_Y  // Clutch
	ABS_THROTTLE = C.ABS_Z  // Throttle
	ABS_BREAK    = C.ABS_RZ // Break

	ABS_DPADX = C.ABS_HAT0X // D-pad X
	ABS_DPADY = C.ABS_HAT0Y // D-pad Y
)

const (
	KEY_ESC        key = 1
	KEY_1          key = 2
	KEY_2          key = 3
	KEY_3          key = 4
	KEY_4          key = 5
	KEY_5          key = 6
	KEY_6          key = 7
	KEY_7          key = 8
	KEY_8          key = 9
	KEY_9          key = 10
	KEY_0          key = 11
	KEY_MINUS      key = 12
	KEY_EQUAL      key = 13
	KEY_BACKSPACE  key = 14
	KEY_TAB        key = 15
	KEY_Q          key = 16
	KEY_W          key = 17
	KEY_E          key = 18
	KEY_R          key = 19
	KEY_T          key = 20
	KEY_Y          key = 21
	KEY_U          key = 22
	KEY_I          key = 23
	KEY_O          key = 24
	KEY_P          key = 25
	KEY_LEFTBRACE  key = 26
	KEY_RIGHTBRACE key = 27
	KEY_ENTER      key = 28
	KEY_LEFTCTRL   key = 29
	KEY_A          key = 30
	KEY_S          key = 31
	KEY_D          key = 32
	KEY_F          key = 33
	KEY_G          key = 34
	KEY_H          key = 35
	KEY_J          key = 36
	KEY_K          key = 37
	KEY_L          key = 38
	KEY_SEMICOLON  key = 39
	KEY_APOSTROPHE key = 40
	KEY_GRAVE      key = 41
	KEY_LEFTSHIFT  key = 42
	KEY_BACKSLASH  key = 43
	KEY_Z          key = 44
	KEY_X          key = 45
	KEY_C          key = 46
	KEY_V          key = 47
	KEY_B          key = 48
	KEY_N          key = 49
	KEY_M          key = 50
	KEY_COMMA      key = 51
	KEY_DOT        key = 52
	KEY_SLASH      key = 53
	KEY_RIGHTSHIFT key = 54
	KEY_KPASTERISK key = 55
	KEY_LEFTALT    key = 56
	KEY_SPACE      key = 57
	KEY_CAPSLOCK   key = 58
	KEY_F1         key = 59
	KEY_F2         key = 60
	KEY_F3         key = 61
	KEY_F4         key = 62
	KEY_F5         key = 63
	KEY_F6         key = 64
	KEY_F7         key = 65
	KEY_F8         key = 66
	KEY_F9         key = 67
	KEY_F10        key = 68
	KEY_NUMLOCK    key = 69
	KEY_SCROLLLOCK key = 70
	KEY_KP7        key = 71
	KEY_KP8        key = 72
	KEY_KP9        key = 73
	KEY_KPMINUS    key = 74
	KEY_KP4        key = 75
	KEY_KP5        key = 76
	KEY_KP6        key = 77
	KEY_KPPLUS     key = 78
	KEY_KP1        key = 79
	KEY_KP2        key = 80
	KEY_KP3        key = 81
	KEY_KP0        key = 82
	KEY_KPDOT      key = 83

	KEY_ZENKAKUHANKAKU   key = 85
	KEY_102ND            key = 86
	KEY_F11              key = 87
	KEY_F12              key = 88
	KEY_RO               key = 89
	KEY_KATAKANA         key = 90
	KEY_HIRAGANA         key = 91
	KEY_HENKAN           key = 92
	KEY_KATAKANAHIRAGANA key = 93
	KEY_MUHENKAN         key = 94
	KEY_KPJPCOMMA        key = 95
	KEY_KPENTER          key = 96
	KEY_RIGHTCTRL        key = 97
	KEY_KPSLASH          key = 98
	KEY_SYSRQ            key = 99
	KEY_RIGHTALT         key = 100
	KEY_LINEFEED         key = 101
	KEY_HOME             key = 102
	KEY_UP               key = 103
	KEY_PAGEUP           key = 104
	KEY_LEFT             key = 105
	KEY_RIGHT            key = 106
	KEY_END              key = 107
	KEY_DOWN             key = 108
	KEY_PAGEDOWN         key = 109
	KEY_INSERT           key = 110
	KEY_DELETE           key = 111
	KEY_MACRO            key = 112
	KEY_MUTE             key = 113
	KEY_VOLUMEDOWN       key = 114
	KEY_VOLUMEUP         key = 115
	KEY_PAUSE            key = 119

	KEY_KPCOMMA   key = 121
	KEY_HANGEUL   key = 122
	KEY_HANGUEL   key = KEY_HANGEUL
	KEY_HANJA     key = 123
	KEY_YEN       key = 124
	KEY_LEFTMETA  key = 125
	KEY_RIGHTMETA key = 126
	KEY_COMPOSE   key = 127

	KEY_F13 key = 183
	KEY_F14 key = 184
	KEY_F15 key = 185
	KEY_F16 key = 186
	KEY_F17 key = 187
	KEY_F18 key = 188
	KEY_F19 key = 189
	KEY_F20 key = 190
	KEY_F21 key = 191
	KEY_F22 key = 192
	KEY_F23 key = 193
	KEY_F24 key = 194

	KEY_MICMUTE key = 248
)

type inputId struct {
	bustype uint16
	vendor  uint16
	product uint16
	version uint16
}

type uinputSetup struct {
	id             inputId
	name           [C.UINPUT_MAX_NAME_SIZE]byte
	ff_effects_max uint32
}

type inputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

type uinputUserDev struct {
	id             inputId
	name           [C.UINPUT_MAX_NAME_SIZE]byte
	ff_effects_max uint32
	absmax         [C.ABS_CNT]int32
	absmin         [C.ABS_CNT]int32
	absfuzz        [C.ABS_CNT]int32
	absflat        [C.ABS_CNT]int32
}

type FF_Replay struct {
	length uint16
	delay  uint16
}

type FF_Trigger struct {
	button   uint16
	interval uint16
}

type FF_Envelope struct {
	attack_length uint16
	attack_level  uint16
	fade_length   uint16
	fade_level    uint16
}

type FF_Constant_Effect struct {
	level    int16
	envelope FF_Envelope
}

type FF_Effect struct {
	etype     uint16
	id        int16
	direction uint16

	trigger FF_Trigger
	replay  FF_Replay

	_pad [2]byte // 2 padding

	u [32]byte // 32 becasue of FF_Periodic_Effect
}
