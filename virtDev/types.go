package virtDev

/*
#include <linux/uinput.h>
*/
import "C"

type WheelKey int

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

type KBKey uint8

const (
	KEY_ESC        KBKey = 1
	KEY_1          KBKey = 2
	KEY_2          KBKey = 3
	KEY_3          KBKey = 4
	KEY_4          KBKey = 5
	KEY_5          KBKey = 6
	KEY_6          KBKey = 7
	KEY_7          KBKey = 8
	KEY_8          KBKey = 9
	KEY_9          KBKey = 10
	KEY_0          KBKey = 11
	KEY_MINUS      KBKey = 12
	KEY_EQUAL      KBKey = 13
	KEY_BACKSPACE  KBKey = 14
	KEY_TAB        KBKey = 15
	KEY_Q          KBKey = 16
	KEY_W          KBKey = 17
	KEY_E          KBKey = 18
	KEY_R          KBKey = 19
	KEY_T          KBKey = 20
	KEY_Y          KBKey = 21
	KEY_U          KBKey = 22
	KEY_I          KBKey = 23
	KEY_O          KBKey = 24
	KEY_P          KBKey = 25
	KEY_LEFTBRACE  KBKey = 26
	KEY_RIGHTBRACE KBKey = 27
	KEY_ENTER      KBKey = 28
	KEY_LEFTCTRL   KBKey = 29
	KEY_A          KBKey = 30
	KEY_S          KBKey = 31
	KEY_D          KBKey = 32
	KEY_F          KBKey = 33
	KEY_G          KBKey = 34
	KEY_H          KBKey = 35
	KEY_J          KBKey = 36
	KEY_K          KBKey = 37
	KEY_L          KBKey = 38
	KEY_SEMICOLON  KBKey = 39
	KEY_APOSTROPHE KBKey = 40
	KEY_GRAVE      KBKey = 41
	KEY_LEFTSHIFT  KBKey = 42
	KEY_BACKSLASH  KBKey = 43
	KEY_Z          KBKey = 44
	KEY_X          KBKey = 45
	KEY_C          KBKey = 46
	KEY_V          KBKey = 47
	KEY_B          KBKey = 48
	KEY_N          KBKey = 49
	KEY_M          KBKey = 50
	KEY_COMMA      KBKey = 51
	KEY_DOT        KBKey = 52
	KEY_SLASH      KBKey = 53
	KEY_RIGHTSHIFT KBKey = 54
	KEY_KPASTERISK KBKey = 55
	KEY_LEFTALT    KBKey = 56
	KEY_SPACE      KBKey = 57
	KEY_CAPSLOCK   KBKey = 58
	KEY_F1         KBKey = 59
	KEY_F2         KBKey = 60
	KEY_F3         KBKey = 61
	KEY_F4         KBKey = 62
	KEY_F5         KBKey = 63
	KEY_F6         KBKey = 64
	KEY_F7         KBKey = 65
	KEY_F8         KBKey = 66
	KEY_F9         KBKey = 67
	KEY_F10        KBKey = 68
	KEY_NUMLOCK    KBKey = 69
	KEY_SCROLLLOCK KBKey = 70
	KEY_KP7        KBKey = 71
	KEY_KP8        KBKey = 72
	KEY_KP9        KBKey = 73
	KEY_KPMINUS    KBKey = 74
	KEY_KP4        KBKey = 75
	KEY_KP5        KBKey = 76
	KEY_KP6        KBKey = 77
	KEY_KPPLUS     KBKey = 78
	KEY_KP1        KBKey = 79
	KEY_KP2        KBKey = 80
	KEY_KP3        KBKey = 81
	KEY_KP0        KBKey = 82
	KEY_KPDOT      KBKey = 83

	KEY_ZENKAKUHANKAKU   KBKey = 85
	KEY_102ND            KBKey = 86
	KEY_F11              KBKey = 87
	KEY_F12              KBKey = 88
	KEY_RO               KBKey = 89
	KEY_KATAKANA         KBKey = 90
	KEY_HIRAGANA         KBKey = 91
	KEY_HENKAN           KBKey = 92
	KEY_KATAKANAHIRAGANA KBKey = 93
	KEY_MUHENKAN         KBKey = 94
	KEY_KPJPCOMMA        KBKey = 95
	KEY_KPENTER          KBKey = 96
	KEY_RIGHTCTRL        KBKey = 97
	KEY_KPSLASH          KBKey = 98
	KEY_SYSRQ            KBKey = 99
	KEY_RIGHTALT         KBKey = 100
	KEY_LINEFEED         KBKey = 101
	KEY_HOME             KBKey = 102
	KEY_UP               KBKey = 103
	KEY_PAGEUP           KBKey = 104
	KEY_LEFT             KBKey = 105
	KEY_RIGHT            KBKey = 106
	KEY_END              KBKey = 107
	KEY_DOWN             KBKey = 108
	KEY_PAGEDOWN         KBKey = 109
	KEY_INSERT           KBKey = 110
	KEY_DELETE           KBKey = 111
	KEY_MACRO            KBKey = 112
	KEY_MUTE             KBKey = 113
	KEY_VOLUMEDOWN       KBKey = 114
	KEY_VOLUMEUP         KBKey = 115
	KEY_POWER            KBKey = 116 /* SC System Power Down */
	KEY_KPEQUAL          KBKey = 117
	KEY_KPPLUSMINUS      KBKey = 118
	KEY_PAUSE            KBKey = 119
	KEY_SCALE            KBKey = 120 /* AL Compiz Scale (Expose) */

	KEY_KPCOMMA   KBKey = 121
	KEY_HANGEUL   KBKey = 122
	KEY_HANGUEL   KBKey = KEY_HANGEUL
	KEY_HANJA     KBKey = 123
	KEY_YEN       KBKey = 124
	KEY_LEFTMETA  KBKey = 125
	KEY_RIGHTMETA KBKey = 126
	KEY_COMPOSE   KBKey = 127

	KEY_STOP           KBKey = 128 /* AC Stop */
	KEY_AGAIN          KBKey = 129
	KEY_PROPS          KBKey = 130 /* AC Properties */
	KEY_UNDO           KBKey = 131 /* AC Undo */
	KEY_FRONT          KBKey = 132
	KEY_COPY           KBKey = 133 /* AC Copy */
	KEY_OPEN           KBKey = 134 /* AC Open */
	KEY_PASTE          KBKey = 135 /* AC Paste */
	KEY_FIND           KBKey = 136 /* AC Search */
	KEY_CUT            KBKey = 137 /* AC Cut */
	KEY_HELP           KBKey = 138 /* AL Integrated Help Center */
	KEY_MENU           KBKey = 139 /* Menu (show menu) */
	KEY_CALC           KBKey = 140 /* AL Calculator */
	KEY_SETUP          KBKey = 141
	KEY_SLEEP          KBKey = 142 /* SC System Sleep */
	KEY_WAKEUP         KBKey = 143 /* System Wake Up */
	KEY_FILE           KBKey = 144 /* AL Local Machine Browser */
	KEY_SENDFILE       KBKey = 145
	KEY_DELETEFILE     KBKey = 146
	KEY_XFER           KBKey = 147
	KEY_PROG1          KBKey = 148
	KEY_PROG2          KBKey = 149
	KEY_WWW            KBKey = 150 /* AL Internet Browser */
	KEY_MSDOS          KBKey = 151
	KEY_COFFEE         KBKey = 152 /* AL Terminal Lock/Screensaver */
	KEY_SCREENLOCK     KBKey = KEY_COFFEE
	KEY_ROTATE_DISPLAY KBKey = 153 /* Display orientation for e.g. tablets */
	KEY_DIRECTION      KBKey = KEY_ROTATE_DISPLAY
	KEY_CYCLEWINDOWS   KBKey = 154
	KEY_MAIL           KBKey = 155
	KEY_BOOKMARKS      KBKey = 156 /* AC Bookmarks */
	KEY_COMPUTER       KBKey = 157
	KEY_BACK           KBKey = 158 /* AC Back */
	KEY_FORWARD        KBKey = 159 /* AC Forward */
	KEY_CLOSECD        KBKey = 160
	KEY_EJECTCD        KBKey = 161
	KEY_EJECTCLOSECD   KBKey = 162
	KEY_NEXTSONG       KBKey = 163
	KEY_PLAYPAUSE      KBKey = 164
	KEY_PREVIOUSSONG   KBKey = 165
	KEY_STOPCD         KBKey = 166
	KEY_RECORD         KBKey = 167
	KEY_REWIND         KBKey = 168
	KEY_PHONE          KBKey = 169 /* Media Select Telephone */
	KEY_ISO            KBKey = 170
	KEY_CONFIG         KBKey = 171 /* AL Consumer Control Configuration */
	KEY_HOMEPAGE       KBKey = 172 /* AC Home */
	KEY_REFRESH        KBKey = 173 /* AC Refresh */
	KEY_EXIT           KBKey = 174 /* AC Exit */
	KEY_MOVE           KBKey = 175
	KEY_EDIT           KBKey = 176
	KEY_SCROLLUP       KBKey = 177
	KEY_SCROLLDOWN     KBKey = 178
	KEY_KPLEFTPAREN    KBKey = 179
	KEY_KPRIGHTPAREN   KBKey = 180
	KEY_NEW            KBKey = 181 /* AC New */
	KEY_REDO           KBKey = 182 /* AC Redo/Repeat */

	KEY_F13 KBKey = 183
	KEY_F14 KBKey = 184
	KEY_F15 KBKey = 185
	KEY_F16 KBKey = 186
	KEY_F17 KBKey = 187
	KEY_F18 KBKey = 188
	KEY_F19 KBKey = 189
	KEY_F20 KBKey = 190
	KEY_F21 KBKey = 191
	KEY_F22 KBKey = 192
	KEY_F23 KBKey = 193
	KEY_F24 KBKey = 194

	KEY_MICMUTE KBKey = 248
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
