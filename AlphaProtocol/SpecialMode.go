package AlphaProtocol

import (
	"errors"
	"strings"
)

type SpecialMode byte

const (
	SPECIAL_MODE_TWINKLE      SpecialMode = 0x30
	SPECIAL_MODE_SPARKLE      SpecialMode = 0x31
	SPECIAL_MODE_SNOW         SpecialMode = 0x32
	SPECIAL_MODE_INTERLOCK    SpecialMode = 0x33
	SPECIAL_MODE_SWITCH       SpecialMode = 0x34
	SPECIAL_MODE_SLIDE        SpecialMode = 0x35
	SPECIAL_MODE_SPRAY        SpecialMode = 0x36
	SPECIAL_MODE_STARBUST     SpecialMode = 0x37
	SPECIAL_MODE_WELCOME      SpecialMode = 0x38
	SPECIAL_MODE_SLOT_MACHINE SpecialMode = 0x39
	SPECIAL_MODE_CYCLE_COLOR  SpecialMode = 0x43
)

func GetSpecialModeByString(specialMode string) (SpecialMode, error) {
	switch strings.ToUpper(specialMode) {
	case "TWINKLE":
		return SPECIAL_MODE_TWINKLE, nil
	case "SPARKLE":
		return SPECIAL_MODE_SPARKLE, nil
	case "SNOW":
		return SPECIAL_MODE_SNOW, nil
	case "INTERLOCK":
		return SPECIAL_MODE_INTERLOCK, nil
	case "SWITCH":
		return SPECIAL_MODE_SWITCH, nil
	case "SLIDE":
		return SPECIAL_MODE_SLIDE, nil
	case "SPRAY":
		return SPECIAL_MODE_SPRAY, nil
	case "STARBUST":
		return SPECIAL_MODE_STARBUST, nil
	case "WELCOME":
		return SPECIAL_MODE_WELCOME, nil
	case "SLOT_MACHINE":
		return SPECIAL_MODE_SLOT_MACHINE, nil
	case "CYCLE_COLOR":
		return SPECIAL_MODE_CYCLE_COLOR, nil
	}

	return 0, errors.New("unknown special-mode")
}
