package d2prox

import "time"

// GameSession bundles relevant information about the game session, kept between proxies
type GameSession struct {
	RealmHost     string
	GameHost      string
	AccountName   string
	CharacterName string
	Game          *Game
	Games         []*Game
	Start         time.Time
	Debug         bool
}

// NewGameSession initializes a new game session
func NewGameSession(accountName string) *GameSession {
	return &GameSession{
		AccountName: accountName,
		Start:       time.Now(),
		Games:       make([]*Game, 0, 32),
	}
}
