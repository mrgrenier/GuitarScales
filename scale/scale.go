package scale

import (
	"container/ring"
	"fmt"
	"github.com/mrgrenier/GuitarScales/note"
	"sort"
	"strings"
)

type Scale struct {
	notes    *ring.Ring
	root     *ring.Ring
	interval *Interval
	scales   map[string][]string
}

func NewScale(root note.Note) *Scale {

	n := &Scale{}
	n.interval = NewInterval()
	n.scales = make(map[string][]string)
	n.scales["ionian"] = append(n.scales["ionian"], "1", "2", "3", "4", "5", "6", "7")
	n.scales["dorian"] = append(n.scales["dorian"], "1", "2", "b3", "4", "5", "6", "b7")
	n.scales["phrygian"] = append(n.scales["phrygian"], "1", "b2", "b3", "4", "5", "b6", "b7")
	n.scales["lydian"] = append(n.scales["lydian"], "1", "2", "3", "#4", "5", "6", "7")
	n.scales["mixolydian"] = append(n.scales["mixolydian"], "1", "2", "3", "4", "5", "6", "b7")
	n.scales["aoelian"] = append(n.scales["aoelian"], "1", "2", "b3", "4", "5", "b6", "b7")
	n.scales["locrian"] = append(n.scales["locrian"], "1", "b2", "b3", "4", "b5", "b6", "b7")
	n.scales["harmonic major"] = append(n.scales["harmonic major"], "1", "2", "3", "4", "5", "b6", "7")
	n.scales["harmonic minor"] = append(n.scales["harmonic minor"], "1", "2", "b3", "4", "5", "b6", "7")
	n.scales["bebop dominant"] = append(n.scales["bebop dominant"], "1", "2", "3", "4", "5", "#6", "7")
	n.scales["bebop major"] = append(n.scales["bebop major"], "1", "2", "3", "4", "5", "b6", "6", "7")
	n.scales["minor blues"] = append(n.scales["minor blues"], "1", "b3", "4", "b5", "5", "b7")
	n.scales["voodoo blues"] = append(n.scales["voodoo blues"], "1", "b3", "4", "b5", "5", "6")
	n.scales["major blues"] = append(n.scales["major blues"], "1", "2", "b3", "3", "5", "6")
	n.scales["arabian"] = append(n.scales["arabian"], "1", "2", "b3", "4", "b5", "b6", "6", "7")
	n.scales["balinese"] = append(n.scales["balinese"], "1", "b2", "b3", "5", "b6")
	n.scales["phrygian dominant"] = append(n.scales["phrygian dominant"], "1", "b2", "3", "4", "5", "b6", "b7")
	n.scales["byzantine"] = append(n.scales["byzantine"], "1", "b2", "3", "4", "5", "b6", "7")
	n.scales["chinese"] = append(n.scales["chinese"], "1", "3", "#4", "5", "7")
	n.scales["chromatic"] = append(n.scales["chromatic"], "1", "b2", "2", "b3", "3", "4", "b5", "5", "b6", "6", "b7", "7")
	n.scales["composite"] = append(n.scales["composite"], "1", "b2", "b3", "3", "#4", "5", "b6", "b7")
	n.scales["egyptian"] = append(n.scales["egyptian"], "1", "2", "4", "5", "b7")
	n.scales["enigmatic"] = append(n.scales["enigmatic"], "1", "b2", "3", "#4", "#5", "#6", "7")
	n.scales["hinduston"] = append(n.scales["hinduston"], "1", "2", "3", "4", "5", "b6", "b7")
	n.scales["hungarian major"] = append(n.scales["hungarian major"], "1", "#2", "3", "#4", "5", "6", "b7")
	n.scales["hungarian gypsy "] = append(n.scales["hungarian gypsy "], "1", "2", "b3", "#4", "5", "b6", "7")
	n.scales["japanese"] = append(n.scales["japanese"], "1", "b2", "4", "5", "b6")
	n.scales["japanese hirajoshi"] = append(n.scales["japanese hirajoshi"], "1", "2", "b3", "5", "b6")
	n.scales["japanese kumoi"] = append(n.scales["japanese kumoi"], "1", "2", "b3", "5", "6")
	n.scales["japanese kokin joshi"] = append(n.scales["japanese kokin joshi"], "1", "b2", "4", "5", "b7")
	n.scales["japanese iwato"] = append(n.scales["japanese iwato"], "1", "b2", "4", "b5", "b7")
	n.scales["neapolitan major"] = append(n.scales["neapolitan major"], "1", "b2", "b3", "4", "5", "6", "7")
	n.scales["neapolitan minor"] = append(n.scales["neapolitan minor"], "1", "b2", "b3", "4", "5", "b6", "7")
	n.scales["oriental"] = append(n.scales["oriental"], "1", "b2", "3", "4", "b5", "6", "b7")
	n.scales["overtone"] = append(n.scales["overtone"], "1", "2", "3", "#4", "5", "6", "b7")
	n.scales["pelog"] = append(n.scales["pelog"], "1", "b2", "b3", "5", "b6")
	n.scales["persian"] = append(n.scales["persian"], "1", "b2", "3", "4", "b5", "b6", "7")
	n.scales["romanian"] = append(n.scales["romanian"], "1", "2", "b3", "#4", "5", "6", "b7")
	n.scales["scribian"] = append(n.scales["scribian"], "1", "b2", "3", "5", "6")
	n.scales["symmetrical"] = append(n.scales["symmetrical"], "1", "b2", "b3", "3", "#4", "5", "6", "b7")

	allnotes := []note.Note{
		{Name: "A", Alternate: root.Alternate},
		{Name: "A#", Alternate: root.Alternate},
		{Name: "B", Alternate: root.Alternate},
		{Name: "C", Alternate: root.Alternate},
		{Name: "C#", Alternate: root.Alternate},
		{Name: "D", Alternate: root.Alternate},
		{Name: "D#", Alternate: root.Alternate},
		{Name: "E", Alternate: root.Alternate},
		{Name: "F", Alternate: root.Alternate},
		{Name: "F#", Alternate: root.Alternate},
		{Name: "G", Alternate: root.Alternate},
		{Name: "G#", Alternate: root.Alternate},
	}
	n.notes = ring.New(len(allnotes))

	// add all the notes to the circular ring
	for _, no := range allnotes {
		n.notes.Value = no
		n.notes = n.notes.Next()
	}

	n.SetRoot(root)

	return n
}

func (n *Scale) SetRoot(root note.Note) {
	for i := 0; i < n.notes.Len(); i++ {
		if n.notes.Value == root {
			n.root = n.notes
			break
		}
		n.notes = n.notes.Next()
	}
}

func (n *Scale) ScaleNames() []string {
	var scaleNames []string
	for name := range n.scales {
		scaleNames = append(scaleNames, name)
	}
	sort.Strings(scaleNames)
	return scaleNames
}

func (n *Scale) ScaleNotes(name string) []note.Note {
	var notes []note.Note

	for _, scaleNote := range n.scales[name] {
		notes = append(notes, n.ShowNoteAt(scaleNote))
	}
	return notes
}

func (n *Scale) ScaleInterval(name string) []string {
	var inter []string

	for _, scaleNote := range n.scales[name] {
		inter = append(inter, scaleNote)
	}
	return inter
}

func (n *Scale) GetScaleNotes(scaleName string) string {
	var sb strings.Builder
	for _, scaleNote := range n.scales[scaleName] {
		sb.WriteString(n.ShowNoteAt(scaleNote).String())
	}
	return sb.String()
}

func (n *Scale) ShowAll() {
	fmt.Printf("Root: %s\n", n.root.Value)

	n.notes = n.root

	for _, scale := range n.ScaleNames() {
		fmt.Print(scale + ": ")
		for _, scaleNote := range n.scales[scale] {
			fmt.Print(n.ShowNoteAt(scaleNote))
		}
		fmt.Println()
	}
	fmt.Print()

}

func (n *Scale) ShowNoteAt(interval string) note.Note {

	offset, err := n.interval.IntervalToOffset(interval)
	if err != nil {
		fmt.Println(err.Error())
		return note.Note{}
	}

	n.notes = n.root
	n.notes = n.notes.Move(offset)

	no := n.notes.Value.(note.Note)
	return no
}
