package diagram

import (
	"image"
	"image/color"
	"log"
	"os"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

type PianoDiagram struct {
	scaleOctaveWidth    float64
	scaleHeight         float64
	canvasWidth         int
	canvasHeight        int
	keyWidth            float64
	flatName            map[string]string
	StringFret2Interval map[int]map[string]bool
	dest                *image.RGBA
	gc                  *draw2dimg.GraphicContext
}

const (
	offsetX = 40.0
	offsetY = 140.0
)

func NewPianoDiagram() Diagram {
	scaleOctaveWidth := 6.5
	kb := &PianoDiagram{
		scaleOctaveWidth: scaleOctaveWidth,
		scaleHeight:      1.25,
		canvasWidth:      1188,
		canvasHeight:     940,
		keyWidth:         scaleOctaveWidth / 12,
	}
	kb.dest = image.NewRGBA(image.Rect(0, 0, kb.canvasWidth, kb.canvasHeight))
	kb.gc = draw2dimg.NewGraphicContext(kb.dest)

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

	kb.gc.SetFontData(fontdata)

	kb.StringFret2Interval = make(map[int]map[string]bool)
	for f := 0; f < 12; f++ {
		kb.StringFret2Interval[f] = make(map[string]bool)
	}

	kb.StringFret2Interval[0]["1"] = true
	kb.StringFret2Interval[1]["b2"] = true
	kb.StringFret2Interval[2]["2"] = true
	kb.StringFret2Interval[3]["#2"] = true
	kb.StringFret2Interval[3]["b3"] = true
	kb.StringFret2Interval[4]["3"] = true
	kb.StringFret2Interval[5]["4"] = true
	kb.StringFret2Interval[6]["#4"] = true
	kb.StringFret2Interval[6]["b5"] = true
	kb.StringFret2Interval[7]["5"] = true
	kb.StringFret2Interval[8]["#5"] = true
	kb.StringFret2Interval[8]["b6"] = true
	kb.StringFret2Interval[9]["6"] = true
	kb.StringFret2Interval[10]["#6"] = true
	kb.StringFret2Interval[10]["b7"] = true
	kb.StringFret2Interval[11]["7"] = true

	kb.flatName = make(map[string]string)
	kb.flatName["#2"] = "b3"
	kb.flatName["#4"] = "b5"
	kb.flatName["#5"] = "b6"
	kb.flatName["#6"] = "b7"

	return kb
}

var _ Diagram = (*PianoDiagram)(nil)

func (p *PianoDiagram) DrawDiagram() {

	// Draw one octave: an outer rectangle of (scaleOctaveWidth x scaleHeight),
	// divided into 12 equal keyWidth segments.

	usableW := float64(p.canvasWidth) - 2*offsetX
	unitToPx := usableW / p.scaleOctaveWidth

	w := p.scaleOctaveWidth * unitToPx
	h := p.scaleHeight * unitToPx

	x0 := offsetX
	y0 := offsetY
	x1 := x0 + w
	y1 := y0 + h

	p.gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	p.gc.SetLineWidth(2)

	// Outer rectangle
	p.gc.BeginPath()
	p.gc.MoveTo(x0, y0)
	p.gc.LineTo(x1, y0)
	p.gc.LineTo(x1, y1)
	p.gc.LineTo(x0, y1)
	p.gc.Close()
	p.gc.Stroke()

	// 12 segments (vertical dividers)
	segW := p.keyWidth * unitToPx
	for i := 1; i < 12; i++ {
		x := x0 + float64(i)*segW
		p.gc.BeginPath()
		p.gc.MoveTo(x, y0)
		p.gc.LineTo(x, y1)
		p.gc.Stroke()
	}

}

func (p *PianoDiagram) ColorScale(interval []string) {

	intervalmap := make(map[string]bool)
	for _, i := range interval {
		intervalmap[i] = true
	}

	blankNoteColor := color.RGBA{0xee, 0xee, 0xee, 0xff}
	blankNoteFontColor := color.RGBA{0x00, 0x00, 0x00, 0xff}
	rootNoteColor := color.RGBA{0xff, 0x44, 0x44, 0xff}
	rootNoteFontColor := color.RGBA{0xff, 0xff, 0xff, 0xff}
	scaleNoteColor := color.RGBA{0x00, 0x00, 0x00, 0xff}
	scaleNoteFontColor := color.RGBA{0xff, 0xff, 0xff, 0xff}

	noteColor := blankNoteColor
	fontColor := blankNoteFontColor
	var note string

	usableW := float64(p.canvasWidth) - 2*offsetX
	unitToPx := usableW / p.scaleOctaveWidth
	w := p.scaleOctaveWidth * unitToPx
	h := p.scaleHeight * unitToPx

	segW := p.keyWidth * unitToPx
	x0 := offsetX
	y0 := offsetY
	x1 := x0 + w
	y1 := y0 + h

	for i := 0; i < 12; i++ {
		noteColor = blankNoteColor
		fontColor = blankNoteFontColor
		for note = range p.StringFret2Interval[i] {
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
				note = p.getFlatName(note)
			}
			x := x0 + float64(i)*segW
			x1 = x + w

			p.gc.BeginPath() // Initialize a new path
			p.gc.SetFillColor(noteColor)
			p.gc.SetStrokeColor(noteColor)
			p.gc.MoveTo(x, y0)
			p.gc.LineTo(x1, y0)
			p.gc.LineTo(x1, y1)
			p.gc.LineTo(x, y1)
			p.gc.Close()
			p.DrawInterval(note, x, y0, w/2, fontColor)

		}

	}
}

func (p *PianoDiagram) DrawInterval(note string, x, y, radius float64, textColor color.RGBA) {

	flat := string([]rune{'\u266D'})
	sharp := string([]rune{'\u266F'})
	var noteFontSize float64 = 34
	var accidentalsFontSize float64 = 20

	p.gc.SetFillColor(textColor)
	p.gc.SetStrokeColor(textColor)

	if strings.HasPrefix(note, "b") {
		p.gc.SetFontSize(accidentalsFontSize)
		p.gc.FillStringAt(flat, x-(radius+accidentalsFontSize)/2.25, y+accidentalsFontSize/2)
		x = x + accidentalsFontSize/2
		note = note[1:]
	} else if strings.HasPrefix(note, "#") {
		p.gc.SetFontSize(accidentalsFontSize)
		p.gc.FillStringAt(sharp, x-(radius+accidentalsFontSize)/2.25, y+accidentalsFontSize/2)
		x = x + accidentalsFontSize/2
		note = note[1:]
	}

	p.gc.SetFontSize(noteFontSize)
	p.gc.FillStringAt(note, x-(radius+noteFontSize)/4.75, y+(radius+noteFontSize)/4.5)

}

func (p *PianoDiagram) DrawTitle(scaleName, scaleNotes string, x, y float64) {
	// draw title on the piano diagram
}

func (p *PianoDiagram) SaveScaleDiagram(filename string) {
	// save output image
}

func (p *PianoDiagram) getFlatName(interval string) string {
	if flatname, ok := p.flatName[interval]; ok {
		return (flatname)
	}
	return interval
}
