package brigg

const DEFID = 0

type State int8

const (
	HOVER State = iota
	DEFAULT
)

type Align int8

const (
	START Align = iota
	END
	CENTER
)

// type Size float32

const (
	WRAP_CONTENT float32 = iota
	MATCH_PARENT
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
)
