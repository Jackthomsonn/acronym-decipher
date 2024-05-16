package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackthomson/acronym-decipher/internal"
)

type StrategyUsed string

const (
	ExactMatch   StrategyUsed = "ExactMatch"
	ClosestMatch StrategyUsed = "ClosestMatch"
)

type Response struct {
	Word          string       `json:"word"`
	TimeToResolve int64        `json:"time_to_resolve"`
	StrategyUsed  StrategyUsed `json:"strategy_used"`
}

type Error struct {
	Message string `json:"message"`
}

type Api struct{}

func NewAPI() *Api {
	return &Api{}
}

func (a *Api) GetWord(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	query := r.URL.Query()
	acronym := query.Get("acronym")
	w.Header().Set("Content-Type", "application/json")

	if acronym == "" {
		w.WriteHeader(http.StatusBadRequest)
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.Encode(&Error{Message: "acronym query parameter is required"})
		return
	}

	exact, err := FindExactMatch(acronym)

	if err == nil {
		response := &Response{Word: exact, TimeToResolve: time.Since(start).Abs().Nanoseconds(), StrategyUsed: "ExactMatch"}
		w.WriteHeader(http.StatusOK)
		jsonEncoder := json.NewEncoder(w)
		jsonEncoder.Encode(response)
		return
	}

	closest := FindClosestMatch(acronym)

	response := &Response{Word: internal.GetAcronymValue(closest), TimeToResolve: time.Since(start).Abs().Nanoseconds(), StrategyUsed: "ClosestMatch"}

	w.WriteHeader(http.StatusOK)
	jsonEncoder := json.NewEncoder(w)
	jsonEncoder.Encode(response)
}

func FindExactMatch(acronym string) (string, error) {
	return internal.Exact(acronym)
}

func FindClosestMatch(acronym string) string {
	acronyms := internal.GetAcronyms()
	return internal.Closest(acronym, acronyms)
}
