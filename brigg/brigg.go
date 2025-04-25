package brigg

import "github.com/Hoyoll/brigg/lib"

var (
	Trees *lib.Barrel[Tree] = &lib.Barrel[Tree]{
		Items: make([]Tree, 1000),
	}

	Leaves *lib.Barrel[Leaf] = &lib.Barrel[Leaf]{
		Items: make([]Leaf, 1000),
	}
	Bones *lib.Barrel[Root] = &lib.Barrel[Root]{
		Items: make([]Root, 1000),
	}
	Styles *lib.Barrel[Style] = &lib.Barrel[Style]{
		Items: make([]Style, 1000),
	}
	Listens *lib.Barrel[Listen] = &lib.Barrel[Listen]{
		Items: make([]Listen, 1000),
	}
	States *lib.Barrel[StateMap] = &lib.Barrel[StateMap]{
		Items: make([]StateMap, 1000),
	}
)

func Build(g Genus) *Tree {
	var r Root
	switch g {
	case BOX:
		r = &Box{}
	default:
		r = &Box{}
	}
	el := Bones.Add(r)
	leaf := Leaf{
		Bone:  el,
		Genus: g,
	}
	tree := Tree{
		Self: Leaves.Add(leaf),
	}
	id := Trees.Add(tree)
	return &Trees.Items[id]
}

// func main() {
// 	b := Build[Box](BOX)
// }
