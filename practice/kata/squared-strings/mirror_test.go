package kata

import "testing"

func TestOper(t *testing.T) {
	t.Run("horizontal mirror", func(t *testing.T) {
		s := "abcd\nefgh\nijkl\nmnop"
		got := Oper(HorMirror, s)
		want := "mnop\nijkl\nefgh\nabcd"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("vertical mirror", func(t *testing.T) {
		s := "abcd\nefgh\nijkl\nmnop"
		got := VertMirror(s)
		want := "dcba\nhgfe\nlkji\nponm"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
