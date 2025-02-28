package device

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

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

type FF_Ramp_Effect struct {
	start_level int16
	end_level   int16
	envelope    FF_Envelope
}

type FF_Condition_Effect struct {
	right_saturation uint16
	left_saturation  uint16

	right_coeff int16
	left_coeff  int16

	deadband uint16
	center   int16
}

type FF_Periodic_Effect struct {
	waveform  uint16
	period    uint16
	magnitude int16
	offset    int16
	phase     uint16

	envelope FF_Envelope

	custom_len  uint32
	custom_data *int16
}

type FF_Rumble_Effect struct {
	strong_magnitude uint16
	weak_magnitude   uint16
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

const (
	FF_RUMBLE   = 0x50
	FF_PERIODIC = 0x51
	FF_CONSTANT = 0x52
	FF_SPRING   = 0x53
	FF_FRICTION = 0x54
	FF_DAMPER   = 0x55
	FF_INERTIA  = 0x56
	FF_RAMP     = 0x57

	FF_EFFECT_MIN = FF_RUMBLE
	FF_EFFECT_MAX = FF_RAMP

	FF_SQUARE   = 0x58
	FF_TRIANGLE = 0x59
	FF_SINE     = 0x5a
	FF_SAW_UP   = 0x5b
	FF_SAW_DOWN = 0x5c
	FF_CUSTOM   = 0x5d

	EVIOCSFF  = 0x40304580
	EVIOCRMFF = 0x40044581
)

func (d *Device) TestEffect() {
	fd, err := syscall.Open(d.dev_name, syscall.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer syscall.Close(fd)

	right(fd)
	left(fd)
}

func left(fd int) {
	ConstantEffect(fd, -1, 2000)
	time.Sleep(1000 * time.Millisecond)
}
func right(fd int) {
	ConstantEffect(fd, 1, 2000)
	time.Sleep(1000 * time.Millisecond)
}

func binPrint(ef FF_Effect) {
	var effectBytes bytes.Buffer
	binary.Write(&effectBytes, binary.LittleEndian, ef)

	fmt.Println(len(effectBytes.Bytes()))
	fmt.Println(effectBytes.Bytes())
}

var effect FF_Effect

func ConstantEffect(
	fd int,
	strength float32,
	duration uint16,
) {
	// erase effect
	if effect.id != -1 {
		if errno := eraseEffect(fd, &effect.id); errno != 0 {
			fmt.Println(errno)
		}
	}

	// set new values
	{
		effect.etype = FF_CONSTANT
		effect.id = -1
		effect.trigger.button = 0
		effect.trigger.interval = 0
		effect.replay.length = 0xffff
		effect.replay.delay = 0
		effect.direction = 0xC000

		constantEffect := FF_Constant_Effect{
			level: int16(strength * 32767),
			envelope: FF_Envelope{
				attack_length: 0,
				attack_level:  uint16(strength * 32767),
				fade_length:   0,
				fade_level:    uint16(strength * 32767),
			},
		}

		constantBytes := bytes.Buffer{}
		binary.Write(&constantBytes, binary.LittleEndian, constantEffect)
		copy(effect.u[:], constantBytes.Bytes())
	}

	binPrint(effect)

	// write effect
	errno := writeEffect(fd, &effect)
	if errno != 0 {
		fmt.Println(errno)
	}

	// start effect
	now := time.Now()
	err := writeInputEvent(fd, InputEvent{
		Sec:   now.Unix(),
		Usec:  int64(now.Nanosecond() / 1000),
		Type:  EV_FF,
		Code:  uint16(effect.id),
		Value: 1,
	})

	if err != nil {
		fmt.Println(err)
	}
}

func eraseEffect(
	fd int,
	effectId *int16,
) syscall.Errno {
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		EVIOCRMFF,
		uintptr(*effectId),
	)

	return errno
}

func writeEffect(
	fd int,
	effect *FF_Effect,
) syscall.Errno {
	_, _, errno := syscall.Syscall(
		syscall.SYS_IOCTL,
		uintptr(fd),
		EVIOCSFF,
		uintptr(unsafe.Pointer(effect)),
	)

	return errno
}

func writeInputEvent(
	fd int,
	evt InputEvent,
) error {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, evt); err != nil {
		return err
	}

	_, err := syscall.Write(fd, buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
