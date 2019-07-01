package daikin

import "github.com/tarm/serial"

var s *serial.Port

func InitSerial() error {
	var err error

	config := &serial.Config{}
	s, err = serial.OpenPort(config)

	return err
}

func Write(json string) error {
	_, err := s.Write([]byte(json))

	return err
}

func Read() ([]byte, error) {
	buf := make([]byte, 128)
	_, err := s.Read(buf)

	return buf, err
}