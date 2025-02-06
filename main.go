package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/jochenvg/go-udev"
)

type InputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

type model struct {
	dev      *udev.Device
	dev_path string
	dev_name string
}

const FF_AUTOCENTER = 0x61
const EV_FF = 0x15

const vendorID = 0x046d
const productID = 0xc24f

func init_model() model {
	u := udev.Udev{}

	enumerator := u.NewEnumerate()

	devices, err := enumerator.Devices()
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
		return model{}
	}

	path := fmt.Sprintf("/sys%s/device/device", dev.Properties()["DEVPATH"])

	devname := dev.Properties()["DEVNAME"]
	m := model{dev: dev, dev_path: path, dev_name: devname}

	return m
}

func main() {
	m := init_model()

	go m.print_events()

	m.set_autocenter(20000)
	m.set_range(500)

	select {}
}

func (m *model) set_range(value int) {
	file, err := os.OpenFile(fmt.Sprintf("%s/range", m.dev_path), os.O_WRONLY, 0666)
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

func (m *model) print_events() {
	for true {
		var data []byte = make([]byte, 300)

		fd, err := syscall.Open(m.dev_name, syscall.O_RDONLY, 0644)
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

		fmt.Println(event)

		if event.Code == 0 {
			fmt.Println(event.Value)
		}

		time.Sleep(1 * time.Second)
	}
}

func (m *model) set_autocenter(value int32) {
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
	fmt.Println(report)

	fd, err := syscall.Open(m.dev_name, syscall.O_WRONLY, 0644)
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
