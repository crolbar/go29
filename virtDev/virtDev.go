package virtDev

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"fmt"
// 	"os"
// 	"syscall"
// 	"time"
// )

// // IOCTL constants (values from typical Linux uinput headers)
// const (
// 	UI_SET_EVBIT   = 1074025828 // _IOW(UINPUT_IOCTL_BASE, 100, int)
// 	UI_SET_KEYBIT  = 1074025829 // _IOW(UINPUT_IOCTL_BASE, 101, int)
// 	UI_DEV_CREATE  = 21761      // _IO(UINPUT_IOCTL_BASE, 1)
// 	UI_DEV_DESTROY = 21762      // _IO(UINPUT_IOCTL_BASE, 2)
// )

// const (
// 	EV_KEY     = 0x01
// 	EV_SYN     = 0x00
// 	KEY_A      = 23
// 	SYN_REPORT = 0
// )

// // inputID corresponds to the struct input_id in <linux/input.h>
// type inputID struct {
// 	Bustype uint16
// 	Vendor  uint16
// 	Product uint16
// 	Version uint16
// }

// // uinputUserDev corresponds to the struct uinput_user_dev in <linux/uinput.h>
// type uinputUserDev struct {
// 	Name         [80]byte
// 	ID           inputID
// 	FFEffectsMax int32
// 	AbsMax       [64]int32
// 	AbsMin       [64]int32
// 	AbsFuzz      [64]int32
// 	AbsFlat      [64]int32
// }

// // inputEvent corresponds to the struct input_event in <linux/input.h>
// // On a 64-bit system this is typically 24 bytes.
// type inputEvent struct {
// 	TimeSec  int64  // tv_sec
// 	TimeUsec int64  // tv_usec
// 	Type     uint16 // event type
// 	Code     uint16 // event code
// 	Value    int32  // event value
// }

// type VirtDev struct {
// 	f *os.File
// }

// func NewVirtDev() *VirtDev {
// 	f, err := os.OpenFile("/dev/uinput", os.O_WRONLY, 0)
// 	if err != nil {
// 		fmt.Printf("Error opening /dev/uinput: %v\n", err)
// 		return nil
// 	}
// 	// defer f.Close()

// 	vd := VirtDev{f: f}

// 	if err := vd.enableKeys(); err != nil {
// 		return nil
// 	}

// 	if err := vd.setupDev(); err != nil {
// 		return nil
// 	}

// 	return &vd
// }

// // sendSynEvent writes an EV_SYN event to signal event completion.
// func (vd *VirtDev) sendSynEvent() {
// 	var buf bytes.Buffer
// 	syn := inputEvent{
// 		TimeSec:  time.Now().Unix(),
// 		TimeUsec: int64(time.Now().Nanosecond() / 1000),
// 		Type:     EV_SYN,
// 		Code:     SYN_REPORT,
// 		Value:    0,
// 	}
// 	if err := binary.Write(&buf, binary.LittleEndian, syn); err != nil {
// 		fmt.Printf("Error encoding SYN event: %v\n", err)
// 		return
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing SYN event: %v\n", err)
// 		return
// 	}
// }

// func (vd *VirtDev) enableKeys() error {
// 	// Enable EV_KEY events.
// 	if err := ioctl(vd.f.Fd(), UI_SET_EVBIT, EV_KEY); err != nil {
// 		fmt.Printf("Error setting EV_KEY: %v\n", err)
// 		return err
// 	}
// 	// // Enable KEY_A.
// 	if err := ioctl(vd.f.Fd(), UI_SET_KEYBIT, KEY_A); err != nil {
// 		fmt.Printf("Error setting KEY_A: %v\n", err)
// 		return err
// 	}
// 	return nil
// }

// func (vd *VirtDev) setupDev() error {
// 	var uidev uinputUserDev
// 	copy(uidev.Name[:], "go-uinput-keyboard")

// 	uidev.ID.Bustype = 0x03
// 	uidev.ID.Vendor = 0x1
// 	uidev.ID.Product = 0x1
// 	uidev.ID.Version = 1

// 	// Write the uinput_user_dev structure to the file descriptor.
// 	var buf bytes.Buffer
// 	if err := binary.Write(&buf, binary.LittleEndian, uidev); err != nil {
// 		fmt.Printf("Error encoding uinput_user_dev: %v\n", err)
// 		return err
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing uinput_user_dev: %v\n", err)
// 		return err
// 	}

// 	// Create the device.
// 	if err := ioctl(vd.f.Fd(), UI_DEV_CREATE, 0); err != nil {
// 		fmt.Printf("Error creating device: %v\n", err)
// 		return err
// 	}

// 	// Allow some time for the device to be registered.
// 	time.Sleep(time.Second)

// 	return nil
// }

// func (vd *VirtDev) DestroyDev() error {
// 	if err := ioctl(vd.f.Fd(), UI_DEV_DESTROY, 0); err != nil {
// 		fmt.Printf("Error destroying device: %v\n", err)
// 		return err
// 	}
// 	vd.f.Close()

// 	return nil
// }

// func ioctl(fd uintptr, cmd, arg uintptr) error {
// 	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, cmd, arg)
// 	if errno != 0 {
// 		return errno
// 	}
// 	return nil
// }

// func (vd *VirtDev) PressA() {
// 	var buf bytes.Buffer
// 	// Send a key press for KEY_A.
// 	keyDown := inputEvent{
// 		TimeSec:  time.Now().Unix(),
// 		TimeUsec: int64(time.Now().Nanosecond() / 1000),
// 		Type:     EV_KEY,
// 		Code:     KEY_A,
// 		Value:    1, // Key press
// 	}
// 	buf.Reset()
// 	if err := binary.Write(&buf, binary.LittleEndian, keyDown); err != nil {
// 		fmt.Printf("Error encoding keyDown: %v\n", err)
// 		return
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing keyDown: %v\n", err)
// 		return
// 	}

// 	// Send a SYN event to mark the end of this key event packet.
// 	vd.sendSynEvent()
// }

// func (vd *VirtDev) ReleaseA() {
// 	var buf bytes.Buffer
// 	// Send a key release for KEY_A.
// 	keyUp := inputEvent{
// 		TimeSec:  time.Now().Unix(),
// 		TimeUsec: int64(time.Now().Nanosecond() / 1000),
// 		Type:     EV_KEY,
// 		Code:     KEY_A,
// 		Value:    0, // Key release
// 	}
// 	buf.Reset()
// 	if err := binary.Write(&buf, binary.LittleEndian, keyUp); err != nil {
// 		fmt.Printf("Error encoding keyUp: %v\n", err)
// 		return
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing keyUp: %v\n", err)
// 		return
// 	}

// 	// Send another SYN event.
// 	vd.sendSynEvent()
// }

// func (vd *VirtDev) SendA() {
// 	var buf bytes.Buffer
// 	// Send a key press for KEY_A.
// 	keyDown := inputEvent{
// 		TimeSec:  time.Now().Unix(),
// 		TimeUsec: int64(time.Now().Nanosecond() / 1000),
// 		Type:     EV_KEY,
// 		Code:     KEY_A,
// 		Value:    1, // Key press
// 	}
// 	buf.Reset()
// 	if err := binary.Write(&buf, binary.LittleEndian, keyDown); err != nil {
// 		fmt.Printf("Error encoding keyDown: %v\n", err)
// 		return
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing keyDown: %v\n", err)
// 		return
// 	}

// 	// Send a SYN event to mark the end of this key event packet.
// 	vd.sendSynEvent()

// 	// Wait a short while before releasing the key.
// 	time.Sleep(100 * time.Millisecond)

// 	// Send a key release for KEY_A.
// 	keyUp := inputEvent{
// 		TimeSec:  time.Now().Unix(),
// 		TimeUsec: int64(time.Now().Nanosecond() / 1000),
// 		Type:     EV_KEY,
// 		Code:     KEY_A,
// 		Value:    0, // Key release
// 	}
// 	buf.Reset()
// 	if err := binary.Write(&buf, binary.LittleEndian, keyUp); err != nil {
// 		fmt.Printf("Error encoding keyUp: %v\n", err)
// 		return
// 	}
// 	if _, err := vd.f.Write(buf.Bytes()); err != nil {
// 		fmt.Printf("Error writing keyUp: %v\n", err)
// 		return
// 	}

// 	// Send another SYN event.
// 	vd.sendSynEvent()

// 	fmt.Println("Keyboard event (KEY_A) sent via virtual device.")
// }
