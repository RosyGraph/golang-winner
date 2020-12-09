package llist

import "testing"

func TestPush(t *testing.T) {
	tc := []struct {
		name string
		head *node
		data int
		want node
	}{
		{
			name: "create a new node",
			head: nil,
			data: 2,
			want: node{2, nil},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			got := push(c.head, c.data)
			if got != c.want {
				t.Errorf("got %v want %v", got, c.want)
			}

		})
	}

	t.Run("prepend a node to an existing node", func(t *testing.T) {
		n := node{0, nil}
		got := push(&n, 9)
		if got.next.data != n.data {
			t.Errorf("got %v want %v", got, n.data)
		}

	})
}
