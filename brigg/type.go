package brigg

import "image/color"

type Root interface {
	IsPure() bool
	SetPureness(bool)
	GetStyle() *Style
	Listen() *Listen
	GetIO() *Listen
	State() *StateMap
	ChangeState(State)
}

type Composer interface {
	Render()
	Compose(*Tree)
	Position(*Tree)
	GetDimension() (float32, float32)
	SetPosition(float32, float32)
	GetPosition() (float32, float32)
	CheckIO(Root) bool
}

type Component interface {
	Box | Text
}

type Tree struct {
	Self   int   // indice for Leaf
	Parent int   // indice for Leaf
	Branch []int // indices for Tree
}

type Leaf struct {
	Bone   int // indice for Root
	Shadow int // indice for Composer
	Genus  Genus
}

type Listen struct {
	State map[State]func(*Style)
	Key   map[int]func(*Style)
}

func (l *Listen) On(s State, f func(*Style)) {
	l.State[s] = f
}

func (l *Listen) OnKey(key int, f func(*Style)) {
	l.Key[key] = f
}

func (l *Listen) Trigger(s State, style *Style) {
	fun, ok := l.State[s]
	if ok {
		fun(style)
	}
}

type StateMap struct {
	item map[State]int
}

func (s *StateMap) Add(state State, sty *Style) {
	s.item[state] = Styles.Add(*sty)
}

func (s *StateMap) Get(state State) *Style {
	style, ok := s.item[state]
	if !ok {
		return nil
	}
	st, _ := Styles.Get(style)
	return st
}

type Style struct {
	Align     Align
	Gravity   Gravity
	PaddingY  float32
	PaddingX  float32
	MaxHeight float32
	MaxWidth  float32
	Height    float32
	Width     float32
	Spacing   float32
	FontSize  float32
	Radius    float32
	Gap       float32
	Color     *color.RGBA
	Text      *string
	FontPath  *string
}
