package brick

import (
	"image/color"
)

type Style struct {
	PaddingY  float32
	PaddingX  float32
	MaxHeight Size
	MaxWidth  Size
	Height    Size
	Width     Size
	Gravity   Gravity
	Spacing   float32
	FontSize  float32
	Gap       float32
	Align     Align
	Text      string
	Color     *color.RGBA
	FontPath  string
	Radius    float32
}

type Tree struct {
	Self   *Leaf
	Parent *Leaf
	Branch []*Tree
}

func (t *Tree) Child(childs ...*Tree) {
	for _, v := range childs {
		v.Parent = t.Self
	}
	t.Branch = append(t.Branch, childs...)
}

type Leaf struct {
	Genus  Genus
	Bone   Root
	Shadow Composer
}

type Listener struct {
	MouseClick map[int]func(*Style)
	Normal     []func(*Style)
	Hover      func(*Style)
	Key        map[int]func(*Style)
}

func (l *Listener) OnClick(key int, process func(*Style)) {
	l.MouseClick[key] = process
}

func (l *Listener) OnHover(process func(*Style)) {
	l.Hover = process
}

func (l *Listener) OnKey(key int, process func(*Style)) {
	l.Key[key] = process
}

type State struct {
	Style *Style
}

type States struct {
	Hover  *Style
	Click  *Style
	Key    *Style
	Normal *Style
}

func (s *States) Fill(st *Style) {
	clone := func(s Style) *Style {
		return &s
	}
	s.Normal = clone(*st)
	s.Click = clone(*st)
	s.Key = clone(*st)
	s.Hover = clone(*st)
}

func (s *States) GetState(st StateFlag) *Style {
	var res *Style
	switch st {
	case CLICK:
		res = s.Click
	case HOVER:
		res = s.Hover
	case KEYDOWN:
		res = s.Key
	default:
		res = s.Normal
	}
	return res
}
