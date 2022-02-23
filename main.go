package main

import (
	"github.com/mrgrenier/GuitarScales/diagram"
	"github.com/mrgrenier/GuitarScales/note"
	"github.com/mrgrenier/GuitarScales/scale"
)

func main() {

	root := note.Note{Name: "C", Alternate: note.FLAT}

	scale := scale.NewScale(root)
	scale_names := scale.ScaleNames()
	for _, scaleName := range scale_names {
		fretdiagram := diagram.NewFretBoard()
		fretdiagram.DrawFretBoard()
		inter := scale.ScaleInterval(scaleName)
		fretdiagram.ColorScale(inter)
		scaleNotes := scale.GetScaleNotes(scaleName)
		fretdiagram.DrawTitle(scaleName, scaleNotes, 40, 100)
		fretdiagram.SaveScaleDiagram("./output/" + scaleName + ".png")

	}

}
