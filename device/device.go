package device

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go29/event_codes"
	"go29/udev"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type InputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

const FF_AUTOCENTER = 0x61
const EV_FF = 0x15

const vendorID = 0x046d
const productID = 0xc24f

type Device struct {
	dev      *udev.Device
	dev_path string
	dev_name string

	p *tea.Program
}

type Send struct {
	Value int
}

type SendThrottle struct {
	Value int
}

type SendButton struct {
	Value int
}

type SendDpad struct {
	Value int
	Code  int
}

func (d *Device) SetProgram(p *tea.Program) {
	d.p = p
}

func NewDevice() Device {
	u := udev.NewUdev()

	enumerator := u.NewEnumerate()

	devices, err := enumerator.GetDevices()
	if err != nil {
		fmt.Println("error scanning devices:", err)
		panic("")
	}

	var dev *udev.Device

	for _, device := range devices {
		props := device.Properties()
		product_id, _ := strconv.ParseInt(props["ID_MODEL_ID"], 16, 64)
		vendor_id, _ := strconv.ParseInt(props["ID_VENDOR_ID"], 16, 64)

		if !strings.Contains(props["DEVNAME"], "event") {
			continue
		}

		if product_id == productID && vendor_id == vendorID {
			dev = device
			break
		}
	}

	// TODO
	if dev == nil {
		return Device{}
	}

	path := fmt.Sprintf("/sys%s/device/device", dev.Properties()["DEVPATH"])

	devname := dev.Properties()["DEVNAME"]
	m := Device{dev: dev, dev_path: path, dev_name: devname}

	return m
}

func (d *Device) SetRange(value int) {
	file, err := os.OpenFile(fmt.Sprintf("%s/range", d.dev_path), os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%d", value))
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (d *Device) GetRange() int {
	file, err := os.OpenFile(fmt.Sprintf("%s/range", d.dev_path), os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer file.Close()

	b := make([]byte, 24)
	_, err = file.Read(b)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return 0
	}

	var r int = 0

	for i := 0; b[i] != 10 && i < len(b); i++ {
		if i > 0 {
			r *= 10
		}
		r += int(b[i] - '0')
	}

	return r
}

func (d *Device) SetAutocenter(perc int) {
	value := int32((float32(perc) / 100.0) * 65535)
	now := time.Now()
	ev := InputEvent{
		Sec:   now.Unix(),
		Usec:  int64(now.Nanosecond() / 1000),
		Type:  EV_FF,
		Code:  FF_AUTOCENTER,
		Value: value,
	}

	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, ev); err != nil {
		fmt.Println("binary.Write failed:", err)
		return
	}
	report := buf.Bytes()

	fd, err := syscall.Open(d.dev_name, syscall.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	_, err = syscall.Write(fd, report)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}

func (d *Device) PrintEvents() {
	for true {
		var data []byte = make([]byte, 300)

		fd, err := syscall.Open(d.dev_name, syscall.O_RDONLY, 0644)
		if err != nil {
			fmt.Println("Error opening dev event", err)
		}

		_, err = syscall.Read(fd, data)
		if err != nil {
			fmt.Println("Error reading input_event:", err)
		}

		var event InputEvent
		reader := bytes.NewReader(data)
		err = binary.Read(reader, binary.LittleEndian, &event)
		if err != nil {
			fmt.Println("binary.Read Error:", err)
			return
		}

		// fmt.Println(event)

		switch event.Code {
		case event_codes.ABS_X:
			// fmt.Println(event.Value)
			d.p.Send(func() tea.Msg {
				return Send{Value: int(event.Value)}
			}())
		case event_codes.ABS_Z:
			d.p.Send(func() tea.Msg {
				return SendThrottle{Value: int(event.Value)}
			}())
		case event_codes.ABS_RY:
			d.p.Send(func() tea.Msg {
				return SendButton{Value: int(event.Value)}
			}())
		case event_codes.ABS_HAT0Y:
			fallthrough
		case event_codes.ABS_HAT0X:
			d.p.Send(func() tea.Msg {
				return SendDpad{Code: int(event.Code), Value: int(event.Value)}
			}())
		}

		time.Sleep(1 * time.Millisecond)
	}
}
