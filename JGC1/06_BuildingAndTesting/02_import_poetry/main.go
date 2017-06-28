package main

import (
	"fmt"
	"poetry"
)

func main() {
	//p := poetry.NewPoem()  is same as p := poetry.Poem{}
	p := poetry.Poem{{"And from my pillow, looking forth by light", "Of moon or favouring stars, I could behold",
		"The antechapel where the statue stood", "Of Newton with his prism and silent face,",
		"The marble index of a mind for ever", "Voyaging through strange seas of Thought, alone."}}
	v, c, puncs := p.Stats()

	fmt.Printf("The poem has %d vowels, %d consonents, and %d punctuations\n", v, c, puncs)
	fmt.Printf("Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())
}
