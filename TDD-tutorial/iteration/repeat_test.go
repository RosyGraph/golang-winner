package iteration

import "testing"

func TestRepeat(t *testing.T) {

	assertEquals := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', want '%s'.", got, want)
		}
	}

	t.Run("Repeat character 5 times.", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEquals(t, got, want)
	})

	t.Run("Repeat character 1 times.", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assertEquals(t, got, want)
	})
}
