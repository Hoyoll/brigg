package raylib

import (
	"github.com/Hoyoll/brick/brick"
	"github.com/Hoyoll/brick/genus"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BoxRenderer struct {
	Box   *rl.Rectangle
	Style *brick.Style
}

func (b *BoxRenderer) Compose(branch *brick.Tree) {
	if b.Box == nil {
		b.Box = &rl.Rectangle{}
	}
	b.Style = branch.Self.Bone.GetStyle()
	b.Box.Height = b.SetHeight(branch)
	b.Box.Width = b.SetWidth(branch)
}

func (b *BoxRenderer) CheckIO(io brick.Root) bool {
	process := io.GetIO()
	normal := true

	if process == nil {
		return normal
	}

	if rl.CheckCollisionPointRec(MousePos, *b.Box) && process.Hover != nil {
		io.ChangeState(brick.HOVER)
		process.Hover(io.GetStyle())
		normal = false
	}

	for key, fun := range process.MouseClick {
		if rl.CheckCollisionPointRec(MousePos, *b.Box) && rl.IsMouseButtonPressed(rl.MouseButton(key)) {
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
		style := io.GetStyle()
		for _, fun := range process.Normal {
			fun(style)
			normal = false
		}
	}
	return normal
}

func (b *BoxRenderer) Position(branch *brick.Tree) {
	if branch.Parent == nil {
		b.SetPosition(0, 0)
	}

	selfStyle := branch.Self.Bone.GetStyle()
	selfX, selfY := branch.Self.Shadow.GetPosition()

	switch selfStyle.Gravity {
	case brick.Gravity_HORIZONTAL:
		if selfStyle.Align == brick.Align_START {
			tempX := selfX
			for _, v := range branch.Branch {
				shadow := v.Self.Shadow
				style := v.Self.Bone.GetStyle()
				tempX += style.PaddingX
				shadow.SetPosition(tempX, selfY+style.PaddingY)
				tempX += style.PaddingX + selfStyle.Gap
			}
		} else {
			// tempY := selfY
			// for _, v := range branch.Branch {
			// 	shadow := v.Self.Shadow
			// 	style := v.Self.Bone.GetStyle()
			// }
		}
	case brick.Gravity_VERTICAL:
		if selfStyle.Align == brick.Align_START {
			tempY := selfY
			for _, v := range branch.Branch {
				shadow := v.Self.Shadow
				style := v.Self.Bone.GetStyle()
				tempY += style.PaddingY
				shadow.SetPosition(selfX+style.PaddingX, tempY)
				tempY += style.PaddingY + selfStyle.Gap
			}
		} else {

		}
	}

}

func (b *BoxRenderer) SetHeight(branch *brick.Tree) float32 {
	self := branch.Self.Bone.(*genus.Box)
	style := self.Style

	getParentHeight := func() float32 {
		if branch.Parent == nil {
			return float32(rl.GetScreenHeight())
		}
		parentHeight := float32(branch.Parent.Bone.GetStyle().Height)
		if float32(style.MaxHeight) < parentHeight {
			return parentHeight
		}
		return float32(style.MaxHeight)
	}

	if self.Style.Gravity == brick.Gravity_HORIZONTAL {
		switch self.Style.Height {
		case brick.Size_CHILD_SUM:
			var maxChildHeight float32
			for _, v := range branch.Branch {
				_, h := v.Self.Shadow.GetDimension()
				padY := v.Self.Bone.GetStyle().PaddingY
				totalHeight := h + float32(padY)*2 + float32(style.Gap)
				if totalHeight > maxChildHeight {
					maxChildHeight = totalHeight
				}
			}
			return maxChildHeight + float32(style.Gap)

		case brick.Size_MATCH_PARENT:
			return getParentHeight()

		default:
			return float32(style.Height)
		}
	}

	switch style.Height {
	case brick.Size_CHILD_SUM:
		var total float32
		for _, v := range branch.Branch {
			_, h := v.Self.Shadow.GetDimension()
			padY := v.Self.Bone.GetStyle().PaddingY
			total += float32(padY)*2 + h + float32(style.Gap)
		}
		return total + float32(style.Gap)

	case brick.Size_MATCH_PARENT:
		return getParentHeight()

	default:
		return float32(style.Height)
	}
}

func (b *BoxRenderer) SetWidth(branch *brick.Tree) float32 {
	self := branch.Self.Bone.(*genus.Box)
	style := self.Style

	getParentWidth := func() float32 {
		if branch.Parent == nil {
			return float32(rl.GetScreenWidth())
		}
		parentWidth := float32(branch.Parent.Bone.GetStyle().Width)
		if float32(style.MaxWidth) < parentWidth {
			return parentWidth
		}
		return float32(style.MaxWidth)
	}

	if style.Gravity == brick.Gravity_VERTICAL {
		switch style.Width {
		case brick.Size_CHILD_SUM:
			var maxChildWidth float32
			for _, v := range branch.Branch {
				w, _ := v.Self.Shadow.GetDimension()
				padX := v.Self.Bone.GetStyle().PaddingX
				totalWidth := w + float32(padX)*2 + float32(style.Gap)
				if totalWidth > maxChildWidth {
					maxChildWidth = totalWidth
				}
			}
			return maxChildWidth + float32(style.Gap)

		case brick.Size_MATCH_PARENT:
			return getParentWidth()

		default:
			return float32(style.Width)
		}
	}

	switch style.Width {
	case brick.Size_CHILD_SUM:
		var total float32
		for _, v := range branch.Branch {
			w, _ := v.Self.Shadow.GetDimension()
			padX := v.Self.Bone.GetStyle().PaddingX
			total += float32(padX)*2 + w + float32(style.Gap)
		}
		return total + float32(style.Gap)

	case brick.Size_MATCH_PARENT:
		return getParentWidth()

	default:
		return float32(style.Width)
	}
}

func (b *BoxRenderer) Render() {
	rl.DrawRectangleRounded(*b.Box, b.Style.Radius, 5, *b.Style.Color)
}

func (b *BoxRenderer) GetDimension() (float32, float32) {
	return b.Box.Width, b.Box.Height
}

func (b *BoxRenderer) SetPosition(X, Y float32) {
	b.Box.X = X
	b.Box.Y = Y
}

func (b *BoxRenderer) GetPosition() (float32, float32) {
	return b.Box.X, b.Box.Y
}
