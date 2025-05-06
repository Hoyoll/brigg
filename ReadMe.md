# Brigg

Oops, all Go. Meet **Brigg**, the unintentional Go-based UI layout thingy. It’s simple, flexible, and yes, it actually works. Not sure how I got here, but here we are.

## What is this?

So, this is a UI rendering engine written in Go. It's like a declarative layout system that responds to screen sizes, mouse movements, and all that fun stuff. You can use it to create UIs, and it even supports things like hover states and other state. (I swear I didn’t mean to make it this complicated.)

### Features:

- **Declarative Layouts**: Just define stuff in Go structs and functions. No magic—just Go.
- **Renderer Agnostic**: It shipped with Raylib, but you can swap it for whatever renderer fits your fancy. It’s all decoupled, which means it’s kinda future-proof (I hope).
- **Constraints**: Padding, gaps, alignment. You know, all the layout-y things.

### Warning!:
- It is very low to the ground. This is NOT a UI library, i don't provide high level abstraction (not yet anyway) for let's say a button component. What you will get is a way to make that button. What is UI anyway? It just a bunch of rectangle all the way down

- What you do get is a bunch of primitives. A way to make boxes, images, and text. That's it  

## How to Use

1. **Install it** (don’t blame me, okay?):

```bash
go get github.com/Hoyoll/brigg
```

2. **Example**

```
package main

import (
	"github.com/Hoyoll/brigg/brigg"
	"github.com/Hoyoll/brigg/renderer/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func gchild() int {
	root, id := brigg.Build(brigg.BOX)
	bone := root.Bone()
	state := bone.State()
	style := brigg.NewStyle()
	style.SetBox(brigg.Box{
		Width:  30,
		Height: 30,
		Color:  rl.Blue,
	})
	style.SetConstraint(brigg.Constraint{
		PaddingTop:   10,
		PaddingRight: 10,
		PaddingLeft:  10,
	})
	state.Add(brigg.DEFAULT, style)
	return id
}

func child() int {
	root, id := brigg.Build(brigg.BOX)
	bone := root.Bone()
	state := bone.State()
	style := brigg.NewStyle()
	constraint := brigg.Constraint{
		PaddingRight: 10,
		PaddingLeft:  10,
	}
	style.SetBox(brigg.Box{
		Width:  100,
		Height: 100,
		Color:  rl.Red,
	})
	style.SetConstraint(constraint)
	hstyle := brigg.NewStyle()
	hstyle.SetBox(brigg.Box{
		Width:  100,
		Height: 100,
		Color:  rl.Blue,
	})
	hstyle.SetConstraint(constraint)
	state.Add(brigg.DEFAULT, style)
	state.Add(brigg.HOVER, hstyle)(func(s *brigg.Style) bool {
		fmt.Print("Hovered!")
		return false
	})
	root.Child(gchild(), gchild())
	return id
}

func main() {
	root, id := brigg.Build(brigg.BOX)
	bone := root.Bone()
	state := bone.State()
	style := brigg.NewStyle()
	style.SetBox(brigg.Box{
		Width:  brigg.WINDOW,
		Height: brigg.WINDOW,
		Color:  rl.Yellow,
		Rotate: 20,
	})
	style.SetConstraint(brigg.Constraint{
		Gravity: brigg.HORIZONTAL,
		Align:   brigg.START,
		Gap:     12,
	})
	state.Add(brigg.DEFAULT, style)
	root.Child(child(), child(), child())
	renderer := raylib.Renderer(800, 600, "It Works!")
	renderer.Render(id, 20)
}
```