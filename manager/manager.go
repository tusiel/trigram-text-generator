package manager

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"../utils"
	"../weighting"
)

type MarkovManager struct {
	sync.RWMutex
	Chain map[string]map[string]*int
}

func New() *MarkovManager {
	return &MarkovManager{
		Chain: make(map[string]map[string]*int),
	}
}

func (markov *MarkovManager) Train(text string) error {
	var err error

	text, err = utils.RemoveUnwantedCharacters(text)
	if err != nil {
		return err
	}

	text = strings.TrimSpace(text)
	// Replace all new lines with a space unlimited (-1) times
	text = strings.Replace(text, "\n", " ", -1)

	words := strings.Fields(text)

	markov.Lock()
	defer markov.Unlock()

	for ix := range words {
		if ix >= len(words)-3 {
			break
		}

		key := fmt.Sprintf("%s %s", words[ix], words[ix+1])

		_, ok := markov.Chain[key]
		if !ok {
			markov.Chain[key] = make(map[string]*int)
		}

		_, ok = markov.Chain[key][words[ix+2]]
		if !ok {
			markov.Chain[key][words[ix+2]] = new(int)
		}

		*markov.Chain[key][words[ix+2]]++

	}

	return nil

}

func (markov *MarkovManager) getMapKeys() (keys []string) {
	markov.RLock()
	defer markov.RUnlock()

	for key := range markov.Chain {
		keys = append(keys, key)
	}

	return
}

func (markov *MarkovManager) getRandomStart() string {
	keys := markov.getMapKeys()

	rand.Seed(time.Now().UnixNano())

	return keys[rand.Intn(len(keys))]
}

func (markov *MarkovManager) getNext(current string) string {
	markov.RLock()
	defer markov.RUnlock()

	var choices []weighting.Choice
	for item, weight := range markov.Chain[current] {
		c := weighting.Choice{Item: item, Weight: *weight}
		choices = append(choices, c)
	}

	choice, _ := weighting.WeightedChoice(choices)

	return choice.Item.(string)
}

func (markov *MarkovManager) Generate(sentences int) (string, error) {
	var words []string
	words = append(words, strings.Split(markov.getRandomStart(), " ")...)

	counter := 0
	for {
		selection := strings.Join(words[len(words)-2:], " ")

		sentenceEnd, err := utils.IsSentenceEnd(selection)
		if err != nil {
			return "", err
		}

		if sentenceEnd {
			counter++

			if counter >= sentences {
				break
			}

			words = append(words, strings.Split(markov.getRandomStart(), " ")...)
		} else {
			words = append(words, markov.getNext(selection))
		}
	}

	return strings.Join(words, " "), nil
}
