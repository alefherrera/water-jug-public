package models

type Input struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type JugModel struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type JugTuple struct {
	Left  JugModel `json:"left"`
	Right JugModel `json:"right"`
}

type Step struct {
	Name  string `json:"name"`
	Left  int    `json:"left"`
	Right int    `json:"right"`
}

type Solution struct {
	Jugs  JugTuple `json:"jugs"`
	Steps []Step   `json:"steps"`
}

func (receiver *Solution) AddStep(name string, left, right int) {
	receiver.Steps = append(receiver.Steps, Step{
		Left:  left,
		Right: right,
		Name:  name,
	})
}
