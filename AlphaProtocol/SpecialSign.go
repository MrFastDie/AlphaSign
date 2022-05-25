package AlphaProtocol

type SpecialSign byte

const (
	SPEZIAL_SIGN_THANK_YOU            SpecialSign = 0x53
	SPEZIAL_SIGN_NO_SMOKE             SpecialSign = 0x55
	SPEZIAL_SIGN_DONT_DRINK_AND_DRIVE SpecialSign = 0x56
	SPEZIAL_SIGN_RUNNING_ANIMAIL      SpecialSign = 0x57
	SPEZIAL_SIGN_FIREWORKS            SpecialSign = 0x58
	SPEZIAL_SIGN_TURBO_CAR            SpecialSign = 0x59
	SPEZIAL_SIGN_CHERRY_BOMB          SpecialSign = 0x5A
)

func CreateSpecialSign(position Position, specialSign SpecialSign) []byte {
	return []byte{0x41, 0x41, 0x1B, byte(position), 0x6E, byte(specialSign)}
}
