package brigg

// type Root interface {
// 	GetStyle() *Style
// 	GetStyleId() int
// 	GetIO() map[State][]any
// 	State() *StateMap
// 	ChangeState(State)
// }

type Composer interface {
	Render()
	// Compose(*Tree)

	CalcPos(int, int, []int)
	CalcDim(int, int, []int)
	GetDim() (float32, float32)
	GetPos() (float32, float32)
	CheckIO(*Element) bool
}

type Tree struct {
	Pure   bool
	Genus  Genus
	Self   int   // indice for Leaf
	Parent int   // indice for Leaf
	Branch []int // indices for Tree
}

func (t *Tree) Bone() *Element {
	l := Leaves.Items[t.Self]
	return &Bones.Items[l.Bone]
}

func (t *Tree) Child(child ...*Tree) {
	for _, v := range child {
		t.Branch = append(t.Branch, Trees.Add(*v))
	}
}

type Leaf struct {
	Bone int // indice for Root
	// Shadow   int           // indice for Composers
	Renderer map[Genus]int // map of Genus and Composer indc
}

type StateMap struct {
	Style  map[State]int
	Listen map[State]func(*Style) bool
	Key    map[State][]any
}

func (s *StateMap) Get(state State) *Style {
	style, ok := s.Style[state]
	if !ok {
		return nil
	}
	st, _ := Styles.Get(style)
	return st
}

func (s *StateMap) Add(st State, sty *Style) func(
	func(*Style) bool) {

	s.Style[st] = Styles.Add(*sty)
	return func(f func(*Style) bool) {
		s.Listen[st] = f
	}
}

func (s *StateMap) OnKey(st State, key ...any) {
	s.Key[st] = key
}

func (s *StateMap) Trigger(st State) bool {
	styleid, ok := s.Style[st]
	if !ok {
		return false
	}
	fun, funok := s.Listen[st]
	if !funok {
		return true
	}
	return fun(&Styles.Items[styleid])
}
