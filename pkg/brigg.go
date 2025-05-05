package brigg

import "github.com/Hoyoll/brigg/lib"

var (
	Trees *lib.Barrel[Tree] = &lib.Barrel[Tree]{
		Items: make([]Tree, 0, 1000),
	}
	Bones *lib.Barrel[Element] = &lib.Barrel[Element]{
		Items: make([]Element, 0, 1000),
	}
	Styles *lib.Barrel[Style] = &lib.Barrel[Style]{
		Items: make([]Style, 0, 1000),
	}
	States *lib.Barrel[StateMap] = &lib.Barrel[StateMap]{
		Items: make([]StateMap, 0, 1000),
	}
)

var (
	Boxes *lib.Barrel[Box] = &lib.Barrel[Box]{
		Items: make([]Box, 0, 1000),
	}
	Images *lib.Barrel[Image] = &lib.Barrel[Image]{
		Items: make([]Image, 0, 1000),
	}
	Texts *lib.Barrel[Text] = &lib.Barrel[Text]{
		Items: make([]Text, 0, 1000),
	}
	Constraints *lib.Barrel[Constraint] = &lib.Barrel[Constraint]{
		Items: make([]Constraint, 1, 1000),
	}
)

var Composers *lib.Barrel[Composer]

func Build(g Genus) (*Tree, int) {
	tree := Tree{
		Genus: g,
		Bones: Bones.Add(Element{DEFID, DEFID}),
	}
	id := Trees.Add(tree)
	return &Trees.Items[id], id
}
