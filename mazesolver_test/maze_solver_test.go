package mazesolver_test

import (
	"testing"

	mb "github.com/devamondos/maze_solver/mazebuilder"
	ms "github.com/devamondos/maze_solver/mazesolver"
)

const testfile string = "../test.png"

var expectedMoves [9][2]int = [9][2]int{
	{0, 1}, {1, 1}, {1, 2}, {1, 3}, {2, 1}, {3, 1}, {3, 2}, {3, 3}, {4, 3},
}

func TestSolve(t *testing.T) {
	maze := mb.Build(testfile, false)
	solution := ms.Solve(maze, "alwaysLeft", false)

	if !solution.IsSuccessful {
		t.Error("Solution was not found")
	}
	if len(solution.Moves) != len(expectedMoves) {
		t.Errorf("Solution took longer than expected. Expected %d moves but was %d moves.", len(expectedMoves), len(solution.Moves))
	}
	for i := 0; i < len(expectedMoves); i++ {
		expectedMove := expectedMoves[i]
		actualMove := solution.Moves[i]
		if expectedMove[0] != actualMove[0] && expectedMove[1] != actualMove[1] {
			t.Errorf("Mismatch in expected moves. Expected (%d,%d) but was actually (%d,%d)", expectedMove[0], expectedMove[1], actualMove[0], actualMove[1])
		}
	}
}
