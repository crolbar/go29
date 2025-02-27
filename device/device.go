package device

import (
	"errors"
	"fmt"
	"go29/udev"
	"strconv"
	"strings"
)

const (
	vendorID  = 0x046d
	productID = 0xc24f
)

type InputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

type Device struct {
	dev      *udev.Device
	dev_path string
	dev_name string
}

func (d *Device) T() {
	fmt.Println(d.dev_name)
	fmt.Println(d.dev_path)

	d.TestEffect()
	// d.SetRange(540)
}


func NewDevice() (*Device, error) {
	u := udev.NewUdev()

	enumerator := u.NewEnumerate()

	devices, err := enumerator.GetDevices()
	if err != nil {
		return nil, err
	}

	var dev *udev.Device

	for _, device := range devices {
		props := device.Properties()

		if !strings.Contains(props["DEVNAME"], "event") {
			continue
		}

		product_id, _ := strconv.ParseInt(props["ID_MODEL_ID"], 16, 64)
		vendor_id, _ := strconv.ParseInt(props["ID_VENDOR_ID"], 16, 64)

		if product_id == productID && vendor_id == vendorID {
			dev = device
			break
		}
	}

	if dev == nil {
		return nil, errors.New("No device found with vendorID: 0x046d and productID: 0xc24f")
	}

	path := fmt.Sprintf("/sys%s/device/device", dev.Properties()["DEVPATH"])
	devname := dev.Properties()["DEVNAME"]
	m := Device{dev: dev, dev_path: path, dev_name: devname}

	return &m, nil
}
