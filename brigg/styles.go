package brigg

import "image/color"

type Style struct {
	Box        int
	Text       int
	Image      int
	Constraint int
}

func (s *Style) SetText(t *Text) {
	s.Text = Texts.Add(*t)
}

func (s *Style) GetText() *Text {
	return &Texts.Items[s.Text]
}

func (s *Style) SetImage(image *Image) {
	s.Image = Images.Add(*image)
}

func (s *Style) GetImage() *Image {
	return &Images.Items[s.Image]
}

func (s *Style) SetBox(box *Box) {
	s.Box = Boxes.Add(*box)
}

func (s *Style) GetBox() *Box {
	return &Boxes.Items[s.Box]
}

type Box struct {
	Height float32
	Width  float32
	Color  color.RGBA
	Radius float32
	Rotate float32
}

type Text struct {
	Spacing int
	Sizing  int
	Text    int
	Font    int
	Rotate  float32
	Tint    color.RGBA
}

type Image struct {
	Path   string
	Tint   color.RGBA
	Rotate float32
	Scale  float32
}

type Constraint struct {
	Align    Align
	Gravity  Gravity
	Gap      float32
	PaddingY float32
	PaddingX float32
}
