package brigg

// Composer is the interface you implement to create a custom renderer for Brigg.
// It handles layout calculation, drawing, and interaction detection.
type Composer interface {

	// Your mission for each method is:

	// Draws the element using the given style ID
	Render(style int)

	// Calculates position based on layout tree
	CalcPos(treeId int)

	// Calculates width/height based on constraints
	CalcDim(treeId int)

	// Sets the final screen position
	SetPos(X, Y float32)

	// Returns width and height
	GetDim() (float32, float32)

	// Returns screen X and Y
	GetPos() (float32, float32)

	// So you can implement the logic gate from here.
	// Depends on how you want your renderer to be.
	CheckIO(element int, childs []int) (bool, bool)
}

type Tree struct {
	// Pure flags are evaluated every frame.
	// If true, layout recalculation is skipped for performance.
	// Although it depends on how your renderer approach it
	Pure bool

	// Defines how this tree should behave and be rendered (like BOX, TEXT, etc.).
	Genus Genus

	// Index pointing to this tree’s bones — raw, uncomputed layout data.
	Bones int

	// The renderer (implements Composer) that draws this tree based on its Genus.
	Renderer Composer

	// Child trees (layout hierarchy).
	Branch []int
}

// Bone returns the Element (aka "bone") associated with this Tree.
// Bones hold raw, uncomputed layout and visual data.
func (t *Tree) Bone() *Element {
	return &Bones.Items[t.Bones]
}

// Child appends one or more child node IDs to the tree.
// This builds out the layout hierarchy manually.
func (t *Tree) Child(child ...int) {
	t.Branch = append(t.Branch, child...)
}

// StateMap maps UI states (like DEFAULT, HOVER, YOUR_CUSTOM_STATE)
// to styles, event listeners, and key triggers.
// It's used to define how an element should look and behave in different states.
type StateMap struct {

	// Style maps a UI state to a Style ID.
	Style map[State]int

	// Listen maps a UI state to a function that runs when the state is triggered.
	// It receives the current Style and returns bool to determine the if it has
	// side effect. True means the function does not cause recompute, false does
	Listen map[State]func(*Style) bool

	// Key maps a state to specific key events.
	// Used for more advanced input handling.
	Key map[State]map[Event][]any
}

// Get returns the Style associated with a given state.
// Returns nil if no style is found for that state.
func (s *StateMap) Get(state State) *Style {
	style, ok := s.Style[state]
	if !ok {
		return nil
	}
	st, _ := Styles.Get(style)
	return st
}

// Add assigns a Style to a given state, and returns a function
// to optionally attach a listener/callback to that state.
// Example:
//     state.Add(brigg.HOVER, style)(func(s *Style) bool { ... })
// Return true if you want the function to not change the state
func (s *StateMap) Add(st State, sty Style) func(func(*Style) bool) {
	s.Style[st] = Styles.Add(sty)
	return func(f func(*Style) bool) {
		s.Listen[st] = f
	}
}

// OnKey binds a key Event (like a keyboard shortcut or input trigger)
// to a specific state.
// The key event payloads (key...) are passed to the renderer
// — their meaning is up to your implementation.
// Used mostly for keyboard and mouse input cases
func (s *StateMap) OnKey(st State, e Event, key ...any) {
	m := map[Event][]any{}
	m[e] = key
	s.Key[st] = m
}

// Trigger runs the listener function associated with a given state.
// Returns (listenerRan, styleExists) — both bools help determine if
// state logic should continue.
func (s *StateMap) Trigger(st State) (bool, bool) {
	styleid, exist := s.Style[st]
	fun, funok := s.Listen[st]
	if funok {
		funok = fun(&Styles.Items[styleid])
	}
	return funok, exist
}
