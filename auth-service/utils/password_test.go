package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldGenerateHashFromPasswordModelWork(t *testing.T) {
	hashedPassword, err := GenerateHashFromPassword("test-password")

	assert.NoError(t, err)
	assert.True(t, ValidatePassword(hashedPassword, "test-password"))
}

func TestShouldValidatePasswordReturnTrueIfCorrectPassword(t *testing.T) {
	hashedPassword, err := GenerateHashFromPassword("test-password")

	assert.NoError(t, err)
	assert.True(t, ValidatePassword(hashedPassword, "test-password"))
}

func TestShouldValidatePasswordReturnFalseIfIncorrectPassword(t *testing.T) {
	hashedPassword, err := GenerateHashFromPassword("raw-password")

	assert.NoError(t, err)
	assert.True(t, ValidatePassword(hashedPassword, "raw-password"))
}
