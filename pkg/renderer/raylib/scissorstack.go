package raylib

import rl "github.com/gen2brain/raylib-go/raylib"

func scissorBuild() scissor {
	return scissor{}
}

type scissor struct {
	Scissors int
}

func (s *scissor) Begin(rec rl.Rectangle) int {
	s.Scissors++
	rl.BeginScissorMode(int32(rec.X), int32(rec.Y),
		int32(rec.Width), int32(rec.Height))
	return s.Scissors
}

func (s *scissor) End(id int) {
	s.Scissors--
	if s.Scissors == 0 {
		rl.EndScissorMode()
	}
}
