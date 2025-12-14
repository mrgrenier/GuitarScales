package diagram

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/llgcode/draw2d/draw2dpdf"
)

// tilePNGsToPDF reads all PNG files in inputDir and writes them to a multi-page
// Letter PDF (8.5x11 in) with 9 tiles (3x3) per page.
func TilePNGsToPDF(inputDir, outPDFPath string) error {
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
		cols   = 3
		rows   = 3
		margin = 36.0 // 0.5"
		gap    = 12.0
	)

	cellW := (pageW - 2*margin - float64(cols-1)*gap) / float64(cols)
	cellH := (pageH - 2*margin - float64(rows-1)*gap) / float64(rows)

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

		x0 := margin + float64(c)*(cellW+gap)
		y0 := margin + float64(r)*(cellH+gap)

		iw := float64(img.Bounds().Dx())
		ih := float64(img.Bounds().Dy())

		// Scale-to-fit (preserve aspect ratio)
		sx := cellW / iw
		sy := cellH / ih
		s := sx
		if sy < sx {
			s = sy
		}
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
