package virtDev

/*
#include <linux/uinput.h>
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go29/device"
	"log"
	"syscall"
	"time"
	"unsafe"
)

func WheelNew() error {
	fd, err := setupWheel()
	if err != nil {
		return err
	}

	log.Println("Virtual G29 device created successfully.")
	log.Println("Starting event handling loop...")

	fdr, err := syscall.Open("/dev/input/event13", syscall.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	// {
	// 	go func() {
	// 		for {
	// 			// emit(fd, C.EV_MSC, C.MSC_SCAN, 589827)
	// 			// 			emit(fd, C.EV_KEY, C.BTN_THUMB2, 1)
	// 			// 			syncEvents(fd)

	// 			// 			time.Sleep(50 * time.Millisecond)

	// 			// 			emit(fd, C.EV_MSC, C.MSC_SCAN, 589827)
	// 			// 			emit(fd, C.EV_KEY, C.BTN_THUMB2, 0)

	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			// emit(fd, C.EV_FF, C.FF_STATUS_STOPPED, 1)
	// 			syncEvents(fd)

	// 			time.Sleep(100 * time.Millisecond)
	// 		}
	// 	}()
	// }

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

	var realFd int
	{
		d, err := device.NewDevice()
		if err != nil {
			fmt.Println("Error while creating device: ", err)
			return err
		}
		defer d.CloseFD()

		ch := make(chan device.InputEvents, 20)
		d.SpawnEventListenerThread(ch)
		go eventProxy(fd, ch)

		realFd = d.Fd
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
			// fmt.Println(buffer)
			printEvents(buffer, n, fd, realFd)
		}

		// fmt.Println("sleeping")
		time.Sleep(10 * time.Millisecond)
	}

	// ioctl(fd, C.UI_DEV_DESTROY, 0)
	// syscall.Close(fd)

	return nil
}

func eventProxy(fd int, ch chan device.InputEvents) {
	for {
		select {
		case events, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
			}

			data := events.Data
			n := events.N

			var evt device.InputEvent

			reader := bytes.NewReader(data[0:24])

			for l, r := 0, 24; r < n; l, r = r, r+24 {
				reader.Reset(data[l:r])

				err := binary.Read(reader, binary.LittleEndian, &evt)
				if err != nil {
					fmt.Println("binary.Read Error:", err)
				}

				emit(fd, int(evt.Type), int(evt.Code), int(evt.Value))

				// skip SYN & unused KEY events
				// if evt.Type == ec.EV_SYN || evt.Type == ec.EV_KEY {
				// 	continue
				// }

			}
			syncEvents(fd)
		}
	}
}

func printEvents(buf []byte, n int, fd int, realFd int) {
	numEvents := n / 24

	for i := 0; i < numEvents; i++ {
		// Calculate start index for this event
		start := i * 24

		// // Parse timestamp (first 8 bytes)
		// timestamp := binary.LittleEndian.Uint64(buf[start : start+8])

		// Parse event type (next 2 bytes)
		eventType := binary.LittleEndian.Uint16(buf[start+16 : start+18])

		// Parse event code (next 2 bytes)
		eventCode := binary.LittleEndian.Uint16(buf[start+18 : start+20])

		// Parse event value (next 4 bytes)
		eventValue := binary.LittleEndian.Uint32(buf[start+20 : start+24])

		// Print readable event information
		fmt.Printf("Event %d: ", i+1)
		// fmt.Printf("  Timestamp: %d\n", timestamp)
		fmt.Printf("  Type:  0x%04x ", eventType)

		// Translate known event types
		switch eventType {
		case 0x0000:
			fmt.Print("(EV_SYN)")
		case 0x0001:
			fmt.Print("(EV_KEY)")
		case 0x0002:
			fmt.Print("(EV_REL)")
		case 0x0003:
			fmt.Print("(EV_ABS)")
		case 0x0004:
			fmt.Print("(EV_MSC)")
		case 0x0015:
			fmt.Print("(EV_FF)")
		default:
			fmt.Print("(Unknown)")
		}

		fmt.Printf("  Code:  0x%04x ", eventCode)
		fmt.Printf("  Value: 0x%08x (%d)\n", eventValue, eventValue)
		// fmt.Println()

		if eventType == C.EV_FF && eventCode == C.FF_GAIN {
			emit(realFd, int(eventType), int(eventCode), int(eventValue))
		}
	}
}

func setupWheel() (fd int, err error) {
	fd, err = syscall.Open("/dev/uinput", syscall.O_WRONLY|syscall.O_NONBLOCK, 0644)
	if err != nil {
		return
	}

	// set supported events
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

	uidev.absmin[C.ABS_X] = 0
	uidev.absmax[C.ABS_X] = 65535
	uidev.absfuzz[C.ABS_X] = 0
	uidev.absflat[C.ABS_X] = 0

	// clutch: 0 to 255
	uidev.absmin[C.ABS_Y] = 0
	uidev.absmax[C.ABS_Y] = 255
	uidev.absfuzz[C.ABS_Y] = 0
	uidev.absflat[C.ABS_Y] = 0

	// throttle: 0 to 255
	uidev.absmin[C.ABS_Z] = 0
	uidev.absmax[C.ABS_Z] = 255
	uidev.absfuzz[C.ABS_Z] = 0
	uidev.absflat[C.ABS_Z] = 0

	// break: 0 to 255
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

	emit(fd, C.EV_ABS, C.ABS_X, 32767) // Centered steering
	emit(fd, C.EV_ABS, C.ABS_Y, 255)   // Released clutch
	emit(fd, C.EV_ABS, C.ABS_Z, 255)   // Released throttle
	emit(fd, C.EV_ABS, C.ABS_RZ, 255)  // Released break
	emit(fd, C.EV_ABS, C.ABS_HAT0Y, 0)
	emit(fd, C.EV_ABS, C.ABS_HAT0X, 0)
	syncEvents(fd)

	return
}
