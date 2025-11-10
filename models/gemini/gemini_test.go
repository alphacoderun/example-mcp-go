package gemini

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPrompt tests the prompt function of the Gemini model.
func TestPrompt(t *testing.T) {
	assert := assert.New(t)

	userPrompt := "Explain how AI works in a few words"

	actualResponse, err := prompt(userPrompt)

	assert.Nil(err, "Error should be nil")
	assert.NotNil(actualResponse, "Response should not be nil")
	assert.Greater(len(actualResponse), 0, "Response should not be empty")
}
