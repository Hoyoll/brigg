package brigg

// Element is the bone of the UI although not it's marrow
type Element struct {

	// Current State
	CState State

	// StateMap index. This is the marrow
	States int
}

// Add a new StateMap into Element
func (e *Element) State() *StateMap {
	e.States = States.Add(StateMap{
		map[State]int{},
		make(map[State]func(*Style) bool),
		make(map[State]map[Event][]any),
	})
	return &States.Items[e.States]
}

// Exporting an IO listener for the renderer. Not important
// unless you want to make your own renderer.
func (e *Element) GetIO() map[State]map[Event][]any {
	st := States.Items[e.States]
	return st.Key
}

// Return current style index. Not important
// unless you want to make your own renderer.
func (b *Element) GetStyle() int {
	state := States.Items[b.States]
	return state.Style[b.CState]
}

// Change the state of the element, and based on the logic
// gate can change the purity of the element.
// Although if you don't implement the purity filter in
// your renderer it just change the state
func (e *Element) ChangeState(s State) bool {
	state := States.Items[e.States]
	fun, stylechange := state.Trigger(s)
	if stylechange {
		e.CState = s
	}
	return !(fun || stylechange)
}
