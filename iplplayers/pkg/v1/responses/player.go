package responses

import (
	"time"
)

type PlayerResponse struct {
	ID             int       `json:"-"`
	PlayerCode     string    `json:"player_code"`
	PlayerName     string    `json:"player_name"`
	PlayerDob      string    `json:"player_dob"`
	PlayerCountry  string    `json:"player_country"`
	PlayerCategory string    `json:"player_category"`
	Status         uint8     `json:"-"`
	StatusOut      string    `json:"status"`
	Age            int8      `json:"age"`
	CreatedDt      time.Time `json:"created_dt"`
	UpdatedDt      time.Time `json:"updated_dt"`
}

var _table_players = "players"

// TableName get sql table name players
func (m *PlayerResponse) TableName() string {
	return _table_players
}
