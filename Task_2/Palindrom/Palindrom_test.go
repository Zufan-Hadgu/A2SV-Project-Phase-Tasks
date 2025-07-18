package main
import "testing"

func TestReverse(t *testing.T) {
    input := "hello"
    expected := "olleh"
    result := reverse(input)

    if result != expected {
        t.Errorf("reverse(%q) = %q; want %q", input, result, expected)
    }
}