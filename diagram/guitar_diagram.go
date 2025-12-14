package diagram

import (
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type FretBoard struct {
	scaleLength         float64
	numFrets            int
	canvasWidth         int
	canvasHeight        int
	strings             int
	noteposX            []float64
	noteposY            []float64
	StringFret2Interval map[int]map[int]map[string]bool
	flatName            map[string]string
	dest                *image.RGBA
	gc                  *draw2dimg.GraphicContext
}

// NewGuitarDiagram returns the common Diagram interface backed by the guitar fretboard implementation.
func NewGuitarDiagram() Diagram {
	return NewFretBoard()
}

// Compile-time check that *FretBoard implements Diagram.
var _ Diagram = (*FretBoard)(nil)

func NewFretBoard() *FretBoard {
	fb := &FretBoard{scaleLength: 25.5, numFrets: 6,
		canvasWidth: 1188, canvasHeight: 940, strings: 6}

	fb.dest = image.NewRGBA(image.Rect(0, 0, fb.canvasWidth, fb.canvasHeight))
	fb.gc = draw2dimg.NewGraphicContext(fb.dest)

	draw2d.SetFontFolder("./resource/font")
	b, err := os.ReadFile("./resource/font/wqy-zenhei.ttf")
	if err != nil {
		log.Fatal(err)
	}
	font, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}
	fontdata := draw2d.FontData{Name: "wgy-zenhei", Family: draw2d.FontFamilyMono, Style: draw2d.FontStyleNormal}
	draw2d.RegisterFont(
		fontdata,
		font,
	)

	fb.gc.SetFontData(fontdata)

	fb.StringFret2Interval = make(map[int]map[int]map[string]bool)
	for f := 0; f < fb.numFrets; f++ {
		fb.StringFret2Interval[f] = make(map[int]map[string]bool)
		for s := 0; s < 6; s++ {
			fb.StringFret2Interval[f][s] = make(map[string]bool)
		}
	}

	fb.StringFret2Interval[0][0]["7"] = true
	fb.StringFret2Interval[1][0]["1"] = true
	fb.StringFret2Interval[2][0]["b2"] = true
	fb.StringFret2Interval[3][0]["2"] = true
	fb.StringFret2Interval[4][0]["#2"] = true
	fb.StringFret2Interval[4][0]["b3"] = true
	fb.StringFret2Interval[5][0]["3"] = true

	fb.StringFret2Interval[0][1]["3"] = true
	fb.StringFret2Interval[1][1]["4"] = true
	fb.StringFret2Interval[2][1]["#4"] = true
	fb.StringFret2Interval[2][1]["b5"] = true
	fb.StringFret2Interval[3][1]["5"] = true
	fb.StringFret2Interval[4][1]["#5"] = true
	fb.StringFret2Interval[4][1]["b6"] = true
	fb.StringFret2Interval[5][1]["6"] = true

	fb.StringFret2Interval[0][2]["6"] = true
	fb.StringFret2Interval[1][2]["#6"] = true
	fb.StringFret2Interval[1][2]["b7"] = true
	fb.StringFret2Interval[2][2]["7"] = true
	fb.StringFret2Interval[3][2]["1"] = true
	fb.StringFret2Interval[4][2]["b2"] = true
	fb.StringFret2Interval[5][2]["2"] = true

	fb.StringFret2Interval[0][3]["2"] = true
	fb.StringFret2Interval[1][3]["#2"] = true
	fb.StringFret2Interval[1][3]["b3"] = true
	fb.StringFret2Interval[2][3]["3"] = true
	fb.StringFret2Interval[3][3]["4"] = true
	fb.StringFret2Interval[4][3]["#4"] = true
	fb.StringFret2Interval[4][3]["b5"] = true
	fb.StringFret2Interval[5][3]["5"] = true

	fb.StringFret2Interval[0][4]["#4"] = true
	fb.StringFret2Interval[0][4]["b5"] = true
	fb.StringFret2Interval[1][4]["5"] = true
	fb.StringFret2Interval[2][4]["#5"] = true
	fb.StringFret2Interval[2][4]["b6"] = true
	fb.StringFret2Interval[3][4]["6"] = true
	fb.StringFret2Interval[4][4]["#6"] = true
	fb.StringFret2Interval[4][4]["b7"] = true
	fb.StringFret2Interval[5][4]["7"] = true

	fb.StringFret2Interval[0][5]["7"] = true
	fb.StringFret2Interval[1][5]["1"] = true
	fb.StringFret2Interval[2][5]["b2"] = true
	fb.StringFret2Interval[3][5]["2"] = true
	fb.StringFret2Interval[4][5]["#2"] = true
	fb.StringFret2Interval[4][5]["b3"] = true
	fb.StringFret2Interval[5][5]["3"] = true

	fb.flatName = make(map[string]string)
	fb.flatName["#2"] = "b3"
	fb.flatName["#4"] = "b5"
	fb.flatName["#5"] = "b6"
	fb.flatName["#6"] = "b7"

	return fb
}

func (fb *FretBoard) DrawDiagram() {
	var offsetX float64 = 40
	var offsetY float64 = (float64(fb.canvasHeight) / 2) + float64(fb.canvasHeight)/10
	var scale float64 = 150
	var scale_factor float64 = 0.0
	var distanceFromNut float64 = 0.0
	var bridgeToFret float64 = 0.0
	var nut = []float64{(distanceFromNut * scale) + offsetX, (2.2 * scale) + offsetY, (distanceFromNut * scale) + offsetX, (-2.2 * scale) + offsetY}
	var va = []float64{(distanceFromNut * scale) + offsetX, (2.2 * scale) + offsetY, (distanceFromNut * scale) + offsetX, (-2.2 * scale) + offsetY}

	// Initialize the graphic context on an RGBA image
	fb.gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	fb.gc.SetLineWidth(2)

	// Draw the nut
	fb.gc.BeginPath() // Initialize a new path
	fb.gc.MoveTo(nut[0], nut[1])
	fb.gc.LineTo(nut[2], nut[3])
	fb.gc.FillStroke()

	// Draw the frets
	for fret := 1; fret <= fb.numFrets; fret++ {
		bridgeToFret = fb.scaleLength - distanceFromNut
		scale_factor = bridgeToFret / 17.817
		prevDistanceFromNut := distanceFromNut
		distanceFromNut = distanceFromNut + scale_factor

		fretPos := (distanceFromNut * scale) + offsetX
		prevFret := (prevDistanceFromNut * scale) + offsetX

		fb.noteposX = append(fb.noteposX, ((prevFret + fretPos) / 2))
		va = []float64{fretPos, (2.2 * scale) + offsetY, fretPos, (-2.2 * scale) + offsetY}
		fb.gc.BeginPath() // Initialize a new path
		fb.gc.MoveTo(va[0], va[1])
		fb.gc.LineTo(va[2], va[3])
		fb.gc.FillStroke()
	}

	segments := (nut[1] - nut[3]) / (float64(fb.strings) - 1)

	// draw the strings
	for i := 0; i < fb.strings; i++ {
		fb.gc.BeginPath() // Initialize a new path
		fb.gc.MoveTo(nut[0], nut[1]-(segments*float64(i)))
		fb.gc.LineTo(va[0], va[1]-(segments*float64(i)))
		fb.noteposY = append(fb.noteposY, nut[1]-(segments*float64(i)))
		fb.gc.FillStroke()
	}

}

func (fb *FretBoard) ColorScale(interval []string) {

	intervalmap := make(map[string]bool)
	for _, i := range interval {
		intervalmap[i] = true
	}

	// Draw the note circles
	var radius float64 = 40

	blankNoteColor := color.RGBA{0xee, 0xee, 0xee, 0xff}
	blankNoteFontColor := color.RGBA{0x00, 0x00, 0x00, 0xff}
	rootNoteColor := color.RGBA{0xff, 0x44, 0x44, 0xff}
	rootNoteFontColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	scaleNoteColor := color.RGBA{0x00, 0x00, 0x00, 0xff}
	scaleNoteFontColor := color.RGBA{0xff, 0xff, 0xff, 0xff}

	noteColor := blankNoteColor
	fontColor := blankNoteFontColor
	var note string

	for f, x := range fb.noteposX {
		for s, y := range fb.noteposY {
			noteColor = blankNoteColor
			fontColor = blankNoteFontColor
			for note = range fb.StringFret2Interval[f][s] {
				if note == "1" {
					noteColor = rootNoteColor
					fontColor = rootNoteFontColor
					note = "R"
					break
				} else if intervalmap[note] == true {
					noteColor = scaleNoteColor
					fontColor = scaleNoteFontColor
					break
				} else {
					// for the notes not in the interval favor the flat name ins stead of the sharp name
					note = fb.getFlatName(note)
				}

			}

			fb.gc.BeginPath() // Initialize a new path
			fb.gc.SetFillColor(noteColor)
			fb.gc.SetStrokeColor(noteColor)
			fb.gc.MoveTo(x+radius, y)
			fb.gc.ArcTo(x, y, radius, radius, 0, -math.Pi*2)
			fb.gc.FillStroke()
			fb.DrawInterval(note, x, y, radius, fontColor)
		}
	}

}

func (fb *FretBoard) DrawInterval(note string, x, y, radius float64, textColor color.RGBA) {

	flat := string([]rune{'\u266D'})
	sharp := string([]rune{'\u266F'})
	var noteFontSize float64 = 34
	var accidentalsFontSize float64 = 20

	fb.gc.SetFillColor(textColor)
	fb.gc.SetStrokeColor(textColor)

	if strings.HasPrefix(note, "b") {
		fb.gc.SetFontSize(accidentalsFontSize)
		fb.gc.FillStringAt(flat, x-(radius+accidentalsFontSize)/2.25, y+accidentalsFontSize/2)
		x = x + accidentalsFontSize/2
		note = note[1:]
	} else if strings.HasPrefix(note, "#") {
		fb.gc.SetFontSize(accidentalsFontSize)
		fb.gc.FillStringAt(sharp, x-(radius+accidentalsFontSize)/2.25, y+accidentalsFontSize/2)
		x = x + accidentalsFontSize/2
		note = note[1:]
	}

	fb.gc.SetFontSize(noteFontSize)
	fb.gc.FillStringAt(note, x-(radius+noteFontSize)/4.75, y+(radius+noteFontSize)/4.5)

}

func (fb *FretBoard) DrawTitle(scaleName, scaleNotes string, x, y float64) {

	textColor := color.RGBA{0x00, 0x00, 0x00, 0xff}

	var fontSize float64 = 44
	var notesFontSize = fontSize * .75

	scaleName = strings.Title(strings.ToLower(scaleName))
	fb.gc.SetFillColor(textColor)
	fb.gc.SetStrokeColor(textColor)

	fb.gc.SetFontSize(fontSize)
	fb.gc.FillStringAt(scaleName, x, y)

	fb.gc.SetFontSize(notesFontSize)
	fb.gc.FillStringAt(scaleNotes, x, y+fontSize+fontSize/2)

}

func (fb *FretBoard) SaveScaleDiagram(filename string) {
	err := draw2dimg.SaveToPngFile(filename, fb.dest)
	if err != nil {
		return
	}
}

func (fb *FretBoard) getFlatName(interval string) string {
	if flatname, ok := fb.flatName[interval]; ok {
		return (flatname)
	}
	return interval
}
