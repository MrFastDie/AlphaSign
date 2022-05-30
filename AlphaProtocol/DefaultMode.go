package AlphaProtocol

import (
	"errors"
	"strings"
)

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

func GetDefaultModeByString(defaultMode string) (DefaultMode, error) {
	switch strings.ToUpper(defaultMode) {
	case "ROTATE":
		return DEFAULT_MODE_ROTATE, nil
	case "HOLD":
		return DEFAULT_MODE_HOLD, nil
	case "FLASH":
		return DEFAULT_MODE_FLASH, nil
	case "ROLL_UP":
		return DEFAULT_MODE_ROLL_UP, nil
	case "ROLL_DOWN":
		return DEFAULT_MODE_ROLL_DOWN, nil
	case "ROLL_LEFT":
		return DEFAULT_MODE_ROLL_LEFT, nil
	case "ROLL_RIGHT":
		return DEFAULT_MODE_ROLL_RIGHT, nil
	case "WIPE_UP":
		return DEFAULT_MODE_WIPE_UP, nil
	case "WIPE_DOWN":
		return DEFAULT_MODE_WIPE_DOWN, nil
	case "WIPE_LEFT":
		return DEFAULT_MODE_WIPE_LEFT, nil
	case "WIPE_RIGHT":
		return DEFAULT_MODE_WIPE_RIGHT, nil
	case "SCROLL":
		return DEFAULT_MODE_SCROLL, nil
	case "AUTOMODE":
		return DEFAULT_MODE_AUTOMODE, nil
	case "ROLL_IN":
		return DEFAULT_MODE_ROLL_IN, nil
	case "ROLL_OUT":
		return DEFAULT_MODE_ROLL_OUT, nil
	case "WIPE_IN":
		return DEFAULT_MODE_WIPE_IN, nil
	case "WIPE_OUT":
		return DEFAULT_MODE_WIPE_OUT, nil
	case "COMPRESSED_ROTATE":
		return DEFAULT_MODE_COMPRESSED_ROTATE, nil
	case "EXPLODE":
		return DEFAULT_MODE_EXPLODE, nil
	case "CLOCK":
		return DEFAULT_MODE_CLOCK, nil
	case "SPECIAL":
		return DEFAULT_MODE_SPECIAL, nil
	}

	return 0, errors.New("unknown default-mode")
}
