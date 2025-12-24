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
		inter := scale.ScaleInterval(scaleName)
		scaleNotes := scale.GetScaleNotes(scaleName)

		fretdiagram := diagram.NewFretBoard()
		fretdiagram.DrawDiagram()
		fretdiagram.ColorScale(inter)
		fretdiagram.DrawTitle(scaleName, scaleNotes, 40, 100)
		fretdiagram.SaveScaleDiagram("./output/guitar/" + scaleName + ".png")

		bassdiagram := diagram.NewBassDiagram()
		bassdiagram.DrawDiagram()
		bassdiagram.ColorScale(inter)
		bassdiagram.DrawTitle(scaleName, scaleNotes, 40, 100)
		bassdiagram.SaveScaleDiagram("./output/bass/" + scaleName + ".png")

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
	bass := diagram.NewFretBoard()
	if err := bass.TilePNGsToPDF("./output/bass/", "./output/bass_scales.pdf"); err != nil {
		log.Fatal(err)
	}

	piano := diagram.NewPianoDiagram()
	if err := piano.TilePNGsToPDF("./output/piano/", "./output/piano_scales.pdf"); err != nil {
		log.Fatal(err)
	}

}
