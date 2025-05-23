package analysis_test

import (
	"testing"

	"github.com/pieroshka/gf/internal/analysis"
	"github.com/stretchr/testify/assert"
)

func TestLexical(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          []byte
		expectedOutput []byte
	}{}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := analysis.NewLexical()

			out := analyzer.Run([]byte(tt.input))
			assert.Equal(t, tt.expectedOutput, out)
		})
	}
}
