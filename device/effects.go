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
	FF_CONSTANT = 0x52
	EVIOCSFF    = 0x40304580
	EVIOCRMFF   = 0x40044581
)

func (d *Device) SetConstantEffect(strength float32) {
	// erase effect
	if d.effect.id != -1 {
		if errno := eraseEffect(d.Fd, &d.effect.id); errno != 0 {
			fmt.Println(errno)
		}
	}

	// set new values
	{
		d.effect.etype = FF_CONSTANT
		d.effect.id = -1
		d.effect.direction = 0xC000

		constantEffect := FF_Constant_Effect{
			level:    int16(strength * 0x7fff),
			envelope: FF_Envelope{},
		}

		constantBytes := bytes.Buffer{}
		binary.Write(&constantBytes, binary.LittleEndian, constantEffect)
		copy(d.effect.u[:], constantBytes.Bytes())
	}

	// write effect
	errno := writeEffect(d.Fd, d.effect)
	if errno != 0 {
		fmt.Println(errno)
	}

	// start effect
	now := time.Now()
	err := writeInputEvent(d.Fd, InputEvent{
		Sec:   now.Unix(),
		Usec:  int64(now.Nanosecond() / 1000),
		Type:  EV_FF,
		Code:  uint16(d.effect.id),
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
