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

	_, err = s.Write(AlphaProtocol.SetDateTime(time.Now()))
	if err != nil {
		log.Fatal(err)
	}

	//_, err = s.Write(AlphaProtocol.SetStaticText(append([]byte("Hey, its: "), AlphaProtocol.PrepareDateTime()...), AlphaProtocol.COLOR_COLOR_MIX))
	//_, err = s.Write(AlphaProtocol.SetSpecialText([]byte("Sex?"), AlphaProtocol.POSITION_MIDDLE, AlphaProtocol.DEFAULT_MODE_FLASH, AlphaProtocol.SPECIAL_MODE_SNOW))
	_, err = s.Write(AlphaProtocol.SetStaticText([]byte("Hey Jonas!"), AlphaProtocol.COLOR_RED))
	//_, err = s.Write(AlphaProtocol.SetSpecialSign(AlphaProtocol.POSITION_MIDDLE, AlphaProtocol.SPECIAL_SIGN_FIREWORKS))
	if err != nil {
		log.Fatal(err)
	}
}
