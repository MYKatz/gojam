package gojam

import (
	"testing"
)

func TestTraining(t *testing.T) {
	mark := NewMarkov(1, " ")
	mark.TrainOnExample("I like to eat cats")
	mark.TrainOnExample("I like to jump rope")
	mark.TrainOnExample("I don't like to jump on cats")
	mark.PrintMap()
}
