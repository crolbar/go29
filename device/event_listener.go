package device

import (
	"fmt"
	"syscall"

	tea "github.com/charmbracelet/bubbletea"
)

type InputEvents struct {
	Data []byte
	N    int
}

func (d *Device) SpawnEventListenerThread(p *tea.Program) {
	fd, err := syscall.Open(d.dev_name, syscall.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("Error opening dev event", err)
	}

	go eventListener(fd, p)
}

func eventListener(fd int, p *tea.Program) {
	epfd, err := syscall.EpollCreate1(0)
	if err != nil {
		fmt.Println("Error creating epoll:", err)
		return
	}
	defer syscall.Close(epfd)

	ev := &syscall.EpollEvent{
		Events: syscall.EPOLLIN,
		Fd:     int32(fd),
	}
	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, ev)
	if err != nil {
		fmt.Println("Error adding fd to epoll:", err)
		return
	}

	var data []byte = make([]byte, 24*64) // 64 input events (sizeof(input_event) == 24)

	for {
		epollEvents := make([]syscall.EpollEvent, 1)
		n, err := syscall.EpollWait(epfd, epollEvents, 500)

		if n > 0 {
			n, err = syscall.Read(fd, data)
			if err != nil {
				fmt.Println("Error reading input_event:", err)
				continue
			}

			p.Send(InputEvents{Data: data, N: n})
		}
	}
}
