package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"syscall"
	"time"
)

type InputEvent struct {
	Sec   int64
	Usec  int64
	Type  uint16
	Code  uint16
	Value int32
}

const FF_AUTOCENTER = 0x61;
const EV_FF = 0x15;

func main() {
	fmt.Printf("hi\n")
	// const vendorID = 0x046d
	// const productID = 0xc24f

    // u := udev.Udev{}

    // enumerator := u.NewEnumerate()

    // devices, err := enumerator.Devices()
    // if err != nil {
    //     log.Fatalf("Error scanning devices: %v", err)
	// }
    

	// var dev *udev.Device

    // for _, device := range devices {
	// 	props := device.Properties()
	// 	product_id, _ := strconv.ParseInt(props["ID_MODEL_ID"], 16, 64)
	// 	vendor_id, _ := strconv.ParseInt(props["ID_VENDOR_ID"], 16, 64)

	// 	if !strings.Contains(props["DEVNAME"], "event") {
	// 		continue
	// 	}
		
	// 	if product_id == productID && vendor_id == vendorID {
	// 		dev = device
	// 		break;
	// 	}
    // }



	// fmt.Println(dev.Properties())

	// fmt.Println()
	// fmt.Println()

	// path := fmt.Sprintf("/sys/%s/device/device", dev.Properties()["DEVPATH"])
	// fmt.Println("path:", path)


    // file, err := os.OpenFile(fmt.Sprintf("%s/range", path), os.O_WRONLY, 0666)
    // if err != nil {
    //     fmt.Println("Error opening file:", err)
    //     return
    // }
    // defer file.Close()

    // _, err = file.WriteString("900")
    // if err != nil {
    //     fmt.Println("Error writing to file:", err)
    //     return
    // }




	for true {
		var data []byte = make([]byte, 300)

		fd, err := syscall.Open("/dev/input/event13", syscall.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Read(fd, data)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)


		var event InputEvent

		reader := bytes.NewReader(data)

		err = binary.Read(reader, binary.LittleEndian, &event)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println(event)

		if event.Code == 0 {
			fmt.Println(event.Value)
		}


		time.Sleep(1 * time.Second)
	}







	// var autocenter_value int32 = 5315

	// now := time.Now()
	// ev := InputEvent{
	// 	Sec:   now.Unix(),
	// 	Usec:  int64(now.Nanosecond() / 1000),
	// 	Type:  EV_FF,
	// 	Code:  FF_AUTOCENTER,
	// 	Value: autocenter_value,
	// }

	// var buf bytes.Buffer
	// if err := binary.Write(&buf, binary.LittleEndian, ev); err != nil {
	// 	log.Fatalf("binary.Write failed: %v", err)
	// }
	// report := buf.Bytes()
	// fmt.Println(report)

	// fd, err := syscall.Open("/dev/input/event13", syscall.O_WRONLY, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// _, err = syscall.Write(fd, report)
	// if err != nil {
	// 	panic(err)
	// }
}
