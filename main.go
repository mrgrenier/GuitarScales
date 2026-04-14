package main

import (
	"fmt"
	"log"

	"github.com/mrgrenier/GuitarScales/chord"
	"github.com/mrgrenier/GuitarScales/diagram"
	"github.com/mrgrenier/GuitarScales/note"
	"github.com/mrgrenier/GuitarScales/scale"
)

func main() {

	root := note.Note{Name: "A", Alternate: note.SHARP}

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

	// Show chords in the "ionian" scale

	scaleName := "ionian"
	interval := scale.IntervalOffset()

	s := scale.ScaleInterval(scaleName)
	inverse_lk := interval.GetInverse()
	offset_lk := interval.GetOffset()

	// f(x) = (x -2 +12) mod 12
	var modes map[string][]string
	modes = make(map[string][]string)
	//build chords
	for idx, j := range s {
		mode_name := fmt.Sprintf("%d", idx+1)
		for _, k := range s {
			fmt.Printf("%s %d-%d %d %v\n", j, offset_lk[k], offset_lk[j], (offset_lk[k]-offset_lk[j]+12)%12, inverse_lk[(offset_lk[k]-offset_lk[j]+12)%12])
			modes[mode_name] = append(modes[mode_name], inverse_lk[(offset_lk[k]-offset_lk[j]+12)%12]...)
		}
	}
	chords := chord.NewChord(root)
	for idx, i := range s {
		n := scale.ShowNoteAt(i)
		mode_name := fmt.Sprintf("%d", idx+1)

		t, _ := chords.GetChordNames(n, modes[mode_name])
		for _, chordName := range t {
			fmt.Printf("%s %s\n", n.Name, chordName)
		}
	}

}

func indexOfString(slice []string, target string) int {
	for i, s := range slice {
		if s == target {
			return i
		}
	}
	return -1
}

func rotateLeftStrings(s []string, k int) []string {
	if len(s) == 0 {
		return s
	}
	k %= len(s)
	return append(s[k:], s[:k]...)
}

func rotateRightStrings(s []string, k int) []string {
	if len(s) == 0 {
		return s
	}
	k %= len(s)
	return append(s[len(s)-k:], s[:len(s)-k]...)
}
