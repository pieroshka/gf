package analysis_test

import (
	"testing"

	"github.com/pieroshka/gf/internal/analysis"
	"github.com/stretchr/testify/assert"
)

func TestSyntactical(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		input          []byte
		expectedOutput int
		expectedOk     bool
		expectedError  error
	}{
		{
			name:           "no brackets",
			input:          []byte("++--"),
			expectedOutput: 0,
			expectedOk:     true,
			expectedError:  nil,
		},
		{
			name:           "balanced brackets",
			input:          []byte("[++]--[--]"),
			expectedOutput: 0,
			expectedOk:     true,
			expectedError:  nil,
		},
		{
			name:           "unbalanced - missing closing",
			input:          []byte("[[++"),
			expectedOutput: 1,
			expectedOk:     false,
			expectedError:  analysis.ErrClosingBracketNotFound,
		},
		{
			name:           "unbalanced - extra closing",
			input:          []byte("++]"),
			expectedOutput: 2,
			expectedOk:     false,
			expectedError:  analysis.ErrUnexpectedClosingBracket,
		},
		{
			name:           "closing bracket appears first",
			input:          []byte("]+["),
			expectedOutput: 0,
			expectedOk:     false,
			expectedError:  analysis.ErrUnexpectedClosingBracket,
		},
		{
			name:           "nested balanced",
			input:          []byte("[[[]]]"),
			expectedOutput: 0,
			expectedOk:     true,
			expectedError:  nil,
		},
		{
			name:           "multiple valid sections",
			input:          []byte("[][++][--]"),
			expectedOutput: 0,
			expectedOk:     true,
			expectedError:  nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			analyzer := analysis.NewSyntactical()

			out, ok, err := analyzer.Run([]byte(tt.input))
			assert.Equal(t, tt.expectedOutput, out)
			assert.Equal(t, tt.expectedOk, ok)
			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func FuzzSyntactical(f *testing.F) {
	f.Add("....><.,.,.,.")
	f.Add("[[[")
	f.Add("]]]")
	f.Add(">>")
	f.Add("<<")

	f.Fuzz(func(t *testing.T, input string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic on input: %q", input)
			}
		}()

		t.Logf("testing: %s", input)

		analyzer := analysis.NewSyntactical()
		_, _, _ = analyzer.Run([]byte(input))
	})
}
