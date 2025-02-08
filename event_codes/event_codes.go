package event_codes
// EVERYTHING IS SPECIFIC FOR THE Logitech G29 PS

// CODES
const (
	ABS_X     = 0x00 // wheel turn left/right
	ABS_Y     = 0x01 // clutch
	ABS_Z     = 0x02 // throttle
	ABS_RY    = 0x04 // buttons exept d-pad
	ABS_RZ    = 0x05 // break
	ABS_HAT0X = 0x10 // d-pad up/down
	ABS_HAT0Y = 0x11 // d-pad left/right
)

// button values under code 4
const (
	BUTTON_OPTIONS = 0x9000A
	BUTTON_PS      = 0x90019
	BUTTON_SHARE   = 0x90009

	BUTTON_R3      = 0x9000B
	BUTTON_L3      = 0x9000C
	BUTTON_R2      = 0x90007
	BUTTON_L2      = 0x90008
	BUTTON_RPaddle = 0x90005
	BUTTON_LPaddle = 0x90006

	BUTTON_RotaryL = 0x90017
	BUTTON_RotaryR = 0x90016
	BUTTON_Rotary  = 0x90018

	BUTTON_T = 0x90004
	BUTTON_S = 0x90002
	BUTTON_C = 0x90003
	BUTTON_X = 0x90001

	BUTTON_Plus  = 0x90014
	BUTTON_Minus = 0x90015

	SHIFTER_FIRST   = 0x9000D
	SHIFTER_SECOND  = 0x9000E
	SHIFTER_THIRD   = 0x9000F
	SHIFTER_FOURTH  = 0x90010
	SHIFTER_FIFTH   = 0x90011
	SHIFTER_SIXTH   = 0x90012
	SHIFTER_REVERSE = 0x90013
)

// D-pad (code 16-17)
const (
	DPAD_UP      = -1
	DPAD_DOWN    = 1
	DPAD_LEFT    = -1
	DPAD_RIGHT   = 1
	DPAD_RELEASE = 0
)
