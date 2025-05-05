package brigg

type Element struct {
	CState State
	States int // State index
}

func (e *Element) State() *StateMap {
	e.States = States.Add(StateMap{
		map[State]int{},
		make(map[State]func(*Style) bool),
		make(map[State]map[Event][]any),
	})
	return &States.Items[e.States]
}

func (e *Element) GetIO() map[State]map[Event][]any {
	st := States.Items[e.States]
	return st.Key
}

func (b *Element) GetStyle() int {
	state := States.Items[b.States]
	return state.Style[b.CState]
}

func (e *Element) ChangeState(s State) bool {
	state := States.Items[e.States]
	fun, stylechange := state.Trigger(s)
	if stylechange {
		e.CState = s
	}
	return !(fun || stylechange)
}
