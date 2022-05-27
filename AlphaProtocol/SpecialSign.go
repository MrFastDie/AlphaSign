package AlphaProtocol

type SpecialSign byte

const (
	SPECIAL_SIGN_THANK_YOU            SpecialSign = 0x53
	SPECIAL_SIGN_NO_SMOKE             SpecialSign = 0x55
	SPECIAL_SIGN_DONT_DRINK_AND_DRIVE SpecialSign = 0x56
	SPECIAL_SIGN_RUNNING_ANIMAIL      SpecialSign = 0x57
	SPECIAL_SIGN_FIREWORKS            SpecialSign = 0x58
	SPECIAL_SIGN_TURBO_CAR            SpecialSign = 0x59
	SPECIAL_SIGN_CHERRY_BOMB          SpecialSign = 0x5A
)

func createSpecialSign(position Position, specialSign SpecialSign) []byte {
	return []byte{0x41, 0x41, 0x1B, byte(position), 0x6E, byte(specialSign)}
}

func SetSpecialSign(position Position, specialSign SpecialSign) []byte {
	sequences := [][]byte{
		InitSequence(),
		createSpecialSign(position, specialSign),
		DeinitSequence(),
	}

	var returnSequence []byte
	for _, sequence := range sequences {
		returnSequence = append(returnSequence, sequence...)
	}

	return returnSequence
}
