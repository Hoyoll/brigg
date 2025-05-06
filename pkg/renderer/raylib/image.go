package raylib

import (
	brigg "github.com/Hoyoll/brigg/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var cachedImg map[string]rl.Texture2D = make(map[string]rl.Texture2D)

type Image struct {
	Width, Height float32
	X, Y          float32
}

func (i *Image) CalcDim(treeId int) {
	t := brigg.Trees.Items[treeId]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	img, _ := style.GetImage()
	tex, ok := cachedImg[img.Path]
	if !ok {
		tex = rl.LoadTexture(img.Path)
		cachedImg[img.Path] = tex
	}
	if img.Scale == 0 {
		i.Width = float32(tex.Width)
		i.Height = float32(tex.Height)
	} else {
		i.Width = (img.Scale * float32(tex.Width))
		i.Height = (img.Scale * float32(tex.Height))
	}
}

func (img *Image) CalcPos(treeId int) {
	t := brigg.Trees.Items[treeId]
	bone := brigg.Bones.Items[t.Bones]
	style := brigg.Styles.Items[bone.GetStyle()]
	cons := style.GetConstraint()

	switch cons.Align {
	case brigg.START:
		offsetX, offsetY := img.X, img.Y
		if cons.Gravity == brigg.HORIZONTAL {
			for _, v := range t.Branch {
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(img.X+Cc.PaddingLeft, img.Y+Cc.PaddingTop)
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
					Tree.Renderer.SetPos(img.X+Cc.PaddingLeft, img.Y+Cc.PaddingTop)
					continue
				}
				offsetY += Cc.PaddingTop + cons.Gap
				Tree.Renderer.SetPos(offsetX+Cc.PaddingLeft+cons.Gap, offsetY)
				_, y := Tree.Renderer.GetDim()
				offsetY += y + Cc.PaddingBottom
			}
		}
	case brigg.END:
		offsetX, offsetY := img.X+img.Width, img.Y+img.Height
		length := len(t.Branch)
		if cons.Gravity == brigg.HORIZONTAL {
			for i := length - 1; i >= 0; i-- {
				v := t.Branch[i]
				Tree := &brigg.Trees.Items[v]
				Cb := brigg.Bones.Items[Tree.Bones]
				Cs := brigg.Styles.Items[Cb.GetStyle()]
				Cc := brigg.Constraints.Items[Cs.Constraint]
				if Cc.Static {
					Tree.Renderer.SetPos(img.X+Cc.PaddingLeft, img.Y+Cc.PaddingTop)
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
					Tree.Renderer.SetPos(img.X+Cc.PaddingLeft, img.Y+Cc.PaddingTop)
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

func (i *Image) SetPos(X, Y float32) {
	i.X, i.Y = X, Y
}

func (i *Image) GetDim() (float32, float32) {
	return i.Width, i.Height
}

func (i *Image) GetPos() (float32, float32) {
	return i.X, i.Y
}

func (i *Image) CheckIO(element int, childs []int) (bool, bool) {
	bone := &brigg.Bones.Items[element]
	style := brigg.Styles.Items[bone.GetStyle()]
	constraint := brigg.Constraints.Items[style.Constraint]
	cState := bone.CState

	rec := rl.NewRectangle(i.X, i.Y, i.Width, i.Height)
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

func (i *Image) Render(s int) {
	style := brigg.Styles.Items[s]
	image, _ := style.GetImage()
	img := cachedImg[image.Path]

	width := i.Width
	height := i.Height

	src := rl.NewRectangle(0, 0, float32(img.Width), float32(img.Height))

	dst := rl.NewRectangle(i.X+width/2, i.Y+height/2, width, height)

	origin := rl.NewVector2(width/2, height/2)

	rl.DrawTexturePro(img, src, dst, origin, image.Rotate, image.Tint)
}
