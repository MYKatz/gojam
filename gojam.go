package gojam

import (
	"fmt"
	"strings"

	"math/rand"
)

//These signal the start and end of a sentence.
const (
	StartToken = "_START_"
	EndToken   = "_END_"
)

type Markov struct {
	n         int    //the 'n' in n-gram, how long
	separator string //what separates each gram - for words, usually spaces.
	chain     map[string]*Output
}

type Output struct {
	occurrences int            //how many times the 'input' occurs. ie total of all output possibilities = 100%
	grams       map[string]int //how many times the word appears AFTER the input
}

func NewMarkov(grams int, separator string) *Markov {
	m := Markov{grams, separator, make(map[string]*Output)}
	return &m
}

func (m Markov) TrainOnExample(sentence string) { //for single example
	queue := []string{}
	for i := 0; i < m.n; i++ {
		queue = append(queue, "")
	}
	queue[len(queue)-1] = "_START_"
	words := strings.Fields(sentence)
	words = append(words, "_END_")
	for i := range words {
		st := words[i]
		prefix := strings.Join(queue, " ")
		opt := m.chain[prefix]
		if opt == nil {
			m.chain[prefix] = &Output{0, make(map[string]int)}
			opt = m.chain[prefix]
		}
		opt.grams[st] += 1
		opt.occurrences += 1
		//dequeue first element and shift everything
		queue[0] = "" //I've heard this helps with memory
		queue = queue[1:]
		queue = append(queue, st)
	}
}

func (m Markov) GenerateExample() string {
	queue := []string{}
	for i := 0; i < m.n; i++ {
		queue = append(queue, "")
	}
	generated := "_START_"
	sentence := []string{}
	for generated != "_END_" {
		queue[0] = ""
		queue = queue[1:]
		queue = append(queue, generated)
		outputs := m.chain[strings.Join(queue, " ")]
		num := rand.Intn(outputs.occurrences)
		tally := 0
		for k, v := range outputs.grams {
			tally += v
			if tally > num {
				sentence = append(sentence, k)
				generated = k
				break
			}
		}
	}
	sentence[len(sentence)-1] = ""
	return strings.Join(sentence, " ")
}

func (m Markov) PrintMap() {
	fmt.Println(m.chain)
	fmt.Println(m.chain["I"])
}
