package genus

import "github.com/Hoyoll/brick/brick"

type Box struct {
	Style    *brick.Style
	State    brick.StateFlag
	States   *brick.States
	Listener *brick.Listener
	Pure     bool
}

func (b *Box) Listen() *brick.Listener {
	listen := &brick.Listener{
		MouseClick: make(map[int]func(*brick.Style)),
		Key:        make(map[int]func(*brick.Style)),
	}
	b.Listener = listen
	return listen
}

func (b *Box) ChangeState(s brick.StateFlag) {
	b.State = s
	b.Style = b.States.GetState(s)
}

func (b *Box) SetStyle(c *brick.Style) {
	b.States.Fill(c)
	b.Style = b.States.Normal
}

func (b *Box) GetIO() *brick.Listener {
	return b.Listener
}

func (b *Box) GetStyle() *brick.Style {
	return b.Style
}

func (b *Box) IsPure() bool {
	return b.Pure
}

func (b *Box) SetPureness(p bool) {
	b.Pure = p
}
