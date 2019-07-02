package daikin

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/tarm/serial"
)

type Serial struct {
	port    *serial.Port
	manager *Manager
}

func NewSerial(m *Manager) (*Serial, error) {
	config := &serial.Config{
		Name: "COM45",
		Baud: 115200,
	}

	s, err := serial.OpenPort(config)

	return &Serial{
		manager: m,
		port:    s,
	}, err
}

func (s *Serial) Run() chan error {
	c := make(chan error)

	go func() {
		for {
			c <- <-s.readChannel()
		}
	}()

	r := bufio.NewReader(s.port)

	for {
		if b, err := r.ReadBytes('\n'); err == nil {
			json.Unmarshal(b, s.manager.State)
		}
	}

	return c
}

func (s *Serial) readChannel() chan error {
	c := make(chan error)

	for i := range s.manager.SerialChan {
		b, err := json.Marshal(i)

		if err != nil {
			c <- err
		}

		fmt.Printf("Received through channel: %v\n", string(b))
		s.port.Write(b)
	}

	return c
}
