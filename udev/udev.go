package udev

/*
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
*/
import "C"
import (
	"errors"
	"fmt"
)

type Udev struct {
	ptr *C.struct_udev
}

type Enumerate struct {
	ptr *C.struct_udev_enumerate
	u *Udev
}

type Device struct {
	ptr *C.struct_udev_device
}

func newDevice(ptr *C.struct_udev_device) (d *Device) {
	if ptr == nil {
		return nil
	}

	d = &Device{
		ptr: ptr,
	}

	return d
}

func (d *Device) Properties() (r map[string]string) {
	r = make(map[string]string)
	for l := C.udev_device_get_properties_list_entry(d.ptr); l != nil; l = C.udev_list_entry_get_next(l) {
		r[C.GoString(C.udev_list_entry_get_name(l))] = C.GoString(C.udev_list_entry_get_value(l))
	}

	return r
}

func NewUdev() Udev {
	u := Udev{}
	u.ptr = C.udev_new()
	return u
}

func (u *Udev) NewEnumerate() Enumerate {
	e := Enumerate{}
	e.ptr = C.udev_enumerate_new(u.ptr)
	e.u = u
	return e
}

func (e *Enumerate) GetDevices() ([]*Device, error) {
	k := C.udev_enumerate_scan_devices(e.ptr)
	if k < 0 {
		return nil, errors.New(fmt.Sprintf("error K: %d", k))
	}

	m := make([]*Device, 0)
	for l := C.udev_enumerate_get_list_entry(e.ptr); l != nil; l = C.udev_list_entry_get_next(l) {
		s := C.udev_list_entry_get_name(l)
		m = append(m, newDevice(C.udev_device_new_from_syspath(e.u.ptr, s)))
	}

	return m, nil
}
