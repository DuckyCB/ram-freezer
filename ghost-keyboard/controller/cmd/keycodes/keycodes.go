package keycodes

const (
	KEYCODE_NONE = 0x00
	// modifiers
	KEYCODE_CTRL        = 0x01
	KEYCODE_SHIFT       = 0x02
	KEYCODE_ALT         = 0x04
	KEYCODE_META        = 0x08
	KEYCODE_LEFT_CTRL   = 0x01
	KEYCODE_RIGHT_CTRL  = 0x10
	KEYCODE_LEFT_SHIFT  = 0x02
	KEYCODE_RIGHT_SHIFT = 0x20
	KEYCODE_LEFT_ALT    = 0x04
	KEYCODE_RIGHT_ALT   = 0x40
	KEYCODE_LEFT_META   = 0x08
	KEYCODE_RIGHT_META  = 0x80
	// keys
	KEYCODE_A                = 0x04
	KEYCODE_B                = 0x05
	KEYCODE_C                = 0x06
	KEYCODE_D                = 0x07
	KEYCODE_E                = 0x08
	KEYCODE_F                = 0x09
	KEYCODE_G                = 0x0a
	KEYCODE_H                = 0x0b
	KEYCODE_I                = 0x0c
	KEYCODE_J                = 0x0d
	KEYCODE_K                = 0x0e
	KEYCODE_L                = 0x0f
	KEYCODE_M                = 0x10
	KEYCODE_N                = 0x11
	KEYCODE_O                = 0x12
	KEYCODE_P                = 0x13
	KEYCODE_Q                = 0x14
	KEYCODE_R                = 0x15
	KEYCODE_S                = 0x16
	KEYCODE_T                = 0x17
	KEYCODE_U                = 0x18
	KEYCODE_V                = 0x19
	KEYCODE_W                = 0x1a
	KEYCODE_X                = 0x1b
	KEYCODE_Y                = 0x1c
	KEYCODE_Z                = 0x1d
	KEYCODE_NUMBER_1         = 0x1e
	KEYCODE_NUMBER_2         = 0x1f
	KEYCODE_NUMBER_3         = 0x20
	KEYCODE_NUMBER_4         = 0x21
	KEYCODE_NUMBER_5         = 0x22
	KEYCODE_NUMBER_6         = 0x23
	KEYCODE_NUMBER_7         = 0x24
	KEYCODE_NUMBER_8         = 0x25
	KEYCODE_NUMBER_9         = 0x26
	KEYCODE_NUMBER_0         = 0x27
	KEYCODE_ENTER            = 0x28
	KEYCODE_ESCAPE           = 0x29
	KEYCODE_BACKSPACE        = 0x2a
	KEYCODE_TAB              = 0x2b
	KEYCODE_SPACEBAR         = 0x2c
	KEYCODE_MINUS            = 0x2d
	KEYCODE_EQUAL_SIGN       = 0x2e
	KEYCODE_LEFT_BRACKET     = 0x2f
	KEYCODE_RIGHT_BRACKET    = 0x30
	KEYCODE_BACKSLASH        = 0x31
	KEYCODE_HASH             = 0x32
	KEYCODE_SEMICOLON        = 0x33
	KEYCODE_SINGLE_QUOTE     = 0x34
	KEYCODE_ACCENT_GRAVE     = 0x35
	KEYCODE_COMMA            = 0x36
	KEYCODE_PERIOD           = 0x37
	KEYCODE_FORWARD_SLASH    = 0x38
	KEYCODE_CAPS_LOCK        = 0x39
	KEYCODE_F1               = 0x3a
	KEYCODE_F2               = 0x3b
	KEYCODE_F3               = 0x3c
	KEYCODE_F4               = 0x3d
	KEYCODE_F5               = 0x3e
	KEYCODE_F6               = 0x3f
	KEYCODE_F7               = 0x40
	KEYCODE_F8               = 0x41
	KEYCODE_F9               = 0x42
	KEYCODE_F10              = 0x43
	KEYCODE_F11              = 0x44
	KEYCODE_F12              = 0x45
	KEYCODE_PRINT_SCREEN     = 0x46
	KEYCODE_SCROLL_LOCK      = 0x47
	KEYCODE_PAUSE_BREAK      = 0x48
	KEYCODE_INSERT           = 0x49
	KEYCODE_HOME             = 0x4a
	KEYCODE_PAGE_UP          = 0x4b
	KEYCODE_DELETE           = 0x4c
	KEYCODE_END              = 0x4d
	KEYCODE_PAGE_DOWN        = 0x4e
	KEYCODE_RIGHT_ARROW      = 0x4f
	KEYCODE_LEFT_ARROW       = 0x50
	KEYCODE_DOWN_ARROW       = 0x51
	KEYCODE_UP_ARROW         = 0x52
	KEYCODE_CLEAR            = 0x53
	KEYCODE_NUM_LOCK         = 0x53
	KEYCODE_NUMPAD_DIVIDE    = 0x54
	KEYCODE_NUMPAD_MULTIPLY  = 0x55
	KEYCODE_NUMPAD_MINUS     = 0x56
	KEYCODE_NUMPAD_PLUS      = 0x57
	KEYCODE_NUMPAD_ENTER     = 0x58
	KEYCODE_NUMPAD_1         = 0x59
	KEYCODE_NUMPAD_2         = 0x5a
	KEYCODE_NUMPAD_3         = 0x5b
	KEYCODE_NUMPAD_4         = 0x5c
	KEYCODE_NUMPAD_5         = 0x5d
	KEYCODE_NUMPAD_6         = 0x5e
	KEYCODE_NUMPAD_7         = 0x5f
	KEYCODE_NUMPAD_8         = 0x60
	KEYCODE_NUMPAD_9         = 0x61
	KEYCODE_NUMPAD_0         = 0x62
	KEYCODE_NUMPAD_DOT       = 0x63
	KEYCODE_102ND            = 0x64
	KEYCODE_CONTEXT_MENU     = 0x65
	KEYCODE_F13              = 0x68
	KEYCODE_F14              = 0x69
	KEYCODE_F15              = 0x6a
	KEYCODE_F16              = 0x6b
	KEYCODE_F17              = 0x6c
	KEYCODE_F18              = 0x6d
	KEYCODE_F19              = 0x6e
	KEYCODE_F20              = 0x6f
	KEYCODE_F21              = 0x70
	KEYCODE_F22              = 0x71
	KEYCODE_F23              = 0x72
	KEYCODE_EXECUTE          = 0x74
	KEYCODE_HELP             = 0x75
	KEYCODE_SELECT           = 0x77
	KEYCODE_INTL_RO          = 0x87
	KEYCODE_INTL_YEN         = 0x89
	KEYCODE_HANGEUL          = 0x90
	KEYCODE_HANJA            = 0x91
	KEYCODE_MEDIA_PLAY_PAUSE = 0xe8
	KEYCODE_REFRESH          = 0xfa
)

var Empty = []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}

var Key = map[rune]byte{
	'a':  KEYCODE_A,
	'b':  KEYCODE_B,
	'c':  KEYCODE_C,
	'd':  KEYCODE_D,
	'e':  KEYCODE_E,
	'f':  KEYCODE_F,
	'g':  KEYCODE_G,
	'h':  KEYCODE_H,
	'i':  KEYCODE_I,
	'j':  KEYCODE_J,
	'k':  KEYCODE_K,
	'l':  KEYCODE_L,
	'm':  KEYCODE_M,
	'n':  KEYCODE_N,
	'o':  KEYCODE_O,
	'p':  KEYCODE_P,
	'q':  KEYCODE_Q,
	'r':  KEYCODE_R,
	's':  KEYCODE_S,
	't':  KEYCODE_T,
	'u':  KEYCODE_U,
	'v':  KEYCODE_V,
	'w':  KEYCODE_W,
	'x':  KEYCODE_X,
	'y':  KEYCODE_Y,
	'z':  KEYCODE_Z,
	1:    KEYCODE_NUMBER_1,
	2:    KEYCODE_NUMBER_2,
	3:    KEYCODE_NUMBER_3,
	4:    KEYCODE_NUMBER_4,
	5:    KEYCODE_NUMBER_5,
	6:    KEYCODE_NUMBER_6,
	7:    KEYCODE_NUMBER_7,
	8:    KEYCODE_NUMBER_8,
	9:    KEYCODE_NUMBER_9,
	0:    KEYCODE_NUMBER_0,
	'\n': KEYCODE_ENTER,
	//'\b': KEYCODE_ESCAPE,
	//'	':  KEYCODE_TAB,
	' ':  KEYCODE_SPACEBAR,
	'-':  KEYCODE_MINUS,
	'=':  KEYCODE_EQUAL_SIGN,
	'[':  KEYCODE_LEFT_BRACKET,
	']':  KEYCODE_RIGHT_BRACKET,
	'\\': KEYCODE_BACKSLASH,
	'#':  KEYCODE_HASH,
	';':  KEYCODE_SEMICOLON,
	'\'': KEYCODE_SINGLE_QUOTE,
	'`':  KEYCODE_ACCENT_GRAVE,
	',':  KEYCODE_COMMA,
	'.':  KEYCODE_PERIOD,
	'/':  KEYCODE_FORWARD_SLASH,
}

var KeyShift = map[rune]byte{
	'A': KEYCODE_A,
	'B': KEYCODE_B,
	'C': KEYCODE_C,
	'D': KEYCODE_D,
	'E': KEYCODE_E,
	'F': KEYCODE_F,
	'G': KEYCODE_G,
	'H': KEYCODE_H,
	'I': KEYCODE_I,
	'J': KEYCODE_J,
	'K': KEYCODE_K,
	'L': KEYCODE_L,
	'M': KEYCODE_M,
	'N': KEYCODE_N,
	'O': KEYCODE_O,
	'P': KEYCODE_P,
	'Q': KEYCODE_Q,
	'R': KEYCODE_R,
	'S': KEYCODE_S,
	'T': KEYCODE_T,
	'U': KEYCODE_U,
	'V': KEYCODE_V,
	'W': KEYCODE_W,
	'X': KEYCODE_X,
	'Y': KEYCODE_Y,
	'Z': KEYCODE_Z,
	'!': KEYCODE_NUMBER_1,
	'@': KEYCODE_NUMBER_2,
	'#': KEYCODE_NUMBER_3,
	'$': KEYCODE_NUMBER_4,
	'%': KEYCODE_NUMBER_5,
	'^': KEYCODE_NUMBER_6,
	'&': KEYCODE_NUMBER_7,
	'*': KEYCODE_NUMBER_8,
	'(': KEYCODE_NUMBER_9,
	')': KEYCODE_NUMBER_0,
	'_': KEYCODE_MINUS,
	'+': KEYCODE_EQUAL_SIGN,
	'{': KEYCODE_LEFT_BRACKET,
	'}': KEYCODE_RIGHT_BRACKET,
	'|': KEYCODE_BACKSLASH,
	':': KEYCODE_SEMICOLON,
	'"': KEYCODE_SINGLE_QUOTE,
	'<': KEYCODE_COMMA,
	'>': KEYCODE_PERIOD,
	'?': KEYCODE_FORWARD_SLASH,
}

var SpecialKey = map[string]byte{
	"SHIFT":       KEYCODE_SHIFT,
	"CTRL":        KEYCODE_CTRL,
	"ALT":         KEYCODE_ALT,
	"META":        KEYCODE_META,
	"ESC":         KEYCODE_ESCAPE,
	"BACKSPACE":   KEYCODE_BACKSPACE,
	"TAB":         KEYCODE_TAB,
	"CAPS":        KEYCODE_CAPS_LOCK,
	"F1":          KEYCODE_F1,
	"F2":          KEYCODE_F2,
	"F3":          KEYCODE_F3,
	"F4":          KEYCODE_F4,
	"F5":          KEYCODE_F5,
	"F6":          KEYCODE_F6,
	"F7":          KEYCODE_F7,
	"F8":          KEYCODE_F8,
	"F9":          KEYCODE_F9,
	"F10":         KEYCODE_F10,
	"F11":         KEYCODE_F11,
	"F12":         KEYCODE_F12,
	"PRT_SCR":     KEYCODE_PRINT_SCREEN,
	"INS":         KEYCODE_INSERT,
	"DEL":         KEYCODE_DELETE,
	"RIGHT_ARROW": KEYCODE_RIGHT_ARROW,
	"LEFT_ARROW,": KEYCODE_LEFT_ARROW,
	"DOWN_RIGHT":  KEYCODE_DOWN_ARROW,
	"UP_RIGHT":    KEYCODE_UP_ARROW,
	"LEFT_CTRL":   KEYCODE_LEFT_CTRL,
	"LEFT_SHIFT":  KEYCODE_LEFT_SHIFT,
	"LEFT_ALT":    KEYCODE_LEFT_ALT,
	"LEFT_META":   KEYCODE_LEFT_META,
	"RIGHT_CTRL":  KEYCODE_RIGHT_CTRL,
	"RIGHT_SHIFT": KEYCODE_RIGHT_SHIFT,
	"RIGHT_ALT":   KEYCODE_RIGHT_ALT,
	"RIGHT_META":  KEYCODE_RIGHT_META,
}

var ModifierKey = map[string]byte{
	"SHIFT":       KEYCODE_SHIFT,
	"CTRL":        KEYCODE_CTRL,
	"ALT":         KEYCODE_ALT,
	"META":        KEYCODE_META,
	"LEFT_CTRL":   KEYCODE_LEFT_CTRL,
	"LEFT_SHIFT":  KEYCODE_LEFT_SHIFT,
	"LEFT_ALT":    KEYCODE_LEFT_ALT,
	"LEFT_META":   KEYCODE_LEFT_META,
	"RIGHT_CTRL":  KEYCODE_RIGHT_CTRL,
	"RIGHT_SHIFT": KEYCODE_RIGHT_SHIFT,
	"RIGHT_ALT":   KEYCODE_RIGHT_ALT,
	"RIGHT_META":  KEYCODE_RIGHT_META,
}
