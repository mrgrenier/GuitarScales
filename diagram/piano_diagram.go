package diagram

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dpdf"
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
	offsetY = 150.0
)

func NewPianoDiagram() Diagram {
	scaleOctaveWidth := 6.5
	p := &PianoDiagram{
		scaleOctaveWidth: scaleOctaveWidth,
		scaleHeight:      1.25,
		canvasWidth:      1188,
		canvasHeight:     940,
		keyWidth:         scaleOctaveWidth / 12,
	}
	p.dest = image.NewRGBA(image.Rect(0, 0, p.canvasWidth, p.canvasHeight))
	p.gc = draw2dimg.NewGraphicContext(p.dest)

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

	p.gc.SetFontData(fontdata)

	p.StringFret2Interval = make(map[int]map[string]bool)
	for f := 0; f < 12; f++ {
		p.StringFret2Interval[f] = make(map[string]bool)
	}

	p.StringFret2Interval[0]["1"] = true
	p.StringFret2Interval[1]["b2"] = true
	p.StringFret2Interval[2]["2"] = true
	p.StringFret2Interval[3]["#2"] = true
	p.StringFret2Interval[3]["b3"] = true
	p.StringFret2Interval[4]["3"] = true
	p.StringFret2Interval[5]["4"] = true
	p.StringFret2Interval[6]["#4"] = true
	p.StringFret2Interval[6]["b5"] = true
	p.StringFret2Interval[7]["5"] = true
	p.StringFret2Interval[8]["#5"] = true
	p.StringFret2Interval[8]["b6"] = true
	p.StringFret2Interval[9]["6"] = true
	p.StringFret2Interval[10]["#6"] = true
	p.StringFret2Interval[10]["b7"] = true
	p.StringFret2Interval[11]["7"] = true

	p.flatName = make(map[string]string)
	p.flatName["#2"] = "b3"
	p.flatName["#4"] = "b5"
	p.flatName["#5"] = "b6"
	p.flatName["#6"] = "b7"

	return p
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
		}
		x := x0 + float64(i)*segW
		x1 = x + segW
		p.gc.BeginPath() // Initialize a new path
		p.gc.SetFillColor(noteColor)
		p.gc.SetStrokeColor(fontColor)
		p.gc.MoveTo(x, y0)
		p.gc.LineTo(x1, y0)
		p.gc.LineTo(x1, y1)
		p.gc.LineTo(x, y1)
		p.gc.Close()
		p.gc.FillStroke()
		p.DrawInterval(note, x+segW/1.5, y0+h/2.1, segW, fontColor)
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
		p.gc.FillStringAt(flat, x-(radius+accidentalsFontSize)/3, y+accidentalsFontSize)
		x = x + accidentalsFontSize/2
		note = note[1:]
	} else if strings.HasPrefix(note, "#") {
		p.gc.SetFontSize(accidentalsFontSize)
		p.gc.FillStringAt(sharp, x-(radius+accidentalsFontSize)/3, y+accidentalsFontSize)
		x = x + accidentalsFontSize/2
		note = note[1:]
	}

	p.gc.SetFontSize(noteFontSize)
	p.gc.FillStringAt(note, x-(radius+noteFontSize)/4.75, y+(radius+noteFontSize)/4.5)

}

func (p *PianoDiagram) DrawTitle(scaleName, scaleNotes string, x, y float64) {
	// draw title on the piano diagram

	textColor := color.RGBA{0x00, 0x00, 0x00, 0xff}

	var fontSize float64 = 44
	var notesFontSize = fontSize * .75

	scaleName = titleCaseASCIIWords(strings.ToLower(scaleName))
	p.gc.SetFillColor(textColor)
	p.gc.SetStrokeColor(textColor)

	p.gc.SetFontSize(fontSize)
	p.gc.FillStringAt(scaleName, x, y)

	p.gc.SetFontSize(notesFontSize)
	p.gc.FillStringAt(scaleNotes, x, y+fontSize+fontSize/2)

}

func (p *PianoDiagram) SaveScaleDiagram(filename string) {
	err := draw2dimg.SaveToPngFile(filename, p.dest)
	if err != nil {
		return
	}
}

func (p *PianoDiagram) getFlatName(interval string) string {
	if flatname, ok := p.flatName[interval]; ok {
		return (flatname)
	}
	return interval
}

// tilePNGsToPDF reads all PNG files in inputDir and writes them to a multi-page
// Letter PDF (8.5x11 in) with 9 tiles (3x3) per page.
func (p *PianoDiagram) TilePNGsToPDF(inputDir, outPDFPath string) error {
	entries, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("read dir %q: %w", inputDir, err)
	}

	var pngPaths []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if strings.HasSuffix(strings.ToLower(name), ".png") {
			pngPaths = append(pngPaths, filepath.Join(inputDir, name))
		}
	}
	sort.Strings(pngPaths)
	if len(pngPaths) == 0 {
		return fmt.Errorf("no PNG files found in %q", inputDir)
	}

	// Letter size in points (1 inch = 72 points)
	const pageW, pageH = 612.0, 792.0

	// Layout settings (points)
	const (
		cols        = 1
		rows        = 3
		topMargin   = 72.0 // 1 inch
		sideMargin  = 36.0 // 0.5"
		extraOffset = 36.0
		gap         = 12.0
	)

	targetWidthPt := p.scaleOctaveWidth * 72.0
	cellW := (pageW - 2*sideMargin - float64(cols-1)*gap) / float64(cols)
	cellH := (pageH - topMargin - sideMargin - float64(rows-1)*gap) / float64(rows)

	pdf := draw2dpdf.NewPdf("P", "pt", "Letter")
	gc := draw2dpdf.NewGraphicContext(pdf)

	for i, p := range pngPaths {
		if i%(cols*rows) == 0 {
			pdf.AddPage()
		}

		f, err := os.Open(p)
		if err != nil {
			return fmt.Errorf("open %q: %w", p, err)
		}
		img, err := png.Decode(f)
		_ = f.Close()
		if err != nil {
			return fmt.Errorf("decode png %q: %w", p, err)
		}

		idxOnPage := i % (cols * rows)
		r := idxOnPage / cols
		c := idxOnPage % cols

		x0 := sideMargin + float64(c)*(cellW+gap)
		y0 := topMargin + extraOffset + float64(r)*(cellH+gap)

		iw := float64(img.Bounds().Dx())
		ih := float64(img.Bounds().Dy())

		// Scale so the image width equals scaleOctaveWidth inches
		s := targetWidthPt / iw
		drawW := iw * s
		drawH := ih * s

		// Center in the cell
		x := x0 + (cellW-drawW)/2
		y := y0 + (cellH-drawH)/2

		gc.Save()
		gc.Translate(x, y)
		gc.Scale(s, s)
		gc.DrawImage(img)
		gc.Restore()
	}

	if err := os.MkdirAll(filepath.Dir(outPDFPath), 0o755); err != nil {
		return fmt.Errorf("mkdir %q: %w", filepath.Dir(outPDFPath), err)
	}
	if err := draw2dpdf.SaveToPdfFile(outPDFPath, pdf); err != nil {
		return fmt.Errorf("save pdf %q: %w", outPDFPath, err)
	}
	return nil
}
