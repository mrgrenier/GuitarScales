package chord

import (
	"container/ring"
	"fmt"

	"github.com/mrgrenier/GuitarScales/note"
	"github.com/mrgrenier/GuitarScales/scale"
)

type Chord struct {
	notes            *ring.Ring
	root             *ring.Ring
	interval         *scale.Interval
	chords2intervals map[string][]string
	intervals2chords map[int]string
}

func NewChord(root note.Note) *Chord {

	n := &Chord{}
	n.interval = scale.NewInterval()
	n.chords2intervals = make(map[string][]string)
	n.chords2intervals["Major"] = append(n.chords2intervals["Major"], "1", "3", "5")
	n.chords2intervals["Major6th"] = append(n.chords2intervals["Major6th"], "1", "3", "5", "6")
	n.chords2intervals["Major7th"] = append(n.chords2intervals["Major7th"], "1", "3", "5", "7")
	n.chords2intervals["Major9th"] = append(n.chords2intervals["Major9th"], "1", "3", "5", "7", "2")
	n.chords2intervals["Major13th"] = append(n.chords2intervals["Major13th"], "1", "3", "5", "7", "2", "6")
	n.chords2intervals["Minor"] = append(n.chords2intervals["Minor"], "1", "b3", "5")
	n.chords2intervals["Minor6th"] = append(n.chords2intervals["Minor6th"], "1", "b3", "5", "6")
	n.chords2intervals["Minor7th"] = append(n.chords2intervals["Minor7th"], "1", "b3", "5", "b7")
	n.chords2intervals["Minor9th"] = append(n.chords2intervals["Minor9th"], "1", "b3", "5", "b7", "2")
	n.chords2intervals["Minor11th"] = append(n.chords2intervals["Minor11th"], "1", "b3", "5", "b7", "2", "4")
	n.chords2intervals["Minor13th"] = append(n.chords2intervals["Minor13th"], "1", "b3", "5", "b7", "2", "6")
	n.chords2intervals["Dim"] = append(n.chords2intervals["Dim"], "1", "b3", "b5")
	n.chords2intervals["Dim7th"] = append(n.chords2intervals["Dim7th"], "1", "b3", "b5", "6")
	n.chords2intervals["Dim7b5"] = append(n.chords2intervals["Dim7b5"], "1", "b3", "b5", "b7")
	n.chords2intervals["Aug"] = append(n.chords2intervals["Aug"], "1", "3", "#5")
	n.chords2intervals["Aug7th"] = append(n.chords2intervals["Aug7th"], "1", "3", "#5", "b7")
	n.chords2intervals["Dom7th"] = append(n.chords2intervals["Dom7th"], "1", "3", "5", "b7")
	n.chords2intervals["Dom9th"] = append(n.chords2intervals["Dom9th"], "1", "3", "5", "b7", "2")
	n.chords2intervals["Dom11th"] = append(n.chords2intervals["Dom11th"], "1", "5", "b7", "2", "4")
	n.chords2intervals["Sus2"] = append(n.chords2intervals["Sus2"], "1", "2", "5")
	n.chords2intervals["Sus4"] = append(n.chords2intervals["Sus4"], "1", "4", "5")
	n.chords2intervals["Add9"] = append(n.chords2intervals["Add9"], "1", "3", "5", "2")

	n.intervals2chords = make(map[int]string)
	for chord, v := range n.chords2intervals {
		m := 0
		for _, intr := range v {
			i, _ := n.interval.IntervalToOffset(intr)
			m = m | 1<<(11-i)
		}
		n.intervals2chords[m] = chord
	}
	for key, value := range n.interval.GetOffset() {
		fmt.Printf("")
		fmt.Printf("Key: %s, 0x%x\n", key, 1<<(11-value))
	}
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

func (n *Chord) SetRoot(root note.Note) {
	for i := 0; i < n.notes.Len(); i++ {
		if n.notes.Value == root {
			n.root = n.notes
			break
		}
		n.notes = n.notes.Next()
	}
}

func (n *Chord) GetChordNames(root note.Note, intervals []string) ([]string, error) {
	// Save the current root to restore it later
	originalRoot := n.root

	// Temporarily set the new root
	n.SetRoot(root)

	// Ensure we restore the original root when done
	defer func() {
		n.root = originalRoot
	}()

	// Calculate the bitmask for the given intervals
	inputMask := 0
	for _, interval := range intervals {
		offset, err := n.interval.IntervalToOffset(interval)
		if err != nil {
			return nil, fmt.Errorf("invalid interval %s: %v", interval, err)
		}
		inputMask = inputMask | 1<<(11-offset)
	}

	// Find all chords that are subsets of the input intervals
	var matchingChords []string
	for chordMask, chordName := range n.intervals2chords {
		// Check if the chord is a subset of the input intervals
		// A chord is a subset if all its bits are present in the input mask
		if (chordMask & inputMask) == chordMask {
			matchingChords = append(matchingChords, chordName)
		}
	}

	if len(matchingChords) == 0 {
		return nil, fmt.Errorf("no chords found for intervals %v", intervals)
	}

	return matchingChords, nil
}
