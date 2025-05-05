package raylib

import (
	"github.com/Hoyoll/brigg/lib"
	brigg "github.com/Hoyoll/brigg/pkg"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Renderer(width, height int, title string) Raylib {
	return Raylib{
		WinX:  width,
		WinY:  height,
		Title: title,
	}
}

type Raylib struct {
	WinX  int
	WinY  int
	Title string
}

var Rectangles *lib.Barrel[rl.Rectangle]
var Textures *lib.Barrel[rl.RenderTexture2D]

func (r *Raylib) Render(tree, buffer int) {
	brigg.Composers = &lib.Barrel[brigg.Composer]{
		Items: make([]brigg.Composer, 0, buffer),
	}
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(int32(r.WinX), int32(r.WinY), r.Title)
	var tempX, tempY int
	for !rl.WindowShouldClose() {
		cW, cH := rl.GetScreenWidth(), rl.GetScreenHeight()
		pure := tempX == cW && tempY == cH
		if !pure {
			r.calcDim(tree, false)
			tempX = cW
			tempY = cH
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		rl.DrawRectangle(100, 100, 200, 100, rl.Blue)
		r.render(tree)
		rl.EndDrawing()
	}
	rl.CloseWindow()
	for _, v := range cachedFont {
		rl.UnloadFont(v)
	}
	for _, v := range cachedImg {
		rl.UnloadTexture(v)
	}
}

func (r *Raylib) calcDim(treeid int, pure bool) {
	tree := &brigg.Trees.Items[treeid]
	for _, i := range tree.Branch {
		branch := &brigg.Trees.Items[i]
		if !pure {
			branch.Pure = pure
		}
		r.calcDim(i, branch.Pure)
	}
	if pure {
		return
	}
	tree.Renderer = r.GetComposer(tree.Genus)
	tree.Renderer.CalcDim(treeid)
}

func (r *Raylib) render(treeid int) {
	tree := &brigg.Trees.Items[treeid]
	bone := brigg.Bones.Items[tree.Bones]
	if !tree.Pure {
		tree.Renderer.CalcPos(treeid)
	}
	tree.Renderer.Render(bone.GetStyle())
	for _, i := range tree.Branch {
		r.render(i)
	}
	s, f := tree.Renderer.CheckIO(tree.Bones, tree.Branch)
	tree.Pure = s && f
}

func (r *Raylib) GetComposer(g brigg.Genus) brigg.Composer {
	var res brigg.Composer
	switch g {
	case brigg.BOX:
		res = &Box{}
	case brigg.TEXT:
		res = &Text{}
	case brigg.IMAGE:
		res = &Image{}
	}
	return res
}

func buttonDown(buttons map[brigg.Event][]any) bool {
	for e, v := range buttons {
		switch e {
		case brigg.DOWN:
			for _, key := range v {
				switch v := key.(type) {
				case int:
					if !rl.IsKeyDown(int32(v)) {
						return false
					}
				case rl.MouseButton:
					if !rl.IsMouseButtonDown(v) {
						return false
					}
				default:
					return false
				}
			}
		case brigg.UP:
			for _, key := range v {
				switch v := key.(type) {
				case int:
					if !rl.IsKeyUp(int32(v)) {
						return false
					}
				case rl.MouseButton:
					if !rl.IsMouseButtonUp(v) {
						return false
					}
				default:
					return false
				}
			}
		case brigg.PRESSED:
			for _, key := range v {
				switch v := key.(type) {
				case int:
					if !rl.IsKeyPressed(int32(v)) {
						return false
					}
				case rl.MouseButton:
					if !rl.IsMouseButtonPressed(v) {
						return false
					}
				default:
					return false
				}
			}
		case brigg.RELEASED:
			for _, key := range v {
				switch v := key.(type) {
				case int:
					if !rl.IsKeyReleased(int32(v)) {
						return false
					}
				case rl.MouseButton:
					if !rl.IsMouseButtonReleased(v) {
						return false
					}
				default:
					return false
				}
			}
		}
	}
	return true
}
