package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaaa"

	assertCorrectMessage(t, expected, repeated)
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
func TestExampleRepeat(t *testing.T) {
	repeated := ExampleRepeat("n", 7)
	expected := "nnnnnnn"

	assertCorrectMessage(t, expected, repeated)
}
func assertCorrectMessage(t testing.TB, want, got string) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %q and got %q", want, got)
	}
}
