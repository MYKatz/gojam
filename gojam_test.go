package gojam

import (
	"fmt"
	"testing"
)

func TestTraining(t *testing.T) {
	mark := NewMarkov(1, " ")
	mark.TrainOnExample("I like to eat cats")
	mark.TrainOnExample("I like to jump rope")
	mark.TrainOnExample("jump man jump man jump man")
	mark.TrainOnExample("I don't like to jump on cats")
	mark.TrainOnExample("I once ate a bunch of cats who were orange")
	mark.TrainOnExample("I once died")
	mark.TrainOnExample("once I died in another dimension")
	mark.PrintMap()
	s := mark.GenerateExample()
	fmt.Println(s)
}
