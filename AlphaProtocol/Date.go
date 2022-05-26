package AlphaProtocol

import (
	"time"
)

type DateFormat byte

func PrepareSetTime(dateTime time.Time) []byte {
	return append([]byte{0x45, 0x20}, []byte(dateTime.Format("1504"))...)
}

func PrepareSetDate(dateTime time.Time) []byte {
	return append([]byte{0x45, 0x3B}, []byte(dateTime.Format("010206"))...)
}

func PrepareDate() []byte {
	// Second byte represents the date format - other formats might be:
	//• 0BH + “0” (30H) = MM/DD/YY
	//• 0BH + “1” (31H) = DD/MM/YY
	//• 0BH + “2” (32H) = MM-DD-YY
	//• 0BH + “3” (33H) = DD-MM-YY
	//• 0BH + “4” (34H) = MM.DD.YY
	//• 0BH + “5” (35H) = DD.MM.YY
	//• 0BH + “6” (36H) = MM DD YY
	//• 0BH + “7” (37H) = DD MM YY
	//• 0BH + “8” (38H) = MMM.DD, YYYY
	//• 0BH + “9” (39H) = Day of week
	return []byte{0x0B, 0x35}
}

func PrepareTime() []byte {
	return []byte{0x13}
}

func PrepareDateTime() []byte {
	return append(PrepareDate(), append([]byte(" "), PrepareTime()...)...)
}
