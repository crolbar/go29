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
	fd int
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

	return &VirtKeyboard{
		fd: fd,
	}, nil
}

func (k *VirtKeyboard) HandleInputEvent(evt device.InputEvent) {
	if evt.Code == C.BTN_TRIGGER && evt.Value == 1 {
		k.PressKey(KEY_A)
	}
	
	if evt.Code == C.BTN_TRIGGER && evt.Value == 0 {
		k.ReleaseKey(KEY_A)
	}
}

func (k *VirtKeyboard) DestroyDev() {
	ioctl(k.fd, C.UI_DEV_DESTROY, 0)
	syscall.Close(k.fd)
}

func (k *VirtKeyboard) PressKey(key int) {
	emit(k.fd, C.EV_KEY, key, 1)
	syncEvents(k.fd)
}

func (k *VirtKeyboard) ReleaseKey(key int) {
	emit(k.fd, C.EV_KEY, key, 0)
	syncEvents(k.fd)
}

func (k *VirtKeyboard) ClickKey(key int) {
	emit(k.fd, C.EV_KEY, key, 1)
	syncEvents(k.fd)
	emit(k.fd, C.EV_KEY, key, 0)
	syncEvents(k.fd)
}
