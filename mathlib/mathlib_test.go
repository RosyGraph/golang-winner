package mathlib

import "testing"

func TestIntex(t *testing.T) {
	got := Intex()
	want := `\draw[dashed] (0,0) -- (0,1) -- (1,1);
\draw[dashed] (1,0) -- (1,4) -- (2,4);
\draw[dashed] (2,0) -- (2,9) -- (3,9);
\draw[dashed] (3,0) -- (3,9);
`
	if got != want {
		t.Errorf("got\n'%s'\nwant\n'%s'", got, want)
	}
}
