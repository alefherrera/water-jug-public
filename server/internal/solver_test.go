package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolver_Do(t *testing.T) {
	solver := NewSolver()
	t.Run("Success", func(t *testing.T) {
		t.Run("3_5_4", func(t *testing.T) {
			solution, err := solver.Do(3, 5, 4)
			assert.NoError(t, err)
			assert.Len(t, solution.Steps, 7)
			assert.Equal(t, 0, solution.Steps[0].Left)
			assert.Equal(t, 0, solution.Steps[0].Right)
			assert.Equal(t, 0, solution.Steps[1].Left)
			assert.Equal(t, 5, solution.Steps[1].Right)
			assert.Equal(t, 3, solution.Steps[2].Left)
			assert.Equal(t, 2, solution.Steps[2].Right)
			assert.Equal(t, 0, solution.Steps[3].Left)
			assert.Equal(t, 2, solution.Steps[3].Right)
			assert.Equal(t, 2, solution.Steps[4].Left)
			assert.Equal(t, 0, solution.Steps[4].Right)
			assert.Equal(t, 2, solution.Steps[5].Left)
			assert.Equal(t, 5, solution.Steps[5].Right)
			assert.Equal(t, 3, solution.Steps[6].Left)
			assert.Equal(t, 4, solution.Steps[6].Right)
		})
		t.Run("3_7_4", func(t *testing.T) {
			solution, err := solver.Do(3, 7, 4)
			assert.NoError(t, err)
			assert.Len(t, solution.Steps, 3)
			assert.Equal(t, 0, solution.Steps[0].Left)
			assert.Equal(t, 0, solution.Steps[0].Right)
			assert.Equal(t, 0, solution.Steps[1].Left)
			assert.Equal(t, 7, solution.Steps[1].Right)
			assert.Equal(t, 3, solution.Steps[2].Left)
			assert.Equal(t, 4, solution.Steps[2].Right)
		})
	})
	t.Run("No solution", func(t *testing.T) {
		t.Run("Zero", func(t *testing.T) {
			solution, err := solver.Do(0, 0, 0)
			assert.EqualError(t, err, NoSolution.Error())
			assert.Len(t, solution.Steps, 0)
		})
		t.Run("Greater than both", func(t *testing.T) {
			solution, err := solver.Do(3, 5, 6)
			assert.EqualError(t, err, NoSolution.Error())
			assert.Len(t, solution.Steps, 0)
		})
		t.Run("No gcd", func(t *testing.T) {
			solution, err := solver.Do(3, 6, 4)
			assert.EqualError(t, err, NoSolution.Error())
			assert.Len(t, solution.Steps, 0)
		})
	})
}
