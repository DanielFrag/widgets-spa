package utils

import (
	"regexp"
	"testing"
)

func TestGenerateRandomAlphaNumericString(t *testing.T) {
	t.Run("ValidLength", func(t *testing.T) {
		r, _ := regexp.Compile(`^(\d|[A-Za-z])*$`)
		session := GenerateRandomAlphaNumericString(8)
		if len(session) != 8 {
			t.Error("Generated string has an incorrect length")
		}
		if !r.MatchString(session) {
			t.Error("Invalid char finded")
		}
	})
	t.Run("InvalidLength", func(t *testing.T) {
		session := GenerateRandomAlphaNumericString(0)
		if session != "" {
			t.Error("Generated string must be an empty string")
		}
	})
}
