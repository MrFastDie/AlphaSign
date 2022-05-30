package AlphaProtocol

import (
	"errors"
	"strings"
)

type Position byte

const (
	POSITION_MIDDLE Position = 0x20
	POSITION_CENTER Position = 0x22
	POSITION_BOTTOM Position = 0x26
	POSITION_FILL   Position = 0x30
	POSITION_LEFT   Position = 0x31
	POSITION_RIGHT  Position = 0x32
)

func GetPositionByString(position string) (Position, error) {
	switch strings.ToUpper(position) {
	case "MIDDLE":
		return POSITION_MIDDLE, nil
	case "CENTER":
		return POSITION_CENTER, nil
	case "BOTTOM":
		return POSITION_BOTTOM, nil
	case "FILL":
		return POSITION_FILL, nil
	case "LEFT":
		return POSITION_LEFT, nil
	case "RIGHT":
		return POSITION_RIGHT, nil
	}

	return 0, errors.New("unknown position")
}
