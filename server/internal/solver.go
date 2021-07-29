package internal

import (
	"errors"
	"fmt"
	"water-jug/internal/models"
)

type Solver struct {
}

func NewSolver() *Solver {
	return &Solver{}
}

var NoSolution = errors.New("no solution")

const (
	transferTemplate = "Transfer from %s to %s"
	fillTemplate     = "Fill %s"
	emptyTemplate    = "Empty %s"
)

func (receiver *Solver) Do(x, y, z int) (models.Solution, error) {

	if isExpectedGreaterThanJugs(x, y, z) {
		return models.Solution{}, NoSolution
	}

	if isNotDividedBy(x, y, z) {
		return models.Solution{}, NoSolution
	}

	jugA := NewJug("A", x)
	jugB := NewJug("B", y)

	c := make(chan models.Solution, 2)
	defer close(c)
	go solve(*jugA, *jugB, z, c)
	go solve(*jugB, *jugA, z, c)

	resultA := <-c
	resultB := <-c

	if len(resultA.Steps) < len(resultB.Steps) {
		return resultA, nil
	} else {
		return resultB, nil
	}
}

func isNotDividedBy(x int, y int, z int) bool {
	aux := gcd(x, y)
	if aux == 0 {
		return true
	}
	return z%aux != 0
}

func isExpectedGreaterThanJugs(x int, y int, z int) bool {
	return z > x && z > y
}

func solve(left Jug, right Jug, expected int, c chan<- models.Solution) {

	result := models.Solution{Jugs: models.JugTuple{Left: left.ToModel(), Right: right.ToModel()}}

	for left.value != expected && right.value != expected {
		right.Transfer(&left)

		if right.IsEmpty() && left.IsEmpty() {
			result.AddStep("", left.value, right.value)
		} else {
			result.AddStep(fmt.Sprintf(transferTemplate, right.name, left.name), left.value, right.value)
		}

		if left.value == expected || right.value == expected {
			break
		}

		if right.IsEmpty() {
			right.Fill()
			result.AddStep(fmt.Sprintf(fillTemplate, right.name), left.value, right.value)
		}

		if left.IsFull() {
			left.Empty()
			result.AddStep(fmt.Sprintf(emptyTemplate, left.name), left.value, right.value)
		}

	}

	c <- result
}
