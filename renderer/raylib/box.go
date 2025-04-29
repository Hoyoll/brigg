package raylib

import (
	"github.com/Hoyoll/brigg/brigg"
)

type Box struct {
	Rectangle int
	Style     int
}

func (b *Box) getSum() {

}

func (b *Box) getBiggest() {

}

func (b *Box) CalcDim(s, p int, c []int) {

}

func (b *Box) CalcPos(s, p int, c []int) {

}

func (b *Box) CheckIO(r *brigg.Element) bool {
	return true
}

func (b *Box) GetDim() (float32, float32) {
	r := Rectangles.Items[b.Rectangle]
	return r.Width, r.Height
}

func (b *Box) GetPos() (float32, float32) {
	r := Rectangles.Items[b.Rectangle]
	return r.X, r.Y
}

func (b *Box) Render() {
	// r := Rectangles.Items[b.Rectangle]
	// rl.DrawRectangleRounded(r)
}
