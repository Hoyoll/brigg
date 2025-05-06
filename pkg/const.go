package brigg

// DEFID is the default index ID used for initializing elements.
// It usually means "no actual data assigned yet".
const DEFID = 0

// State defines UI element interaction states, like default or hover.
type State int8

const (
	// DEFAULT is the normal, non-interacted state.
	DEFAULT State = iota

	// HOVER is the state when a user hovers over an element.
	HOVER

	SCROLL_UP
	SCROLL_DOWN
)

// Align defines how children are aligned within a container.
type Align int8

const (
	// START aligns children to the start of the axis (left or top).
	START Align = iota

	// END aligns children to the end of the axis (right or bottom).
	END
)

// Special size values for layout sizing.
const (
	// WRAP_CONTENT means size should be based on children's total size.
	WRAP_CONTENT float32 = iota

	// WINDOW means size should take up the full window.
	WINDOW
)

type Overflow int8

const (
	WRAP Overflow = iota
	HIDDEN 
	LEAK
	SCROLL
)

// Gravity defines layout direction inside a container.
type Gravity int8

const (
	// VERTICAL stacks children top-to-bottom.
	VERTICAL Gravity = iota

	// HORIZONTAL arranges children left-to-right.
	HORIZONTAL
)

// Genus represents the type or "kind" of a layout element.
type Genus int8

const (
	// BOX is a rectangular container. Can hold children and apply constraints.
	BOX Genus = iota

	// TEXT is a text-rendering element.
	TEXT

	// IMAGE is an image-rendering element.
	IMAGE
)

// Event represents input events like mouse or keyboard interactions.
type Event int8

const (
	// DOWN is triggered when an input (mouse/key) is being
	// pressed and remains there.
	DOWN Event = iota

	// UP is triggered when the input is not pressed anymore.
	UP

	// PRESSED is active while an input is pressed, one-time.
	PRESSED

	// RELEASED is triggered only on released, one-time.
	RELEASED
)
