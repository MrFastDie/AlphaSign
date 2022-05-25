package AlphaProtocol

func PrepareSpecialText(message []byte, position Position, defaultMode DefaultMode, specialMode SpecialMode) []byte {
	if defaultMode == DEFAULT_MODE_SPECIAL {
		return append([]byte{0x41, 0x41, 0x1B, byte(position), byte(defaultMode), byte(specialMode)}, message...)
	}

	return append([]byte{0x41, 0x41, 0x1B, byte(position), byte(defaultMode)}, message...)
}

func PrepareText(message []byte) []byte {
	return append([]byte{0x41, 0x41}, message...)
}

func PrepareStaticText(message []byte, color Color) []byte {
	return PrepareSpecialText(append(AddColor(color), message...), POSITION_MIDDLE, DEFAULT_MODE_HOLD, SPECIAL_MODE_INTERLOCK)
}
