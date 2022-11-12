package movie

import (
	"fmt"
	"strings"

	"github.com/rajatjindal/bollywood-wasm/bapi/pkg/types"
)

func newGame(id, movieName string) *types.GameWithMeta {
	g := &types.GameWithMeta{
		Game: &types.Game{
			Id:               id,
			GuessMap:         map[rune]types.GuessResult{},
			RemainingGuesses: types.MaxGuesses,
		},
		Meta: &types.Meta{
			RawMovieName:        movieName,
			MovieNameMap:        map[rune]types.GuessResult{},
			NormalizedMovieName: nomalize(movieName),
		},
	}

	for _, c := range g.Meta.NormalizedMovieName {
		if c == ' ' {
			continue
		}

		g.Meta.MovieNameMap[c] = types.NotGuessed
	}

	g.Game.MovieName = calculatedMovieName(g)
	return g
}

func makeAGuess(g *types.GameWithMeta, s rune) {
	if _, exists := g.Meta.MovieNameMap[s]; exists {
		g.Game.GuessMap[s] = types.RightGuess
		g.Meta.MovieNameMap[s] = types.RightGuess
		return
	}

	g.Game.GuessMap[s] = types.WrongGuess
}

func guessesRemaining(g *types.GameWithMeta) int {
	remaining := types.MaxGuesses
	for _, v := range g.Game.GuessMap {
		if v == types.RightGuess {
			continue
		}

		remaining--
	}

	return remaining
}

func hasGuessedMovieRight(g *types.GameWithMeta) bool {
	for k := range g.Meta.MovieNameMap {
		if _, ok := g.Game.GuessMap[k]; !ok {
			fmt.Printf("chat %c does not exist\n", k)
			return false
		}
	}

	return true
}

func calculatedMovieName(g *types.GameWithMeta) string {
	name := []rune{}
	for index, k := range g.Meta.NormalizedMovieName {
		if _, ok := g.Game.GuessMap[k]; ok {
			name = append(name, []rune(g.Meta.RawMovieName)[index])
			continue
		}

		if k == ' ' {
			name = append(name, ' ')
			continue
		}

		name = append(name, '-')
	}

	return string(name)
}

func nomalize(name string) string {
	return strings.ToLower(name)
}
