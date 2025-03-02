package virtDev

/*
#include <linux/uinput.h>
*/
import "C"

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
