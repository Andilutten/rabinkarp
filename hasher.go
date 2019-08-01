package rabinkarp

import "math"

type (
	// Hasher is used to hash values 
	// in a RabinKarp func
	Hasher interface {
		Hash(string) int
	}
	// Default hasher
	naiveHasher int
)

const (
	NaiveHasher = naiveHasher(0)
	base = float64(256)
)

func (naiveHasher) Hash(v string) int {
	var sum int
	for i, c := range v {
		sum += int(c) * int(math.Pow(base, float64(i)))
	}
	return sum
}

