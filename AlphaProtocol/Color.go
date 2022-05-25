package AlphaProtocol

type Color byte

const (
	COLOR_RED        Color = 0x31
	COLOR_GREEN      Color = 0x32
	COLOR_AMBER      Color = 0x33
	COLOR_DIM_RED    Color = 0x34
	COLOR_DIM_GREEN  Color = 0x35
	COLOR_BROWN      Color = 0x36
	COLOR_ORANGE     Color = 0x37
	COLOR_YELLOW     Color = 0x38
	COLOR_RAINNBOW_1 Color = 0x39
	COLOR_RAINNBOW_2 Color = 0x41
	COLOR_COLOR_MIX  Color = 0x42
	COLOR_AUTOCOLOR  Color = 0x43
)

func AddColor(color Color) []byte {
	return []byte{0x1C, byte(color)}
}
