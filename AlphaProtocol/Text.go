package AlphaProtocol

func prepareSpecialText(message []byte, position Position, defaultMode DefaultMode, specialMode SpecialMode) []byte {
	if defaultMode == DEFAULT_MODE_SPECIAL {
		return append([]byte{0x41, 0x41, 0x1B, byte(position), byte(defaultMode), byte(specialMode)}, message...)
	}

	return append([]byte{0x41, 0x41, 0x1B, byte(position), byte(defaultMode)}, message...)
}

func SetSpecialText(message []byte, position Position, defaultMode DefaultMode, specialMode SpecialMode) []byte {
	sequences := [][]byte{
		InitSequence(),
		prepareSpecialText(message, position, defaultMode, specialMode),
		DeinitSequence(),
	}

	var returnSequence []byte
	for _, sequence := range sequences {
		returnSequence = append(returnSequence, sequence...)
	}

	return returnSequence
}

func prepareText(message []byte) []byte {
	return append([]byte{0x41, 0x41}, message...)
}

func SetText(message []byte) []byte {
	sequences := [][]byte{
		InitSequence(),
		prepareText(message),
		DeinitSequence(),
	}

	var returnSequence []byte
	for _, sequence := range sequences {
		returnSequence = append(returnSequence, sequence...)
	}

	return returnSequence
}

func prepareStaticText(message []byte, color Color) []byte {
	return prepareSpecialText(append(AddColor(color), message...), POSITION_MIDDLE, DEFAULT_MODE_HOLD, SPECIAL_MODE_INTERLOCK)
}

func SetStaticText(message []byte, color Color) []byte {
	sequences := [][]byte{
		InitSequence(),
		prepareStaticText(message, color),
		DeinitSequence(),
	}

	var returnSequence []byte
	for _, sequence := range sequences {
		returnSequence = append(returnSequence, sequence...)
	}

	return returnSequence
}
