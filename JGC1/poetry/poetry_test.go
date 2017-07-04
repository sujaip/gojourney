package poetry

//godoc testing
import (
	"testing"
)

/* We could also use t.Errorf in place of t.Fatalf in the test suite
   t.Errorf prints the error message and continues the testing
   t.Fatalf prints the message and terminates the testing */

func TestNumStanzas(t *testing.T) {
	p := Poem{{"And from my pillow, looking forth by light",
		"Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood",
		"Of Newton with his prism and silent face,",
		"The marble index of a mind for ever",
		"Voyaging through strange seas of Thought, alone."}}
	if p.NumStanzas() != 1 {
		t.Fatalf("Unexpected stanza count %d", p.NumStanzas())
	}

	p = Poem{}
	if p.NumStanzas() != 0 {
		t.Fatalf("Empty poem is not empty!")
	}
}

func TestNumLines(t *testing.T) {
	p := Poem{{"And from my pillow, looking forth by light",
		"Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood",
		"Of Newton with his prism and silent face,",
		"The marble index of a mind for ever",
		"Voyaging through strange seas of Thought, alone."}}
	if p.NumLines() != 6 {
		t.Fatalf("Unexpected line count %d", p.NumLines())
	}

	p = Poem{}
	if p.NumLines() != 0 {
		t.Fatalf("Empty poem is not empty!")
	}
}

func TestStats(t *testing.T) {
	p := Poem{}
	v, c, puncs := p.Stats()
	if v != 0 || puncs != 0 || c != 0 {
		t.Fatalf("Bad number of vowels, punctuations, or consonants (%d %d %d)", v, puncs, c)
	}

	p = Poem{{"Hello"}}
	v, c, puncs = p.Stats()
	if v != 2 || puncs != 0 || c != 3 {
		t.Fatalf("Bad number of vowels, punctuations, or consonants (%d %d %d)", v, puncs, c)
	}

	p = Poem{{"Hello, World!"}}
	v, c, puncs = p.Stats()
	if v != 3 || puncs != 3 || c != 7 {
		t.Fatalf("Bad number of vowels, punctuations, or consonants (%d %d %d)", v, puncs, c)
	}
}
