package main

import (
	"fmt"
	"github.com/music-theory/diagram"
	"github.com/music-theory/note"
)
import "github.com/music-theory/scale"

func main() {

	root := note.Note{Name: "C", Alternate: note.FLAT}

	scale := scale.NewScale(root)
	scale_names := scale.ScaleNames()
	fmt.Println(scale_names)

	scale.ScaleNotes("byzantine")
	scale.SetRoot(root)

	scale.ShowAll()

	fretdiagram := diagram.NewFretBoard()
	fretdiagram.DrawFretBoard()

	inter := scale.ScaleInterval("hinduston")
	fretdiagram.ColorScale(inter)
	fretdiagram.DrawTitle("hinduston", 40, 100)
	fretdiagram.SaveScaleDiagram("fretdiagram.png")
}
