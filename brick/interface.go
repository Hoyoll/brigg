package brick

type Root interface {
	IsPure() bool
	SetPureness(bool)
	GetStyle() *Style
	SetStyle(*Style)
	GetIO() *Listener
	ChangeState(StateFlag)
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
