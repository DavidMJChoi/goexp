package word

import (
	"testing"
	// word "github.com/DavidMJChoi/goexp/src/goplch11/word1"
)

func TestPalindrome1(t *testing.T) {
	if !IsPalindrome("abba") {
		t.Error(`IsPalindrome("abba") = false`)
	}
	if !IsPalindrome("kayak") {
		t.Error(`IsPalindrome("kayak") = false`)
	}
}
func TestNonPalindrome1(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}
