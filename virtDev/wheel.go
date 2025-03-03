package virtDev

/*
#include <linux/uinput.h>
*/
import "C"
import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"
)

func WheelNew() error {
	fd, err := syscall.Open("/dev/uinput", syscall.O_WRONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return err
	}

	{
		// EV
		{
			ioctl(fd, C.UI_SET_EVBIT, C.EV_SYN)
			ioctl(fd, C.UI_SET_EVBIT, C.EV_KEY) // Buttons
			ioctl(fd, C.UI_SET_EVBIT, C.EV_ABS) // Axes
			ioctl(fd, C.UI_SET_EVBIT, C.EV_FF)  // Force feedback
			ioctl(fd, C.UI_SET_EVBIT, C.EV_MSC)
		}

		// KEY
		{
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_THUMB2)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_THUMB)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TOP)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE5)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE2)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE6)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY8)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY6)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY7)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE3)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_BASE4)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY9)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY5)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY4)

			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TOP2)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_PINKIE)

			ioctl(fd, C.UI_SET_KEYBIT, 0x12C)
			ioctl(fd, C.UI_SET_KEYBIT, 0x12D)
			ioctl(fd, C.UI_SET_KEYBIT, 0x12E)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_DEAD)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY1)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY2)
			ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TRIGGER_HAPPY3)
		}

		// MSC
		{
			ioctl(fd, C.UI_SET_MSCBIT, C.MSC_SCAN)
		}
		// ABS
		{
			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_X)  // Steering wheel
			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_Y)  // Clutch
			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_Z)  // Throttle
			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_RZ) // Break

			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_HAT0X) // D-pad X
			ioctl(fd, C.UI_SET_ABSBIT, C.ABS_HAT0Y) // D-pad Y
		}

		// FF
		{
			ioctl(fd, C.UI_SET_FFBIT, C.FF_CONSTANT)
			ioctl(fd, C.UI_SET_FFBIT, C.FF_GAIN)
			ioctl(fd, C.UI_SET_FFBIT, C.FF_AUTOCENTER)
		}
	}

	var uidev uinputUserDev

	uidev.id.bustype = C.BUS_USB
	uidev.id.vendor = 0x046d
	uidev.id.product = 0xc24f
	uidev.id.version = 0x0111
	copy(uidev.name[:], "Logitech G29 Driving Force Racing Wheel")

	uidev.absmin[C.ABS_X] = -32767
	uidev.absmax[C.ABS_X] = 32767
	uidev.absfuzz[C.ABS_X] = 0
	uidev.absflat[C.ABS_X] = 0

	// Throttle: 0 to 255
	uidev.absmin[C.ABS_Y] = 0
	uidev.absmax[C.ABS_Y] = 255
	uidev.absfuzz[C.ABS_Y] = 0
	uidev.absflat[C.ABS_Y] = 0

	// Brake: 0 to 255
	uidev.absmin[C.ABS_Z] = 0
	uidev.absmax[C.ABS_Z] = 255
	uidev.absfuzz[C.ABS_Z] = 0
	uidev.absflat[C.ABS_Z] = 0

	// Clutch: 0 to 255
	uidev.absmin[C.ABS_RZ] = 0
	uidev.absmax[C.ABS_RZ] = 255
	uidev.absfuzz[C.ABS_RZ] = 0
	uidev.absflat[C.ABS_RZ] = 0

	// D-pad X: -1 to 1
	uidev.absmin[C.ABS_HAT0X] = -1
	uidev.absmax[C.ABS_HAT0X] = 1
	uidev.absfuzz[C.ABS_HAT0X] = 0
	uidev.absflat[C.ABS_HAT0X] = 0

	// D-pad Y: -1 to 1
	uidev.absmin[C.ABS_HAT0Y] = -1
	uidev.absmax[C.ABS_HAT0Y] = 1
	uidev.absfuzz[C.ABS_HAT0Y] = 0
	uidev.absflat[C.ABS_HAT0Y] = 0

	uidev.ff_effects_max = 24

	err = ioctl(fd, C.UI_DEV_SETUP, uintptr(unsafe.Pointer(&uidev)))
	if err != nil {
		log.Fatal("SETUP", err)
	}
	err = ioctl(fd, C.UI_DEV_CREATE, 0)
	if err != nil {
		log.Fatal("CREATE", err)
	}

	time.Sleep(time.Second)

	emit(fd, C.EV_ABS, C.ABS_X, 0)     // Centered steering
	emit(fd, C.EV_ABS, C.ABS_Y, 0)     // Released throttle
	emit(fd, C.EV_ABS, C.ABS_Z, 0)     // Released brake
	emit(fd, C.EV_ABS, C.ABS_RZ, 0)    // Released clutch
	emit(fd, C.EV_ABS, C.ABS_HAT0X, 0) // Centered D-pad X
	emit(fd, C.EV_ABS, C.ABS_HAT0Y, 0) // Centered D-pad Y
	syncEvents(fd)

	log.Println("Virtual G29 device created successfully.")
	log.Println("Starting event handling loop...")

	fdr, err := syscall.Open("/dev/input/event13", syscall.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	{
		go func() {
			for {
				var stat syscall.Stat_t
				err := syscall.Fstat(fdr, &stat)
				if err != nil {
					log.Printf("Fstat error: %v", err)
				} else {
					log.Printf("File descriptor %d status: dev=%d, ino=%d", fdr, stat.Dev, stat.Ino)
				}
				time.Sleep(5 * time.Second)
			}
		}()
	}

	var buffer []byte = make([]byte, 24*64)
	for {
		n, err := syscall.Read(fdr, buffer)
		if err != nil {
			if err == syscall.EAGAIN {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			log.Printf("Error reading from device: %v", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if n > 0 {
			log.Printf("Received %d bytes of force feedback data", n)
		}

		fmt.Println("sleeping")
		time.Sleep(10 * time.Millisecond)
	}

	// ioctl(fd, C.UI_DEV_DESTROY, 0)
	// syscall.Close(fd)

	return nil
}
