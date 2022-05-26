package AlphaProtocol

type SpecialFunction byte

const (
	SPECIAL_FUNCTION_CLEAR_MEMORY SpecialFunction = 0x24
)

func PrepareClearMemory() []byte {
	return []byte{0x45, byte(SPECIAL_FUNCTION_CLEAR_MEMORY)}
}
