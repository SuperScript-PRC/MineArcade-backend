package packets_arcade

type PlayerMatchRequest struct {
	PlayerId string `json:"playerId"`
	MatchId  string `json:"matchId"`
}
