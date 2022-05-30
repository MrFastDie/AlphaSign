package AlphaProtocol

type SpecialFunction byte

const (
	SPECIAL_FUNCTION_CLEAR_MEMORY SpecialFunction = 0x24
)

func prepareClearMemory() []byte {
	return []byte{0x45, byte(SPECIAL_FUNCTION_CLEAR_MEMORY)}
}

func SetClearMemory() []byte {
	sequences := [][]byte{
		InitSequence(),
		prepareClearMemory(),
		DeinitSequence(),
	}

	var returnSequence []byte
	for _, sequence := range sequences {
		returnSequence = append(returnSequence, sequence...)
	}

	return returnSequence
}
