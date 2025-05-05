package raylib

import (
	brigg "github.com/Hoyoll/brigg/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Box struct {
	Width, Height float32
	X, Y          float32
}

func (b *Box) CalcDim(treeid int) {
	t := brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	box, _ := style.GetBox()
	cons := style.GetConstraint()

	switch cons.Gravity {
	case brigg.HORIZONTAL:
		switch box.Height {
		case brigg.WRAP_CONTENT:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					continue
				}
				_, tempY := Tree.Renderer.GetDim()
				tallest := tempY + (Cc.PaddingTop + Cc.PaddingBottom) + cons.Gap
				if b.Height < tallest {
					b.Height = tallest
				}
			}
			b.Height += cons.Gap
		case brigg.WINDOW:
			b.Height = float32(rl.GetScreenHeight())
		default:
			b.Height = box.Height
		}
		switch box.Width {
		case brigg.WRAP_CONTENT:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					continue
				}
				tempW, _ := Tree.Renderer.GetDim()
				b.Width += tempW + (Cc.PaddingLeft + Cc.PaddingRight) + cons.Gap
			}
			b.Width += cons.Gap
		case brigg.WINDOW:
			b.Width = float32(rl.GetScreenWidth())
		default:
			b.Width = box.Width
		}
	case brigg.VERTICAL:
		switch box.Height {
		case brigg.WRAP_CONTENT:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					continue
				}
				_, tempY := Tree.Renderer.GetDim()
				b.Height += tempY + (Cc.PaddingTop + Cc.PaddingBottom) + cons.Gap
			}
			b.Height += cons.Gap
		case brigg.WINDOW:
			b.Height = float32(rl.GetScreenHeight())
		default:
			b.Height = box.Height
		}
		switch box.Width {
		case brigg.WRAP_CONTENT:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					continue
				}
				tempW, _ := Tree.Renderer.GetDim()
				widest := tempW + (Cc.PaddingLeft + Cc.PaddingRight) + cons.Gap
				if b.Width < widest {
					b.Width = widest
				}
			}
			b.Width += cons.Gap
		case brigg.WINDOW:
			b.Width = float32(rl.GetScreenWidth())
		default:
			b.Width = box.Width
		}
	}
}

func (b *Box) CalcPos(treeid int) {
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
				offsetX -= Cc.PaddingRight + cons.Gap
				Tree.Renderer.SetPos(offsetX, offsetY-Cc.PaddingBottom+cons.Gap)
				x, _ := Tree.Renderer.GetDim()
				offsetX -= x + Cc.PaddingLeft
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
				offsetY -= Cc.PaddingBottom + cons.Gap
				Tree.Renderer.SetPos(offsetX-Cc.PaddingRight+cons.Gap, offsetY)
				_, y := Tree.Renderer.GetDim()
				offsetY -= y + Cc.PaddingTop
			}
		}
	default:
		panic("invalid alignment!")
	}
}

func (b *Box) SetPos(X, Y float32) {
	b.X, b.Y = X, Y
}

func (b *Box) CheckIO(element int, childs []int) (bool, bool) {
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

func (b *Box) GetDim() (float32, float32) {
	return b.Width, b.Height
}

func (b *Box) GetPos() (float32, float32) {
	return b.X, b.Y
}

func (b *Box) Render(s int) {
	style := brigg.Styles.Items[s]
	box, _ := style.GetBox()
	rec := rl.NewRectangle(b.X, b.Y, b.Width, b.Height)
	rl.DrawRectangleRounded(rec, box.Radius, 10, box.Color)
}

// func unpack(tree int) (brigg.Constraint, brigg.Composer) {
// 	Tree := brigg.Trees.Items[tree]
// 	Cb := brigg.Bones.Items[Tree.Bones]
// 	Cr := Tree.Renderer
// 	Cs := brigg.Styles.Items[Cb.GetStyle()]
// 	Cc := brigg.Constraints.Items[Cs.Constraint]
// 	return Cc, Cr
// }
