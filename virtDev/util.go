package virtDev

/*
#include <linux/uinput.h>
*/
import "C"
import (
	"bytes"
	"encoding/binary"
	"syscall"
)

func ioctl(fd int, cmd, arg uintptr) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), cmd, arg)
	if errno != 0 {
		return errno
	}
	return nil
}

func emit(fd int, _type int, code int, val int) error {
	var evt inputEvent

	evt.Type = uint16(_type)
	evt.Code = uint16(code)
	evt.Value = int32(val)

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, evt); err != nil {
		// fmt.Println("binary.Write failed:", err)
		return err
	}

	syscall.Write(fd, buf.Bytes())
	
	return nil
}

func syncEvents(fd int) error {
	return emit(fd, C.EV_SYN, C.SYN_REPORT, 0)
}
