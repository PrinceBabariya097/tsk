package color

type colorCode struct {
	code string
}

type fontStyle struct {
	code string
}

var (
	Reset   = colorCode{"\033[0m"}
	Black   = colorCode{"\033[30m"}
	Red     = colorCode{"\033[31m"}
	Green   = colorCode{"\033[32m"}
	Yellow  = colorCode{"\033[33m"}
	Blue    = colorCode{"\033[34m"}
	Magenta = colorCode{"\033[35m"}
	Cyan    = colorCode{"\033[36m"}
	Gray    = colorCode{"\033[37m"}
	White   = colorCode{"\033[97m"}

	// Bright colors
	BrightBlack   = colorCode{"\033[90m"}
	BrightRed     = colorCode{"\033[91m"}
	BrightGreen   = colorCode{"\033[92m"}
	BrightYellow  = colorCode{"\033[93m"}
	BrightBlue    = colorCode{"\033[94m"}
	BrightMagenta = colorCode{"\033[95m"}
	BrightCyan    = colorCode{"\033[96m"}
	BrightWhite   = colorCode{"\033[97m"}
)

// Font styles
var (
	Bold      = fontStyle{"\033[1m"}
	Dim       = fontStyle{"\033[2m"}
	Italic    = fontStyle{"\033[3m"}
	Underline = fontStyle{"\033[4m"}
	Blink     = fontStyle{"\033[5m"}
	Reverse   = fontStyle{"\033[7m"}
	Hidden    = fontStyle{"\033[8m"}
	Strike    = fontStyle{"\033[9m"}
)

func ApplyStyle(input string, color colorCode, styles ...fontStyle) string {
	c := ""
	for _, style := range styles {
		c += style.code
	}
	c += color.code
	return c + input + Reset.code
}
