package raylib

import (
	"github.com/Hoyoll/brick/brick"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var MousePos rl.Vector2
var WindowWidth int
var WindowHeight int

func Render(branch *brick.Tree) {
	var oldWinHeight int
	var oldWinWidth int
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(500, 500, "nice")
	for !rl.WindowShouldClose() {
		WindowHeight = rl.GetScreenHeight()
		WindowWidth = rl.GetScreenWidth()
		if WindowHeight != oldWinHeight || WindowWidth != oldWinWidth {
			branch.Self.Bone.SetPureness(false)
		}
		oldWinHeight = WindowHeight
		oldWinWidth = WindowWidth
		CalcLayout(branch)
		rl.ClearBackground(rl.Black)
		CalcPosition(branch)
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func GetComposer(genus brick.Genus) brick.Composer {
	var res brick.Composer
	switch genus {
	case brick.Genus_BOX:
		res = &BoxRenderer{}
	case brick.Genus_TEXT:
		res = &TextRenderer{}
	}
	return res
}

func CalcLayout(branch *brick.Tree) {
	if branch.Self.Shadow == nil {
		branch.Self.Shadow = GetComposer(branch.Self.Genus)
	}
	isPure := branch.Self.Bone.IsPure()
	for _, child := range branch.Branch {
		if !isPure {
			child.Self.Bone.SetPureness(isPure)
		}
		CalcLayout(child)
	}
	if !isPure {
		branch.Self.Shadow.Compose(branch)
	}
}

func CalcPosition(branch *brick.Tree) {
	shadow := branch.Self.Shadow
	bone := branch.Self.Bone
	if !bone.IsPure() {
		shadow.Position(branch)
	}
	bone.SetPureness(shadow.CheckIO(bone))
	shadow.Render()
	for _, v := range branch.Branch {
		CalcPosition(v)
	}
}
