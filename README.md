# AlphaSign
This is an implementation of the Alpha LED Sign serial api

It spins up a local webserver on port `:3000` which allows to control the different modes via web interface. As of today it uses `COM3` with a baud rate of `9600`

Following colors are supported in text mode:
 - {RED}
 - {GREEN}
 - {YELLOW}
 - {RAINBOW1}
 - {RAINBOW2}
 - {COLORMIX}
 - {DATE}
 - {TIME}

The supported modes are available through the UI
