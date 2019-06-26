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
	s := mark.GenerateExample()
	fmt.Println(s)
	fmt.Println(string(mark.ToJSON()))
}

func TestSerialization(t *testing.T) {
	mark := NewMarkov(1, " ")
	mark.TrainOnExample("I like to eat cats")
	mark.TrainOnExample("I like to jump rope")
	mark.TrainOnExample("jump man jump man jump man")
	mark.TrainOnExample("I don't like to jump on cats")
	mark.TrainOnExample("I once ate a bunch of cats who were orange")
	mark.TrainOnExample("I once died")
	mark.TrainOnExample("once I died in another dimension")
	serialized := mark.ToJSON()
	fmt.Println(string(serialized))
	markymark := NewMarkov(1, " ")
	err := markymark.FromJSON(serialized)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	markymark.PrintMap()
	for i := 0; i < 10; i++ {
		fmt.Println(markymark.GenerateExample())
	}
}
