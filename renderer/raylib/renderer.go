package raylib

import (
	"github.com/Hoyoll/brigg/brigg"
	"github.com/Hoyoll/brigg/lib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Raylib struct {
	WinX int
	WinY int
}

var Rectangles *lib.Barrel[rl.Rectangle]

func (r *Raylib) Render(tree *brigg.Tree, buffer int) {
	brigg.Composers = &lib.Barrel[brigg.Composer]{
		Items: make([]brigg.Composer, 1, buffer),
	}
	Rectangles = &lib.Barrel[rl.Rectangle]{
		Items: make([]rl.Rectangle, 1, buffer),
	}
	var tempX, tempY int
	pure := true
	rl.InitWindow(int32(r.WinX), int32(r.WinY), "test")
	for !rl.WindowShouldClose() {
		tempX = rl.GetScreenWidth()
		tempY = rl.GetScreenHeight()
		if tempX != r.WinX || tempY != r.WinY {
			pure = false
			r.WinX = tempX
			r.WinY = tempY
		}
		r.calcDim(tree, pure)
		r.calcPos(tree, pure)
	}
	rl.CloseWindow()
}

func (r *Raylib) calcDim(tree *brigg.Tree, pure bool) {
	for _, i := range tree.Branch {
		branch := &brigg.Trees.Items[i]
		if !pure {
			branch.Pure = pure
		}
		r.calcDim(branch, branch.Pure)
	}
	if pure {
		return
	}
	self := &brigg.Leaves.Items[tree.Self]
	composerId, ok := self.Renderer[tree.Genus]
	if !ok {
		composerId = r.GetComposer(tree.Genus)
		self.Renderer[tree.Genus] = composerId
	}
	composer := brigg.Composers.Items[composerId]
	composer.CalcDim(self.Bone, tree.Parent, tree.Branch)
}

func (r *Raylib) calcPos(tree *brigg.Tree, pure bool) {

}

func (r *Raylib) GetComposer(g brigg.Genus) int {
	var res brigg.Composer
	switch g {
	case brigg.BOX:
		res = &Box{}
	case brigg.TEXT:
	}
	return brigg.Composers.Add(res)
}
