package d2prox

import (
	"encoding/hex"
	"fmt"
	"time"
)

// GameSession bundles relevant information about the game session, kept between proxies
type GameSession struct {
	RealmHost     string
	GameHost      string
	KeyHash       string
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

// Hash returns a unique hash of this key + account pair
func (s *GameSession) Hash() string {
	return fmt.Sprintf("%s%s", hex.EncodeToString([]byte(s.AccountName)), s.KeyHash)
}

// GameCounts returns the number of games joined in the past hour and the past 12 hours
func (s *GameSession) GameCounts() (int, int) {
	past1h := 0
	past12h := 0
	for _, game := range s.Games {
		since := time.Now().Sub(game.Start)
		if since < time.Hour {
			past1h++
		}
		if since < 12*time.Hour {
			past12h++
		}
	}
	return past1h, past12h
}

// MaxGamesPer1H is the maximum number of games within a 1 hour period
const MaxGamesPer1H = 20

// MaxGamesPer12H is the maximum number of games within a 12 hour period
const MaxGamesPer12H = 100

// GamesLeftUntilBan returns the number of games that can be created until the IP is temporarily banned
func (s *GameSession) GamesLeftUntilBan() int {
	last1h, last12h := s.GameCounts()
	left1h := MaxGamesPer1H - last1h
	left12h := MaxGamesPer12H - last12h
	if left1h < left12h {
		return left1h
	}
	return left12h
}
