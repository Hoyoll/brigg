package genus

import "github.com/Hoyoll/brick/brick"

func Build(g brick.Genus) *brick.Tree {
	leaf := &brick.Leaf{}
	switch g {
	case brick.Genus_BOX:
		leaf.Genus = brick.Genus_BOX
		leaf.Bone = &Box{
			States: &brick.States{},
		}
	case brick.Genus_TEXT:
		leaf.Genus = brick.Genus_TEXT
		leaf.Bone = &Text{
			States: &brick.States{},
		}
	default:
		leaf.Genus = brick.Genus_BOX
		leaf.Bone = &Box{
			States: &brick.States{},
		}
	}
	tree := &brick.Tree{
		Self: leaf,
	}
	return tree
}
