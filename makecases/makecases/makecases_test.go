package makecases

import "testing"

func TestNewCase(t *testing.T) {
	t.Run("sample 0", func(t *testing.T) {
		got := NewCase("11 734 24")
		want := Case{11, 734, 24}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestNewCases(t *testing.T) {

}
