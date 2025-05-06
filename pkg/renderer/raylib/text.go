package raylib

import (
	brigg "github.com/Hoyoll/brigg/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var cachedFont map[string]rl.Font = make(map[string]rl.Font)

type Text struct {
	Width, Height float32
	X, Y          float32
}

func (t *Text) CalcDim(treeid int) {
	tree := brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[tree.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	text, _ := style.GetText()
	font, ok := cachedFont[text.Font]
	if !ok {
		font = rl.LoadFont(text.Font)
		cachedFont[text.Font] = font
	}
	v2 := rl.MeasureTextEx(font, text.Text, text.Sizing, text.Spacing)
	t.Width = v2.X
	t.Height = v2.Y
}

func (b *Text) CalcPos(treeid int) {
	t := brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	cons := style.GetConstraint()

	switch cons.Align {
	case brigg.START:
		offsetX, offsetY := b.X, b.Y
		if cons.Gravity == brigg.HORIZONTAL {
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.X+Cc.PaddingLeft, b.Y+Cc.PaddingTop)
					continue
				}
				offsetX += Cc.PaddingLeft + cons.Gap
				Tree.Renderer.SetPos(offsetX, offsetY+Cc.PaddingTop+cons.Gap)
				w, _ := Tree.Renderer.GetDim()
				offsetX += w + Cc.PaddingRight
			}
		} else {
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.X+Cc.PaddingLeft, b.Y+Cc.PaddingTop)
					continue
				}
				offsetY += Cc.PaddingTop + cons.Gap
				Tree.Renderer.SetPos(offsetX+Cc.PaddingLeft+cons.Gap, offsetY)
				_, y := Tree.Renderer.GetDim()
				offsetY += y + Cc.PaddingBottom
			}
		}
	case brigg.END:
		offsetX, offsetY := b.X+b.Width, b.Y+b.Height
		length := len(t.Branch)
		if cons.Gravity == brigg.HORIZONTAL {
			for i := length - 1; i >= 0; i-- {
				v := t.Branch[i]
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.X+Cc.PaddingLeft, b.Y+Cc.PaddingTop)
					continue
				}
				x, y := Tree.Renderer.GetDim()
				offsetX -= Cc.PaddingRight + x
				Tree.Renderer.SetPos(offsetX, offsetY-(Cc.PaddingBottom+y))
				offsetX -= Cc.PaddingLeft + cons.Gap
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				v := t.Branch[i]
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.X+Cc.PaddingLeft, b.Y+Cc.PaddingTop)
					continue
				}
				x, y := Tree.Renderer.GetDim()
				offsetY -= Cc.PaddingBottom + y
				Tree.Renderer.SetPos(offsetX-(Cc.PaddingRight+x), offsetY)
				offsetY -= cons.Gap + Cc.PaddingTop
			}
		}
	default:
		panic("invalid alignment!")
	}
}

func (t *Text) GetDim() (float32, float32) {
	return t.Width, t.Height
}

func (t *Text) GetPos() (float32, float32) {
	return t.X, t.Y
}

func (t *Text) SetPos(X, Y float32) {
	t.X, t.Y = X, Y
}

func (b *Text) CheckIO(element int, childs []int) (bool, bool) {
	bone := &brigg.Bones.Items[element]
	style := brigg.Styles.Items[bone.GetStyle()]
	constraint := brigg.Constraints.Items[style.Constraint]
	cState := bone.CState

	rec := rl.NewRectangle(b.X, b.Y, b.Width, b.Height)
	mouse := rl.NewVector2(float32(rl.GetMouseX()), float32(rl.GetMouseY()))

	mover := rl.CheckCollisionPointRec(mouse, rec)

	if !constraint.Ghost {
		for _, child := range childs {
			Tree := &brigg.Trees.Items[child]
			x, y := Tree.Renderer.GetPos()
			w, h := Tree.Renderer.GetDim()
			childRec := rl.NewRectangle(x, y, w, h)

			if !rl.CheckCollisionPointRec(mouse, childRec) {
				continue
			}
			mover = false
			break
		}
	}

	if !mover {
		return cState == brigg.DEFAULT, bone.ChangeState(brigg.DEFAULT)
	}

	for state, inputs := range bone.GetIO() {
		if buttonDown(inputs) {
			return cState == state, bone.ChangeState(state)
		}
	}

	return cState == brigg.HOVER, bone.ChangeState(brigg.HOVER)
}

func (t *Text) Render(s int) {
	style := brigg.Styles.Items[s]
	text, _ := style.GetText()
	dst := rl.NewVector2(t.X+t.Width/2, t.Y+t.Height/2)
	origin := rl.NewVector2(t.Width/2, t.Height/2)
	rl.DrawTextPro(cachedFont[text.Font], text.Text,
		dst, origin, text.Rotate, text.Sizing, text.Spacing,
		text.Tint)
}
