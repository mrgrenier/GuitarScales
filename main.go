package main

import (
	"log"

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
		fretdiagram.DrawDiagram()
		inter := scale.ScaleInterval(scaleName)
		fretdiagram.ColorScale(inter)
		scaleNotes := scale.GetScaleNotes(scaleName)
		fretdiagram.DrawTitle(scaleName, scaleNotes, 40, 100)
		fretdiagram.SaveScaleDiagram("./output/guitar/" + scaleName + ".png")
		pianodiagram := diagram.NewPianoDiagram()
		pianodiagram.DrawDiagram()
		pianodiagram.ColorScale(inter)
		pianodiagram.DrawTitle(scaleName, scaleNotes, 40, 50)
		pianodiagram.SaveScaleDiagram("./output/piano/" + scaleName + ".png")

	}

	guitar := diagram.NewFretBoard()
	if err := guitar.TilePNGsToPDF("./output/guitar/", "./output/guitar_scales.pdf"); err != nil {
		log.Fatal(err)
	}
	piano := diagram.NewPianoDiagram()
	if err := piano.TilePNGsToPDF("./output/piano/", "./output/piano_scales.pdf"); err != nil {
		log.Fatal(err)
	}

}
