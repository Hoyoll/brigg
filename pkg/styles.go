package brigg

import (
	"image/color"
)

// NewStyle creates a new Style with initialized content map.
// Styles hold layout and visual data for different Genus types (TEXT, BOX, IMAGE, etc.).
func NewStyle() Style {
	return Style{
		Content: make(map[Genus]int),
	}
}

// Style represents a visual and layout configuration for an element.
// It maps each Genus (like BOX or TEXT) to its raw style data (like a Box or Text struct).
type Style struct {
	// Constraint points to layout constraints (padding, alignment, etc.).
	Constraint int

	// Content maps Genus types (TEXT, BOX, IMAGE, etc.) to their respective style IDs.
	Content map[Genus]int
}

// SetText assigns a Text style to this Style.
func (s *Style) SetText(t Text) {
	s.Content[TEXT] = Texts.Add(t)
}

// GetText retrieves the Text style from this Style, if present.
func (s *Style) GetText() (*Text, bool) {
	k, ok := s.Content[TEXT]
	return &Texts.Items[k], ok
}

// SetImage assigns an Image style to this Style.
func (s *Style) SetImage(image Image) {
	s.Content[IMAGE] = Images.Add(image)
}

// GetImage retrieves the Image style from this Style, if present.
func (s *Style) GetImage() (*Image, bool) {
	k, ok := s.Content[IMAGE]
	return &Images.Items[k], ok
}

// SetBox assigns a Box style to this Style.
func (s *Style) SetBox(box Box) {
	s.Content[BOX] = Boxes.Add(box)
}

// GetBox retrieves the Box style from this Style, if present.
func (s *Style) GetBox() (*Box, bool) {
	k, ok := s.Content[BOX]
	return &Boxes.Items[k], ok
}

// SetConstraint sets the layout constraints (padding, alignment, etc.) for this Style.
func (s *Style) SetConstraint(c Constraint) {
	s.Constraint = Constraints.Add(c)
}

// GetConstraint retrieves the Constraint struct for this Style.
func (s *Style) GetConstraint() *Constraint {
	return &Constraints.Items[s.Constraint]
}

// Box represents a rectangular visual component.
// Used for layout and basic visual styling.
type Box struct {

	// This is a way to determine a child ofset inside scroll
	// mode
	OffsetX, OffsetY float32

	// Border radius
	Radius float32

	// Dimensions
	Height, Width       float32
	MaxHeight, MaxWidth float32

	// Overflow
	Overflow Overflow

	// Fill color
	Color color.RGBA
}

// Text represents styled text content.
type Text struct {

	// Letter spacing and font size
	Spacing, Sizing float32

	// Text content and fontpath
	Text, Font string

	// Rotation angle
	Rotate float32

	// Text color
	Tint color.RGBA
}

// Image represents an image element with visual modifiers.
type Image struct {

	// Rotation and scale factor
	Rotate, Scale float32

	// File path to the image
	Path string

	// Tint color overlay
	Tint color.RGBA
}

// Constraint defines layout behavior like padding, alignment, and spacing.
type Constraint struct {

	// Horizontal padding
	PaddingLeft, PaddingRight float32

	// Vertical padding
	PaddingTop, PaddingBottom float32

	// Space between children
	Gap float32

	// Layout direction (horizontal/vertical)
	Gravity Gravity

	// Child alignment (start, end)
	Align Align

	// Make the element static ie, it does not affect the
	// layout computation for other element
	Static bool

	// Element will listen to IO even if it get blocked by
	// other element
	Ghost bool
}
