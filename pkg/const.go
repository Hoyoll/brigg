package brigg

const DEFID = 0

type State int8

const (
	DEFAULT State = iota
	HOVER
)

type Align int8

const (
	START Align = iota
	END
)

const (
	WRAP_CONTENT float32 = iota
	WINDOW
)

type Gravity int8

const (
	VERTICAL Gravity = iota
	HORIZONTAL
)

type Genus int8

const (
	BOX Genus = iota
	TEXT
	IMAGE
)

type Event int8

const (
	DOWN Event = iota
	UP
	PRESSED
	RELEASED
)
