package Serial

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

var config = &serial.Config{Name: "COM3", Baud: 9600, Size: 7, Parity: 'E', StopBits: 1, ReadTimeout: 1 * time.Second}
var Port *serial.Port

func init() {
	if Port != nil {
		return
	}

	var err error
	Port, err = serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}
}
