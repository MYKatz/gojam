package gojam

import (
	"encoding/json"
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
	n         int               //the 'n' in n-gram, how long
	separator string            //what separates each gram - for words, usually spaces.
	Chain     map[string]Output `json:"_chain"`
}

type Output struct {
	Occurrences int            `json:"o"`     //how many times the 'input' occurs. ie total of all output possibilities = 100%
	Grams       map[string]int `json:"grams"` //how many times the word appears AFTER the input
}

func (self *Output) increment() {
	self.Occurrences++
}

func NewMarkov(grams int, separator string) *Markov {
	m := Markov{grams, separator, make(map[string]Output)}
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
		opt, exists := m.Chain[prefix]
		if !exists {
			m.Chain[prefix] = Output{0, make(map[string]int)}
			opt = m.Chain[prefix]
		}
		opt.Grams[st] += 1
		opt.increment()
		m.Chain[prefix] = opt
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
		outputs := m.Chain[strings.Join(queue, " ")]
		num := rand.Intn(outputs.Occurrences)
		tally := 0
		for k, v := range outputs.Grams {
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

func (m Markov) ToJSON() []byte {
	jsonString, _ := json.Marshal(m.Chain)
	return jsonString
}

func (m Markov) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &m.Chain)
	return err
}

func (m Markov) PrintMap() {
	fmt.Println(m.Chain)
}
