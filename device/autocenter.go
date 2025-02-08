package device

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"time"
)

const (
	FF_AUTOCENTER = 0x61
	EV_FF         = 0x15
)

func (d *Device) SetAutocenter(perc int) {
	value := int32((float32(perc) / 100.0) * 65535)

	now := time.Now()
	evt := InputEvent{
		Sec:   now.Unix(),
		Usec:  int64(now.Nanosecond() / 1000),
		Type:  EV_FF,
		Code:  FF_AUTOCENTER,
		Value: value,
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, evt); err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}

	fd, err := syscall.Open(d.dev_name, syscall.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer syscall.Close(fd)

	_, err = syscall.Write(fd, buf.Bytes())
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
