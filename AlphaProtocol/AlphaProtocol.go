package AlphaProtocol

func InitSequence() []byte {
	return []byte{
		0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x5A, 0x30, 0x30, 0x02,
	}
}

func DeinitSequence() []byte {
	return []byte{
		0x04,
	}
}
