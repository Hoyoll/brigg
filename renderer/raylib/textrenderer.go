package raylib

import (
	"github.com/Hoyoll/brick/brick"
	"github.com/Hoyoll/brick/genus"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextRenderer struct {
	Dimension rl.Vector2
	X, Y      float32
	Font      rl.Font
	Box       *rl.Rectangle
	Style     *brick.Style
}

func (t *TextRenderer) Compose(branch *brick.Tree) {
	self := branch.Self.Bone.(*genus.Text)
	t.Style = self.GetStyle()
	t.Font = rl.LoadFont(t.Style.FontPath)
	t.Dimension = rl.MeasureTextEx(t.Font, self.Text, t.Style.FontSize, t.Style.Spacing)
	t.Box = &rl.Rectangle{
		Width:  t.Dimension.X,
		Height: t.Dimension.Y,
	}
}

func (t *TextRenderer) CheckIO(io brick.Root) bool {
	process := io.GetIO()
	normal := true

	if process == nil {
		return normal
	}

	if rl.CheckCollisionPointRec(MousePos, *t.Box) && process.Hover != nil {
		io.ChangeState(brick.HOVER)
		process.Hover(io.GetStyle())
		normal = false
	}

	for key, fun := range process.MouseClick {
		if rl.CheckCollisionPointRec(MousePos, *t.Box) && rl.IsMouseButtonDown(rl.MouseButton(key)) {
			io.ChangeState(brick.CLICK)
			fun(io.GetStyle())
			normal = false
			break
		}
	}

	for key, fun := range process.Key {
		if rl.IsKeyDown(int32(key)) {
			io.ChangeState(brick.KEYDOWN)
			fun(io.GetStyle())
			normal = false
			break
		}
	}
	if normal {
		io.ChangeState(brick.NORMAL)
	}
	return normal
}

func (t *TextRenderer) Position(branch *brick.Tree) {

}

func (t *TextRenderer) Render() {

}

func (t *TextRenderer) GetDimension() (float32, float32) {
	return t.Dimension.X, t.Dimension.Y
}

func (t *TextRenderer) SetPosition(X, Y float32) {
	t.X = X
	t.Y = Y
}

func (t *TextRenderer) GetPosition() (float32, float32) {
	return t.X, t.Y
}
