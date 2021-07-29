package internal

import "water-jug/internal/models"

type Jug struct {
	name     string
	capacity int
	value    int
}

func NewJug(name string, capacity int) *Jug {
	return &Jug{name: name, capacity: capacity}
}

func (receiver *Jug) Fill() {
	receiver.value = receiver.capacity
}

func (receiver *Jug) Empty() {
	receiver.value = 0
}

func (receiver *Jug) Transfer(other *Jug) {
	amount := min(receiver.value, other.available())
	receiver.value -= amount
	other.value += amount
}

func (receiver Jug) available() int {
	return receiver.capacity - receiver.value
}

func (receiver Jug) IsEmpty() bool {
	return receiver.value == 0
}

func (receiver Jug) IsFull() bool {
	return receiver.value == receiver.capacity
}

func (receiver Jug) ToModel() models.JugModel {
	return models.JugModel{
		Name:     receiver.name,
		Capacity: receiver.capacity,
	}
}
