package brigg

import "github.com/Hoyoll/brigg/lib"

var (
	Trees *lib.Barrel[Tree] = &lib.Barrel[Tree]{
		Items: make([]Tree, 0, 100),
	}
	Bones *lib.Barrel[Element] = &lib.Barrel[Element]{
		Items: []Element{},
	}
	Styles *lib.Barrel[Style] = &lib.Barrel[Style]{
		Items: []Style{},
	}
	States *lib.Barrel[StateMap] = &lib.Barrel[StateMap]{
		Items: []StateMap{},
	}
)

var (
	Boxes *lib.Barrel[Box] = &lib.Barrel[Box]{
		Items: []Box{},
	}
	Images *lib.Barrel[Image] = &lib.Barrel[Image]{
		Items: []Image{},
	}
	Texts *lib.Barrel[Text] = &lib.Barrel[Text]{
		Items: []Text{},
	}
	Constraints *lib.Barrel[Constraint] = &lib.Barrel[Constraint]{
		Items: []Constraint{},
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
