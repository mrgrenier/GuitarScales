package note

var altName = map[string]string{
	"A":  "A",
	"A#": "Bb",
	"B":  "B",
	"C":  "C",
	"C#": "Db",
	"D":  "D",
	"D#": "Eb",
	"E":  "E",
	"F":  "F",
	"F#": "Gb",
	"G":  "G",
	"G#": "Ab",
}

type ALTERNATE_NAME int

const (
	SHARP ALTERNATE_NAME = iota
	FLAT
)

type Note struct {
	Name      string
	Alternate ALTERNATE_NAME
}

func (note *Note) SetAlternate(alternate ALTERNATE_NAME) {
	note.Alternate = alternate
}

func (note Note) String() string {
	var name string
	switch note.Alternate {
	case SHARP:
		name = note.Name
	case FLAT:
		name = altName[note.Name]
	default:
		name = note.Name
	}
	return name + " "
}
