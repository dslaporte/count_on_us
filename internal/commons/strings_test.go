package commons

import (
	commons "count_on_us/internal/commons/strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsModule(t *testing.T) {
	type testCases struct {
		description  string
		input        any
		expectOutput any
	}

	tests := []testCases{
		{
			description:  "should return true when string is empty",
			input:        "",
			expectOutput: true,
		},
		{
			description:  "should return false when string is not empty",
			input:        "test",
			expectOutput: false,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := commons.IsEmpty(test.input.(string))
			assert.Equal(t, test.expectOutput, result)
		})
	}
}
