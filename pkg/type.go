package brigg

type Composer interface {
	Render(style int)
	CalcPos(treeId int)
	CalcDim(treeId int)
	SetPos(X, Y float32)
	GetDim() (float32, float32)
	GetPos() (float32, float32)
	CheckIO(element int, childs []int) (bool, bool)
}

type Tree struct {
	Pure     bool
	Genus    Genus
	Bones    int
	Renderer Composer
	Branch   []int
}

func (t *Tree) Bone() *Element {
	return &Bones.Items[t.Bones]
}

func (t *Tree) Child(child ...int) {
	t.Branch = append(t.Branch, child...)
}

type StateMap struct {
	Style  map[State]int
	Listen map[State]func(*Style) bool
	Key    map[State]map[Event][]any
}

func (s *StateMap) Get(state State) *Style {
	style, ok := s.Style[state]
	if !ok {
		return nil
	}
	st, _ := Styles.Get(style)
	return st
}

func (s *StateMap) Add(st State, sty Style) func(func(*Style) bool) {
	s.Style[st] = Styles.Add(sty)
	return func(f func(*Style) bool) {
		s.Listen[st] = f
	}
}

func (s *StateMap) OnKey(st State, e Event, key ...any) {
	m := map[Event][]any{}
	m[e] = key
	s.Key[st] = m
}

func (s *StateMap) Trigger(st State) (bool, bool) {
	styleid, exist := s.Style[st]
	fun, funok := s.Listen[st]
	if funok {
		funok = fun(&Styles.Items[styleid])
	}
	return funok, exist
}
