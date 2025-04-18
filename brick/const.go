package brick

type StateFlag int

const (
	HOVER StateFlag = iota
	CLICK
	KEYDOWN
	NORMAL
)

type Pure int

const (
	Pure_PRIME Pure = iota
	Pure_RECALC
	Pure_REPAINT
)

type Align int

const (
	Align_START Align = iota
	Align_END
)

type Size float32

const (
	Size_CHILD_SUM Size = iota
	Size_MATCH_PARENT
)

type Gravity int

const (
	Gravity_VERTICAL Gravity = iota
	Gravity_HORIZONTAL
)

type Genus int

const (
	Genus_BOX Genus = iota
	Genus_TEXT
)

type Key int32

const (
	// Keyboard Function Keys
	KeySpace        = 32
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyKbMenu       = 348
	KeyLeftBracket  = 91
	KeyBackSlash    = 92
	KeyRightBracket = 93
	KeyGrave        = 96

	// Keyboard Alpha Numeric Keys
	KeyApostrophe = 39
	KeyComma      = 44
	KeyMinus      = 45
	KeyPeriod     = 46
	KeySlash      = 47
	KeyZero       = 48
	KeyOne        = 49
	KeyTwo        = 50
	KeyThree      = 51
	KeyFour       = 52
	KeyFive       = 53
	KeySix        = 54
	KeySeven      = 55
	KeyEight      = 56
	KeyNine       = 57
	KeySemicolon  = 59
	KeyEqual      = 61
	KeyA          = 65
	KeyB          = 66
	KeyC          = 67
	KeyD          = 68
	KeyE          = 69
	KeyF          = 70
	KeyG          = 71
	KeyH          = 72
	KeyI          = 73
	KeyJ          = 74
	KeyK          = 75
	KeyL          = 76
	KeyM          = 77
	KeyN          = 78
	KeyO          = 79
	KeyP          = 80
	KeyQ          = 81
	KeyR          = 82
	KeyS          = 83
	KeyT          = 84
	KeyU          = 85
	KeyV          = 86
	KeyW          = 87
	KeyX          = 88
	KeyY          = 89
	KeyZ          = 90

	MouseLeftButton   = 0
	MouseRightButton  = 1
	MouseMiddleButton = 2
)
