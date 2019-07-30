/*
Package rabinkarp provides the Rabin-Karp string pattern matching algorithm.
*/
package rabinkarp

type (
	// RabinKarp is a function that runs the rabin karp algorithm
	RabinKarp func(string, string) (bool, Match)
)

// DefaultMatcher provides a instance
// of RabinKarp with the naive hasher
// implementation.
var DefaultMatcher = Make(NaiveHasher)

// Make a closure of RabinKarp which uses 
// the specified hasher.
func Make(h Hasher) RabinKarp {
	return func (p string, t string) (bool, Match) {
		var match Match
		ph := h.Hash(p)
		for i := 0; i < len(t) - len(p) + 1; i++ {
			st := t[i: i+len(p)]
			sth := h.Hash(st)
			if sth != ph {
				continue
			}
			for i, c := range p {
				if c != rune(st[i]) {
					continue
				}
			}
			match.Index = uint(i)
			match.Value = t
			return true, match
		}
		return false, match
	}
}

// MatchAll runs RabinKarp on every string in the slice 
// and stores all pattern matches in a matches slice and 
// returns it.
func (f RabinKarp) MatchAll(p string, v []string) Matches {
	var matches Matches
	for _, value := range v {
		if ok, match := f(p, value); ok {
			matches = append(matches, match)
		}
	}
	return matches
}
