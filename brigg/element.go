package brigg

type Element struct {
	states int // State index
	state  State
}

func (e *Element) State() *StateMap {
	e.states = States.Add(StateMap{
		map[State]int{},
		map[State]func(*Style) bool{},
		map[State][]any{},
	})
	s, _ := States.Get(e.states)
	return s
}

func (e *Element) GetIO() map[State][]any {
	if e.states == DEFID {
		return nil
	}
	st := &States.Items[e.states]
	return st.Key
}

func (b *Element) GetStyle() int {
	state := States.Items[b.states]
	return state.Style[b.state]
}

// func (e *Element) GetStyle() *Style {
// 	state, _ := States.Get(e.states)
// 	style := state.Get(e.state)
// 	return style
// }

func (e *Element) ChangeState(s State) {
	if e.states == DEFID {
		return
	}
	state := States.Items[e.states]
	succ := state.Trigger(s)
	if succ {
		e.state = s
	}
}
