package diagram

// Diagram is the common API that all instrument diagrams must implement
// (guitar, piano, etc).
type Diagram interface {
	DrawDiagram()
	ColorScale(interval []string)
	DrawTitle(scaleName, scaleNotes string, x, y float64)
	SaveScaleDiagram(filename string)
}
