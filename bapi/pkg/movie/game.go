package movie

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/fermyon/spin/sdk/go/redis"
	utilrand "github.com/rajatjindal/bollywood-wasm/bapi/pkg/rand"
	"github.com/rajatjindal/bollywood-wasm/bapi/pkg/types"
)

func NewGame(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ng := &types.NewGame{}
	err = ng.UnmarshalJSON(b)
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rand.Seed(int64(ng.CreatedAt))
	id := fmt.Sprintf("%s-%s-%s", utilrand.String(ng.CreatedAt, 4), utilrand.String(ng.CreatedAt+1122335566, 4), utilrand.String(ng.CreatedAt+8899776523, 4))
	index := rand.Intn(len(hollywoodMovies))
	game := newGame(id, hollywoodMovies[index])

	b, err = game.MarshalJSON()
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := redis.Set(os.Getenv("REDIS_ADDRESS"), game.Game.Id, b); err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cors(w)
	b, err = game.Game.MarshalJSON()
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}

func MakeAGuess(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	guess := &types.Guess{}
	err = guess.UnmarshalJSON(b)
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id := guess.Id
	key := guess.Key
	rkey := []rune(key)[0]

	payload, err := redis.Get(os.Getenv("REDIS_ADDRESS"), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game := &types.GameWithMeta{}
	err = game.UnmarshalJSON(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if guessesRemaining(game) <= 0 {
		cors(w)
		http.Error(w, "game over", http.StatusBadRequest)
		return
	}

	makeAGuess(game, rkey)
	game.Game.MovieName = calculatedMovieName(game)
	game.Game.RemainingGuesses = guessesRemaining(game)
	game.Game.GuessedSuccessfully = hasGuessedMovieRight(game)

	d, err := game.MarshalJSON()
	if err != nil {
		cors(w)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := redis.Set(os.Getenv("REDIS_ADDRESS"), game.Game.Id, d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cors(w)
}

func GetGameState(w http.ResponseWriter, r *http.Request) {
	id := strings.ReplaceAll(r.URL.Path, "/bapi/game/", "")
	payload, err := redis.Get(os.Getenv("REDIS_ADDRESS"), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	game := &types.GameWithMeta{}
	err = game.UnmarshalJSON(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := game.Game.MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cors(w)
	w.Write(b)
}

func cors(w http.ResponseWriter) {
	w.Header().Set("access-control-allow-origin", "*")
	w.Header().Set("access-control-allow-methods", "OPTIONS, GET, POST")
	w.Header().Set("access-control-allow-headers", "authorization, content-type")
}
