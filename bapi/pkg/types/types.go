package types

const MaxGuesses = 9

type GuessResult int

const (
	WrongGuess GuessResult = 0
	RightGuess GuessResult = 1
	NotGuessed GuessResult = -1
)

//easyjson:json
type NewGame struct {
	CreatedAt int64 `json:"createdAt"`
}

//easyjson:json
type Guess struct {
	Id  string `json:"id"`
	Key string `json:"key"`
}

//easyjson:json
type GameWithMeta struct {
	Game *Game `json:"game"`
	Meta *Meta `json:"meta"`
}

//easyjson:json
type Meta struct {
	RawMovieName        string
	NormalizedMovieName string
	MovieNameMap        map[rune]GuessResult
}

//easyjson:json
type Game struct {
	Id                  string               `json:"id"`
	GuessMap            map[rune]GuessResult `json:"guessMap"`
	RemainingGuesses    int                  `json:"remainingGuesses"`
	GuessedSuccessfully bool                 `json:"guessedSuccessfully"`
	MovieName           string               `json:"movieName"`
	Plot                string               `json:"plot"`
}
