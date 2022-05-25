package main

import (
	"github.com/MrFastDie/AlphaSign/AlphaProtocol"
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {
	c := &serial.Config{Name: "COM3", Baud: 9600, Size: 7, Parity: 'E', StopBits: 1, ReadTimeout: 1 * time.Second}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write(AlphaProtocol.InitSequence())
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write(AlphaProtocol.PrepareText([]byte("Hello darkness my old friend...")))
	if err != nil {
		log.Fatal(err)
	}

	_, err = s.Write(AlphaProtocol.DeinitSequence())
	if err != nil {
		log.Fatal(err)
	}
}
