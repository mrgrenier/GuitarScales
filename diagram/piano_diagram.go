package diagram

type PianoDiagram struct {
	scaleOctaveWidth float64
	scaleHeight      float64
	canvasWidth      int
	canvasHeight     int
	noteposX         []float64
	noteposY         []float64
}

func NewPianoDiagram() Diagram {
	return &PianoDiagram{
		scaleOctaveWidth: 6.5,
		scaleHeight:      1.25,
	}
}

var _ Diagram = (*PianoDiagram)(nil)

func (p *PianoDiagram) DrawDiagram() {
	// For piano you might interpret this as "DrawKeyboard()",
	// but keep the shared name for interface compatibility.
}

func (p *PianoDiagram) ColorScale(interval []string) {
	// highlight keys matching interval
}

func (p *PianoDiagram) DrawTitle(scaleName, scaleNotes string, x, y float64) {
	// draw title on the piano diagram
}

func (p *PianoDiagram) SaveScaleDiagram(filename string) {
	// save output image
}
