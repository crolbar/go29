package virtDev

/*
#include <linux/uinput.h>
*/
import "C"

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

func KeyNew() error {
	var usetup uinputSetup

	fd, err := syscall.Open("/dev/uinput", syscall.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	ioctl(fd, C.UI_SET_EVBIT, C.EV_KEY)
	ioctl(fd, C.UI_SET_KEYBIT, C.KEY_A)

	ioctl(fd, C.UI_SET_EVBIT, C.EV_KEY)
	ioctl(fd, C.UI_SET_KEYBIT, C.BTN_LEFT)

	ioctl(fd, C.UI_SET_EVBIT, C.EV_REL)
	ioctl(fd, C.UI_SET_RELBIT, C.REL_X)
	ioctl(fd, C.UI_SET_RELBIT, C.REL_Y)

	usetup.id.bustype = C.BUS_USB
	usetup.id.vendor = 0
	usetup.id.product = 0
	copy(usetup.name[:], "go29-keeeb")

	ioctl(fd, C.UI_DEV_SETUP, uintptr(unsafe.Pointer(&usetup)))
	ioctl(fd, C.UI_DEV_CREATE, 0)

	time.Sleep(time.Second)

	for i := 50; i >= 0; i-- {
		emit(fd, C.EV_REL, C.REL_X, 50)
		emit(fd, C.EV_REL, C.REL_Y, 50)
		emit(fd, C.EV_SYN, C.SYN_REPORT, 0)
		time.Sleep(1500 * time.Microsecond)
	}

	emit(fd, C.EV_KEY, C.KEY_A, 1)
	emit(fd, C.EV_SYN, C.SYN_REPORT, 0)
	emit(fd, C.EV_KEY, C.KEY_A, 0)
	emit(fd, C.EV_SYN, C.SYN_REPORT, 0)

	time.Sleep(time.Second)

	ioctl(fd, C.UI_DEV_DESTROY, 0)
	syscall.Close(fd)

	fmt.Println("hi")
	return nil
}
