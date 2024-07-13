package entities

import (
	"time"
)

var _table_players = "players"

// Players represents the player for this application

type Players struct {
	ID             int       `gorm:"primaryKey;autoIncrement"`
	PlayerCode     string    `gorm:"column:player_code"`
	PlayerName     string    `gorm:"column:player_name"`
	PlayerDob      string    `gorm:"column:player_dob"`
	PlayerCountry  string    `gorm:"column:player_country"`
	PlayerCategory string    `gorm:"column:player_category"`
	Status         uint8     `gorm:"column:status"`
	CreatedDt      time.Time `gorm:"column:created_dt;type:TIMESTAMP WITH TIME ZONE; DEFAULT CURRENT_TIMESTAMP"`
	UpdatedDt      time.Time `gorm:"column:updated_dt;type:TIMESTAMP WITH TIME ZONE; DEFAULT CURRENT_TIMESTAMP; ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName get sql table name players
func (m *Players) TableName() string {
	return _table_players
}
