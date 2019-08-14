/*
Nim starts with a number of piles of stones. The piles may be unequal.
Players alternate picking up 1 or more stones from 1 pile.
The player to remove the last stone wins.

For example, there are n = 3 piles of stones having pile = [3, 2, 4] stones in them.
Play may proceed as follows:

Player  Takes           Leaving
                        pile=[3,2,4]
1       2 from pile[1]  pile=[3,4]
2       2 from pile[1]  pile=[3,2]
1       1 from pile[0]  pile=[2,2]
2       1 from pile[0]  pile=[1,2]
1       1 from pile[1]  pile=[1,1]
2       1 from pile[0]  pile=[0,1]
1       1 from pile[1]  WIN

Given the value of n and the number of stones in each pile, determine the
game's winner if both players play optimally.
*/
package main

import "testing"

// Generate int[][3 - 7] with random int values 1 - 10
func TestNim(t *testing.T) {
	// Create int[][3]
}
