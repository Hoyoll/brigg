package brigg

type Box struct {
	pure   bool
	state  State
	states int // State index
	listen int // Listen index
}

type Text struct {
}

func (b *Box) State() *StateMap {
	b.states = States.Add(StateMap{
		map[State]int{},
	})
	s, _ := States.Get(b.states)
	return s
}

func (b *Box) Listen() *Listen {
	b.listen = Listens.Add(Listen{
		State: map[State]func(*Style){},
		Key:   map[int]func(*Style){},
	})
	l, _ := Listens.Get(b.listen)
	return l
}

func (b *Box) GetIO() *Listen {
	if b.listen == -1 {
		return nil
	}
	l, _ := Listens.Get(b.listen)
	return l
}

func (b *Box) GetStyle() *Style {
	state, _ := States.Get(b.states)
	style := state.Get(b.state)
	return style
}

func (b *Box) ChangeState(s State) {
	if b.listen == -1 || b.states == -1 {
		return
	}
	list, _ := Listens.Get(b.listen)
	state, _ := States.Get(b.states)
	style := state.Get(s)
	if style == nil {
		return
	}
	b.state = s
	list.Trigger(b.state, style)
}

func (b *Box) SetPureness(p bool) {
	b.pure = p
}

func (b *Box) IsPure() bool {
	return b.pure
}
