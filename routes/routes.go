package routes

import (
	"io/ioutil"
	"net/http"

	"../config"
	"../manager"
)

var markov *manager.MarkovManager

func init() {
	markov = manager.New()
}

// HandleLearn takes plain text and adds it to the manager to learn
func HandleLearn(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = markov.Train(string(body))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

// HandleGenerate creates random text based off what it has learnt
func HandleGenerate(w http.ResponseWriter, r *http.Request) {
	result, err := markov.Generate(config.GetInt("output.sentences"))
	if err != nil {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Write([]byte(result))
	w.WriteHeader(http.StatusOK)
}
