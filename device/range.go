package device

import (
	"fmt"
	"strconv"
	"syscall"
)

func (d *Device) SetRange(value int) {
	fd, err := syscall.Open(fmt.Sprintf("%s/range", d.dev_path), syscall.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer syscall.Close(fd)

	_, err = syscall.Write(fd, []byte(strconv.Itoa(value)))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (d *Device) GetRange() int {
	fd, err := syscall.Open(fmt.Sprintf("%s/range", d.dev_path), syscall.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer syscall.Close(fd)

	b := make([]byte, 8)
	_, err = syscall.Read(fd, b)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return 0
	}

	var r int = 0

	for i := 0; b[i] != '\n' && i < len(b); i++ {
		if i > 0 {
			r *= 10
		}
		r += int(b[i] - '0')
	}

	return r
}
