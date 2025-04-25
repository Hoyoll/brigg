package raylib

import (
	"github.com/Hoyoll/brigg/brigg"
	"github.com/Hoyoll/brigg/lib"
)

var Composers *lib.Barrel[brigg.Composer]

func Render(buffer int) {
	Composers = &lib.Barrel[brigg.Composer]{
		Items: make([]brigg.Composer, buffer),
	}
}
