package AlphaProtocol

type DefaultMode byte

const (
	DEFAULT_MODE_ROTATE            DefaultMode = 0x61
	DEFAULT_MODE_HOLD              DefaultMode = 0x62
	DEFAULT_MODE_FLASH             DefaultMode = 0x63
	DEFAULT_MODE_ROLL_UP           DefaultMode = 0x65
	DEFAULT_MODE_ROLL_DOWN         DefaultMode = 0x66
	DEFAULT_MODE_ROLL_LEFT         DefaultMode = 0x67
	DEFAULT_MODE_ROLL_RIGHT        DefaultMode = 0x68
	DEFAULT_MODE_WIPE_UP           DefaultMode = 0x69
	DEFAULT_MODE_WIPE_DOWN         DefaultMode = 0x6A
	DEFAULT_MODE_WIPE_LEFT         DefaultMode = 0x6B
	DEFAULT_MODE_WIPE_RIGHT        DefaultMode = 0x6C
	DEFAULT_MODE_SCROLL            DefaultMode = 0x6D
	DEFAULT_MODE_AUTOMODE          DefaultMode = 0x6F
	DEFAULT_MODE_ROLL_IN           DefaultMode = 0x70
	DEFAULT_MODE_ROLL_OUT          DefaultMode = 0x71
	DEFAULT_MODE_WIPE_IN           DefaultMode = 0x72
	DEFAULT_MODE_WIPE_OUT          DefaultMode = 0x73
	DEFAULT_MODE_COMPRESSED_ROTATE DefaultMode = 0x74
	DEFAULT_MODE_EXPLODE           DefaultMode = 0x75
	DEFAULT_MODE_CLOCK             DefaultMode = 0x76
	DEFAULT_MODE_SPECIAL           DefaultMode = 0x6E
)
