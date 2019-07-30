package rabinkarp

type (
	// Hasher is used to hash values 
	// in a RabinKarp func
	Hasher interface {
		Hash(string) int
	}
	// Default hasher
	naiveHasher int
)

const NaiveHasher = naiveHasher(0)

func (naiveHasher) Hash(v string) int {
	var sum int
	for _, c := range v {
		sum += int(c)
	}
	return sum
}
