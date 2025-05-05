package brigg

import (
	"image/color"
)

func NewStyle() Style {
	return Style{
		Content: make(map[Genus]int),
	}
}

type Style struct {
	Constraint int
	Content    map[Genus]int
}

func (s *Style) SetText(t Text) {
	s.Content[TEXT] = Texts.Add(t)
}

func (s *Style) GetText() (*Text, bool) {
	k, ok := s.Content[TEXT]
	return &Texts.Items[k], ok
}

func (s *Style) SetImage(image Image) {

	s.Content[IMAGE] = Images.Add(image)
}

func (s *Style) GetImage() (*Image, bool) {
	k, ok := s.Content[IMAGE]
	return &Images.Items[k], ok
}

func (s *Style) SetBox(box Box) {
	s.Content[BOX] = Boxes.Add(box)
}

func (s *Style) GetBox() (*Box, bool) {
	k, ok := s.Content[BOX]
	return &Boxes.Items[k], ok
}

func (s *Style) SetConstraint(c Constraint) {
	s.Constraint = Constraints.Add(c)
}

func (s *Style) GetConstraint() *Constraint {
	return &Constraints.Items[s.Constraint]
}

type Box struct {
	Radius, Rotate float32
	Height, Width  float32
	Color          color.RGBA
}

type Text struct {
	Spacing, Sizing float32
	Text, Font      string
	Tint            color.RGBA
}

type Image struct {
	Rotate, Scale float32
	Path          string
	Tint          color.RGBA
}

type Constraint struct {
	PaddingLeft, PaddingRight float32
	PaddingTop, PaddingBottom float32
	Gap                       float32
	Gravity                   Gravity
	Align                     Align
	Static                    bool
	Ghost                     bool
}
