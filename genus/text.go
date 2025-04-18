package genus

import "github.com/Hoyoll/brick/brick"

type Text struct {
	Style    *brick.Style
	State    brick.StateFlag
	States   *brick.States
	Listener *brick.Listener
	Text     string
	Pure     bool
}

func (b *Text) ChangeState(s brick.StateFlag) {
	b.State = s
	b.Style = b.States.GetState(s)
}

func (t *Text) GetIO() *brick.Listener {
	return t.Listener
}

func (t *Text) SetStyle(s *brick.Style) {
	t.Style = s
}

func (t *Text) GetStyle() *brick.Style {
	return t.Style
}

func (t *Text) IsPure() bool {
	return t.Pure
}

func (t *Text) SetPureness(p bool) {
	t.Pure = p
}
