package virtDev

/*
#include <linux/uinput.h>
#include <linux/input-event-codes.h>
*/
import "C"

import (
	"go29/device"
	ec "go29/event_codes"
	"syscall"
	"time"
	"unsafe"
)

type VirtKeyboard struct {
	fd     int
	remaps map[wheelKey][]kbKey
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

func (vk *VirtKeyboard) ReloadConfig() error {
	remaps, err := ParseRemapConfig()
	if err != nil {
		return err
	}

	vk.remaps = remaps

	return nil
}

func (vk *VirtKeyboard) HandleInputEvent(evt device.InputEvent) {
	to, exists := vk.remaps[wheelKey(evt.Code)]
	if !exists {
		return
	}

	postActions := make([]func(), 0)

	outOfDeadZone := evt.Value >= 1

	switch evt.Code {
	case ec.ABS_WHEEL: // max 65535, med 32767
		var (
			deadzone = int32(1200)
			left     = evt.Value < 32767-deadzone
			right    = evt.Value > 32767+deadzone
		)

		if left {
			to = to[:len(to)/2]
		}
		if right {
			to = to[len(to)/2:]
		}

		outOfDeadZone = left || right
	case ec.ABS_BREAK, ec.ABS_THROTTLE, ec.ABS_CLUTCH: // start 255, end 0
		outOfDeadZone = 255-evt.Value >= 70
	}

	for _, k := range to {
		// click
		if k.click && outOfDeadZone {
			vk.ClickKey(k.key)
			continue
		}

		// press/release
		switch {
		case !outOfDeadZone:
			vk.ReleaseKey(k.key)
		case outOfDeadZone:
			vk.PressKey(k.key)
		}

		// release after all
		if k.modifier && outOfDeadZone {
			postActions = append(postActions,
				func() {
					vk.ReleaseKey(k.key)
				})
		}
	}

	for _, a := range postActions {
		a()
	}
}

func (k *VirtKeyboard) DestroyDev() {
	ioctl(k.fd, C.UI_DEV_DESTROY, 0)
	syscall.Close(k.fd)
}

func (k *VirtKeyboard) PressKey(key key) {
	emit(k.fd, C.EV_KEY, int(key), 1)
	syncEvents(k.fd)
}
func (k *VirtKeyboard) ReleaseKey(key key) {
	emit(k.fd, C.EV_KEY, int(key), 0)
	syncEvents(k.fd)
}
func (k *VirtKeyboard) ClickKey(key key) {
	emit(k.fd, C.EV_KEY, int(key), 1)
	syncEvents(k.fd)
	emit(k.fd, C.EV_KEY, int(key), 0)
	syncEvents(k.fd)
}
