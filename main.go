package main

import (
	"fmt"
	"io"

	"github.com/jacobsa/go-serial/serial"
)

var magicBytes = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
var brightnessBytes = []byte{0x0A, 0x0A, 0x0A, 0x0A, 0x0A, 0x0A}

var serialPort io.ReadWriteCloser

func main() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyACM0",
		BaudRate:        2000000,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		fmt.Println("opening serial failed")
		return
	}

	defer port.Close()

	writeBytes(magicBytes)

	writeBytes(brightnessBytes)

}

func writeBytes(bytes []byte) int {
	n, err := serialPort.Write(bytes)
	if err != nil {
		fmt.Println("Error writing to serial...")
		return -1
	}

	return n
}
