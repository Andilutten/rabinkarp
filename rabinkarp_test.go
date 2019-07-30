package rabinkarp

import (
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestRabinKarp(t *testing.T) {

	t.Run("Basic match", func(t *testing.T) {
		pattern := "abc"
		text := "dhfkasabcdef"

		matcher := Make(NaiveHasher)
		ok, match := matcher(pattern, text)

		assert.Assert(t, ok)
		assert.Equal(t, match.Index, uint(6))
	})

	t.Run("Match, filter and sort", func(t *testing.T) {
		pattern := "bro"
		inputs := []string{
			"slowbro",
			"browser",
			"bob",
			"st",
			"bebrobaum",
		}
		expected := []string{
			"browser",
			"bebrobaum",
			"slowbro",
		}
		matcher := Make(NaiveHasher)
		matches := matcher.MatchAll(pattern, inputs)

		assert.Equal(t, len(matches), len(expected))

		sort.Sort(matches)

		for i, match := range matches {
			assert.Equal(t, match.Value, expected[i])
		}
	})
}
