package shuffler

import "math/rand"

var (
	counter int = 0
)

/* packages are initialized when go loads it
this involves setting up any global variables
Here the init function of "mayh/rand" is called before our init function
i.e. any init function in the imported package is called before our own init */
func init() {
	counter = 0
}

/* method to access our counter from outside the package */
func GetCount() int {
	return counter
}

type Shuffleable interface {
	Len() int
	Swap(i, j int)
}

type WeightedShuffleable interface {
	Shuffleable
	Weight(i int) int
}

func Shuffle(s Shuffleable) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		s.Swap(i, j)
		counter++
	}
}

func WeightedShuffle(w WeightedShuffleable) {
	/*initialize total to zero */
	total := 0
	/* calculate the total weight */
	for i := 0; i < w.Len(); i++ {
		total += w.Weight(i)
	}

	/* for each element */
	for i := 0; i < w.Len(); i++ {
		/* get a random number between zero and total weight */
		pos := rand.Intn(total)
		cum := 0
		/* for each element from i to the last element */
		for j := i; j < w.Len(); j++ {
			counter++
			/* add weight of element j to cum */
			cum += w.Weight(j)
			if pos < cum {
				total -= w.Weight(j)
				w.Swap(i, j)
				break
			}
		}
	}
}
