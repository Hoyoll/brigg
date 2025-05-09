package raylib

import (
	brigg "github.com/Hoyoll/brigg/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var scissorStack scissor = scissorBuild()

type Box struct {
	Rec rl.Rectangle
	// Width, Height float32
	// X, Y          float32
	RowOrColumn float32
	Bound       int
}

// TO-DO: Refractor this piece of sh*t
func (b *Box) CalcDim(treeid int) {
	t := brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	box, _ := style.GetBox()
	cons := style.GetConstraint()
	b.Rec = rl.Rectangle{}
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
				tallest :=
					tempY + (Cc.PaddingTop + Cc.PaddingBottom) + cons.Gap
				if b.Rec.Height < tallest {
					b.Rec.Height = tallest
				}
			}
			b.Rec.Height += cons.Gap
		case brigg.WINDOW:
			b.Rec.Height = float32(rl.GetScreenHeight())
		default:
			b.Rec.Height = box.Height
		}
		b.RowOrColumn = b.Rec.Height
		switch box.Width {
		case brigg.WRAP_CONTENT:
			var totalWidth float32 = 0
			var perOverflow bool = false
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					continue
				}
				w, _ := Tree.Renderer.GetDim()
				tempWidth :=
					w + (Cc.PaddingLeft + Cc.PaddingRight) + cons.Gap

				if box.MaxWidth == 0 {
					b.Rec.Width += tempWidth
				}

				totalWidth += tempWidth
				overflowed := totalWidth > box.MaxWidth

				if overflowed {
					perOverflow = true
					totalWidth = 0
					b.Rec.Width = box.MaxWidth
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						return
					case brigg.WRAP:
						b.Rec.Height += b.RowOrColumn
						continue
					}
				}

				if !perOverflow {
					b.Rec.Width += tempWidth
				}

			}
			b.Rec.Width += cons.Gap
		case brigg.WINDOW:
			b.Rec.Width = float32(rl.GetScreenWidth())
		default:
			b.Rec.Width = box.Width
		}
	case brigg.VERTICAL:
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
				widest :=
					tempW + (Cc.PaddingLeft + Cc.PaddingRight) + cons.Gap
				if b.Rec.Width < widest {
					b.Rec.Width = widest
				}
			}
			b.Rec.Width += cons.Gap
		case brigg.WINDOW:
			b.Rec.Width = float32(rl.GetScreenWidth())
		default:
			b.Rec.Width = box.Width
		}
		b.RowOrColumn = b.Rec.Width
		switch box.Height {
		case brigg.WRAP_CONTENT:
			var totalHeight float32
			var perOverflow bool = false
			for _, v := range t.Branch {
				tree := &brigg.Trees.Items[v]
				cb := brigg.Bones.Items[tree.Bones]
				cs := brigg.Styles.Items[cb.GetStyle()]
				cc := brigg.Constraints.Items[cs.Constraint]

				if cc.Static {
					continue
				}

				_, h := tree.Renderer.GetDim()
				tempHeight :=
					h + cc.PaddingTop + cc.PaddingBottom + cons.Gap

				if box.MaxHeight == 0 {
					b.Rec.Height += tempHeight
					continue
				}

				totalHeight += tempHeight
				overflowed := totalHeight > box.MaxHeight
				if overflowed {
					perOverflow = true
					totalHeight = 0
					b.Rec.Height = box.MaxHeight
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						return
					case brigg.WRAP:
						b.Rec.Width += b.RowOrColumn
						continue
					}
				}

				if !perOverflow {
					b.Rec.Height += tempHeight
				}
			}
			b.Rec.Height += cons.Gap
		case brigg.WINDOW:
			b.Rec.Height = float32(rl.GetScreenHeight())

		default:
			b.Rec.Height = box.Height
		}
	}
}

func (b *Box) CalcPos(treeid int) {
	t := brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	cons := style.GetConstraint()
	box, _ := style.GetBox()

	switch cons.Align {
	case brigg.START:
		offsetX, offsetY :=
			b.Rec.X+box.ContentOffsetX, b.Rec.Y+box.ContentOffsetY
		switch cons.Gravity {
		case brigg.HORIZONTAL:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.Rec.X+Cc.PaddingLeft,
						b.Rec.Y+Cc.PaddingTop)
					continue
				}
				offsetX += Cc.PaddingLeft + cons.Gap
				Tree.Renderer.SetPos(offsetX,
					offsetY+Cc.PaddingTop+cons.Gap)
				w, _ := Tree.Renderer.GetDim()
				offsetX += w + Cc.PaddingRight
				if box.MaxWidth == 0 {
					continue
				}
				inside :=
					rl.CheckCollisionPointRec(
						rl.NewVector2(offsetX, offsetY),
						b.Rec)
				if !inside {
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						continue
					case brigg.WRAP:
						offsetX = b.Rec.X + box.ContentOffsetX
						offsetY += b.RowOrColumn
					}
				}
			}
		case brigg.VERTICAL:
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.Rec.X+Cc.PaddingLeft,
						b.Rec.Y+Cc.PaddingTop)
					continue
				}
				offsetY += Cc.PaddingTop + cons.Gap
				Tree.Renderer.SetPos(offsetX+Cc.PaddingLeft+cons.Gap,
					offsetY)
				_, y := Tree.Renderer.GetDim()
				offsetY += y + Cc.PaddingBottom
				if box.MaxHeight == 0 {
					continue
				}
				inside :=
					rl.CheckCollisionPointRec(
						rl.NewVector2(offsetX, offsetY),
						b.Rec)
				if !inside {
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						break
					case brigg.WRAP:
						offsetY = b.Rec.Y + box.ContentOffsetY
						offsetX += b.RowOrColumn
					}
				}
			}
		}
	case brigg.END:
		offsetX, offsetY :=
			(b.Rec.X+b.Rec.Width)+box.ContentOffsetX,
			(b.Rec.Y+b.Rec.Height)+box.ContentOffsetY
		length := len(t.Branch)
		switch cons.Gravity {
		case brigg.HORIZONTAL:
			for i := length - 1; i >= 0; i-- {
				v := t.Branch[i]
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.Rec.X+Cc.PaddingLeft,
						b.Rec.Y+Cc.PaddingTop)
					continue
				}
				x, y := Tree.Renderer.GetDim()
				offsetX -= Cc.PaddingRight + x + cons.Gap
				Tree.Renderer.SetPos(offsetX,
					offsetY-(Cc.PaddingBottom+y))
				offsetX -= Cc.PaddingLeft + cons.Gap
				if box.MaxWidth == 0 {
					continue
				}

				inside :=
					rl.CheckCollisionPointRec(
						rl.NewVector2(offsetX, offsetY),
						b.Rec)
				if !inside {
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						continue
					case brigg.WRAP:
						offsetX = b.Rec.X + b.Rec.Width + box.ContentOffsetX
						offsetY -= b.RowOrColumn
					}
				}

			}
		case brigg.VERTICAL:
			for i := length - 1; i >= 0; i-- {
				v := t.Branch[i]
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(b.Rec.X+Cc.PaddingLeft,
						b.Rec.Y+Cc.PaddingTop)
					continue
				}
				x, y := Tree.Renderer.GetDim()
				offsetY -= Cc.PaddingBottom + y + cons.Gap
				Tree.Renderer.SetPos(offsetX-(Cc.PaddingRight+x),
					offsetY)
				offsetY -= cons.Gap + Cc.PaddingTop
				if box.MaxHeight == 0 {
					continue
				}

				inside :=
					rl.CheckCollisionPointRec(
						rl.NewVector2(offsetX, offsetY),
						b.Rec)

				if !inside {
					switch box.Overflow {
					case brigg.HIDE, brigg.LEAK:
						continue
					case brigg.WRAP:
						offsetY = b.Rec.Y + b.Rec.Height + box.ContentOffsetY
						offsetX -= b.RowOrColumn
					}
				}
			}
		}
	default:
		panic("invalid alignment!")
	}
}

func (b *Box) SetPos(X, Y float32) {
	b.Rec.X, b.Rec.Y = X, Y
}

func (b *Box) CheckIO(element int, childs []int) (bool, bool) {
	bone := &brigg.Bones.Items[element]
	style := brigg.Styles.Items[bone.GetStyle()]
	constraint := brigg.Constraints.Items[style.Constraint]
	cState := bone.CState

	mouse := rl.NewVector2(float32(rl.GetMouseX()),
		float32(rl.GetMouseY()))

	mover := rl.CheckCollisionPointRec(mouse, b.Rec)
	if b.Bound != 0 {
		scissorStack.End(b.Bound)
	}

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

	scroll := rl.GetMouseWheelMoveV()

	if scroll.Y != 0 {
		if scroll.Y > 0 {
			return cState == brigg.SCROLL_UP, bone.ChangeState(brigg.SCROLL_UP)
		} else {
			return cState == brigg.SCROLL_DOWN, bone.ChangeState(brigg.SCROLL_DOWN)
		}
	}

	for state, inputs := range bone.GetIO() {
		if buttonDown(inputs) {
			return cState == state, bone.ChangeState(state)
		}
	}

	return cState == brigg.HOVER, bone.ChangeState(brigg.HOVER)
}

func (b *Box) GetDim() (float32, float32) {
	return b.Rec.Width, b.Rec.Height
}

func (b *Box) GetPos() (float32, float32) {
	return b.Rec.X, b.Rec.Y
}

func (b *Box) Render(s int) {
	style := brigg.Styles.Items[s]
	box, _ := style.GetBox()
	if box.Overflow == brigg.HIDE {
		b.Bound = scissorStack.Begin(b.Rec)
	}
	rl.DrawRectangleRounded(b.Rec, box.Radius, 10, box.Color)
}
