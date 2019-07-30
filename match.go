package rabinkarp

type (
	// Match is return when a RabinKarp
	// function founds a match
	Match struct {
		Index uint
		Value string
	}
	// Matches is a slice of match.
	// The type implements sort.Interface
	Matches []Match
)

func (m Matches) Len() int {
	return len(m)
}

func (m Matches) Less(i, j int) bool {
	return m[i].Index < m[j].Index
}

func (m Matches) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
