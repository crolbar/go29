package virtDev

/*
#include <linux/uinput.h>
#include <linux/input-event-codes.h>
*/
import "C"

import (
	"go29/device"
	"syscall"
	"time"
	"unsafe"
)

type VirtKeyboard struct {
	fd     int
	remaps []Remap
}

func NewVirtKeyboard() (*VirtKeyboard, error) {
	var usetup uinputSetup

	fd, err := syscall.Open("/dev/uinput", syscall.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	err = ioctl(fd, C.UI_SET_EVBIT, C.EV_KEY)
	if err != nil {
		return nil, err
	}

	for i := C.KEY_ESC; i <= C.KEY_MICMUTE; i++ {
		err = ioctl(fd, C.UI_SET_KEYBIT, uintptr(i))
		if err != nil {
			return nil, err
		}
	}

	usetup.id.bustype = C.BUS_USB
	usetup.id.vendor = 0
	usetup.id.product = 0
	copy(usetup.name[:], "go29-keeeb")

	err = ioctl(fd, C.UI_DEV_SETUP, uintptr(unsafe.Pointer(&usetup)))
	if err != nil {
		return nil, err
	}
	err = ioctl(fd, C.UI_DEV_CREATE, 0)
	if err != nil {
		return nil, err
	}

	time.Sleep(time.Second / 2)

	remaps, err := ParseRemapConfig()
	if err != nil {
		return nil, err
	}
	return &VirtKeyboard{
		fd:     fd,
		remaps: remaps,
	}, nil
}

func (k *VirtKeyboard) HandleInputEvent(evt device.InputEvent) {
	for _, r := range k.remaps {
		if evt.Code == uint16(r.from) {
			if r.modified && evt.Value == 1{
				k.PressKey(r.to[0])
				k.ClickKeys(r.to[1:])
				k.ReleaseKey(r.to[0])
				continue
			}

			if r.click && evt.Value == 1 {
				k.ClickKeys(r.to)
				continue
			}


			switch evt.Value {
			case 0:
				k.ReleaseKeys(r.to)
			case 1:
				k.PressKeys(r.to)
			}
		}
	}
}

func (k *VirtKeyboard) DestroyDev() {
	ioctl(k.fd, C.UI_DEV_DESTROY, 0)
	syscall.Close(k.fd)
}

func (k *VirtKeyboard) PressKeys(key []KBKey) {
	for _, _k := range key {
		k.PressKey(_k)
	}
}

func (k *VirtKeyboard) PressKey(key KBKey) {
	emit(k.fd, C.EV_KEY, int(key), 1)
	syncEvents(k.fd)
}

func (k *VirtKeyboard) ReleaseKeys(key []KBKey) {
	for _, _k := range key {
		k.ReleaseKey(_k)
	}
}

func (k *VirtKeyboard) ReleaseKey(key KBKey) {
	emit(k.fd, C.EV_KEY, int(key), 0)
	syncEvents(k.fd)
}

func (k *VirtKeyboard) ClickKeys(key []KBKey) {
	for _, _k := range key {
		k.ClickKey(_k)
	}
}

func (k *VirtKeyboard) ClickKey(key KBKey) {
	emit(k.fd, C.EV_KEY, int(key), 1)
	syncEvents(k.fd)
	emit(k.fd, C.EV_KEY, int(key), 0)
	syncEvents(k.fd)
}
