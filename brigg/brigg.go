package brigg

import "github.com/Hoyoll/brigg/lib"

var (
	Trees *lib.Barrel[Tree] = &lib.Barrel[Tree]{
		Items: make([]Tree, 1, 1000),
	}
	Leaves *lib.Barrel[Leaf] = &lib.Barrel[Leaf]{
		Items: make([]Leaf, 1, 1000),
	}
	Bones *lib.Barrel[Element] = &lib.Barrel[Element]{
		Items: make([]Element, 1, 1000),
	}
	Styles *lib.Barrel[Style] = &lib.Barrel[Style]{
		Items: make([]Style, 1, 1000),
	}
	States *lib.Barrel[StateMap] = &lib.Barrel[StateMap]{
		Items: make([]StateMap, 1, 1000),
	}
)

var (
	Boxes *lib.Barrel[Box] = &lib.Barrel[Box]{
		Items: make([]Box, 1, 1000),
	}
	Images *lib.Barrel[Image] = &lib.Barrel[Image]{
		Items: make([]Image, 1, 1000),
	}
	Texts *lib.Barrel[Text] = &lib.Barrel[Text]{
		Items: make([]Text, 1, 1000),
	}
)

var Composers *lib.Barrel[Composer]

func Build(g Genus) *Tree {
	leaf := Leaf{
		Bone: Bones.Add(Element{DEFID, DEFID}),
	}
	tree := Tree{
		Genus:  g,
		Self:   Leaves.Add(leaf),
		Parent: DEFID,
	}
	id := Trees.Add(tree)
	return &Trees.Items[id]
}
