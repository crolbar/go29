package virtDev

/*
#include <linux/uinput.h>
*/
import "C"
import (
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
	defer syscall.Close(fd)

	{
		// Enable event types
		ioctl(fd, C.UI_SET_EVBIT, C.EV_KEY)       // Buttons
		ioctl(fd, C.UI_SET_EVBIT, C.EV_ABS)       // Axes
		ioctl(fd, C.UI_SET_EVBIT, C.EV_FF)        // Force feedback
		ioctl(fd, C.UI_SET_EVBIT, C.EV_FF_STATUS) // Force feedback status

		// Configure buttons - G29 has many buttons
		// Main buttons
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_SOUTH)  // X
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_EAST)   // Circle
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_NORTH)  // Triangle
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_WEST)   // Square
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_START)  // Options
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_SELECT) // Share
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_MODE)   // PS button

		// Wheel buttons
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_GEAR_DOWN) // Left paddle
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_GEAR_UP)   // Right paddle
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TL)        // L2
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_TR)        // R2

		// Additional G29 buttons
		for i := C.BTN_0; i <= C.BTN_9; i++ {
			ioctl(fd, C.UI_SET_KEYBIT, uintptr(i))
		}

		// Wheel base buttons
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_THUMBL)
		ioctl(fd, C.UI_SET_KEYBIT, C.BTN_THUMBR)

		// Configure axes
		// Main controls
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_X)  // Steering wheel
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_Y)  // Throttle
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_Z)  // Brake
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_RZ) // Clutch

		// D-pad
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_HAT0X) // D-pad X
		ioctl(fd, C.UI_SET_ABSBIT, C.ABS_HAT0Y) // D-pad Y

		// Force feedback effects
		ioctl(fd, C.UI_SET_FFBIT, C.FF_CONSTANT)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_RAMP)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_PERIODIC)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_SPRING)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_FRICTION)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_DAMPER)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_RUMBLE)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_INERTIA)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_CUSTOM)

		// Periodic effect types
		ioctl(fd, C.UI_SET_FFBIT, C.FF_SQUARE)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_TRIANGLE)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_SINE)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_SAW_UP)
		ioctl(fd, C.UI_SET_FFBIT, C.FF_SAW_DOWN)
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

	log.Println("Virtual G29 device created successfully.")
	log.Println("Starting event handling loop...")

	emit(fd, C.EV_ABS, C.ABS_X, 0)     // Centered steering
	emit(fd, C.EV_ABS, C.ABS_Y, 0)     // Released throttle
	emit(fd, C.EV_ABS, C.ABS_Z, 0)     // Released brake
	emit(fd, C.EV_ABS, C.ABS_RZ, 0)    // Released clutch
	emit(fd, C.EV_ABS, C.ABS_HAT0X, 0) // Centered D-pad X
	emit(fd, C.EV_ABS, C.ABS_HAT0Y, 0) // Centered D-pad Y
	syncEvents(fd)

	{
		go func() {
			for {
				var stat syscall.Stat_t
				err := syscall.Fstat(fd, &stat)
				if err != nil {
					log.Printf("Fstat error: %v", err)
				} else {
					log.Printf("File descriptor %d status: dev=%d, ino=%d", fd, stat.Dev, stat.Ino)
				}
				time.Sleep(5 * time.Second)
			}
		}()
	}

	var buffer [64]byte
	for {

		// Check for force feedback effects sent from the system
		n, err := syscall.Read(fd, buffer[:])
		if err != nil {
			if err == syscall.EAGAIN {
				// No data available, not an error in non-blocking mode
				time.Sleep(10 * time.Millisecond)
				continue
			}
			log.Printf("Error reading from device: %v", err)
			time.Sleep(100 * time.Millisecond)
			continue
		}

		if n > 0 {
			// Process the force feedback effect data
			// This would depend on how you want to handle FF effects
			log.Printf("Received %d bytes of force feedback data", n)

			// Here you would typically:
			// 1. Parse the effect data
			// 2. Acknowledge receipt of the effect
			// 3. Simulate the effect on your virtual hardware or pass it to real hardware
		}

		// You can add here code to simulate steering wheel movement, button presses, etc.
		// For example, to simulate oscillating steering:
		/*
			oscillation := int32(32767.0 * math.Sin(float64(time.Now().UnixNano())/1e9))
			writeEvent(fd, C.EV_ABS, C.ABS_X, oscillation)
			syncEvents(fd)
		*/

		time.Sleep(10 * time.Millisecond)
	}

	// time.Sleep(time.Second)

	// ioctl(fd, C.UI_DEV_DESTROY, 0)
	// syscall.Close(fd)

	return nil
}
