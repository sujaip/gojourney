package poetry

import "fmt"

type Line string

/* Stanza is Slice of Strings */
type Stanza []Line

/* Poem is Slice of Slice of strings */
type Poem []Stanza

func NewPoem() Poem {
	return Poem{}
}

func (p Poem) NumStanzas() int {
	return len(p)
}

func (s Stanza) NumLines() int {
	return len(s)
}

func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

func (p Poem) Stats() (numVowels, numConsonants, numPuncs int) {
	/* for each Slice of string in the Slice of Slice of strings */
	for _, s := range p {
		/* for each string in the Slice of String */
		for _, l := range s {
			/* for each rune in the String */
			for _, r := range l {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels++
				case ',', ' ', '!':
					numPuncs++
				default:
					numConsonants++
				}
			}
		}
	}
	return
}

func (s Stanza) String() string {
	result := ""
	for _, l := range s {
		result += fmt.Sprintf("%s\n", l)
	}
	return result
}

func (p Poem) String() string {
	result := "PRINTING AS STRING:\n------------------\n"
	for _, s := range p {
		result += fmt.Sprintf("%s\n", s)
	}
	return result
}
