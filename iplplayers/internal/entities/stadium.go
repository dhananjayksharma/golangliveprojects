package entities

import "time"

var _table_stadium = "stadiums"

type Stadiums struct {
	ID               int       `gorm:"primaryKey;autoIncrement"`
	StadiumCode      string    `gorm:"column:stadium_code"`
	StadiumName      string    `gorm:"column:stadium_name"`
	StadiumBuiltDate string    `gorm:"column:stadium_build_date"`
	StadiumCountry   string    `gorm:"column:stadium_country"`
	StadiumCategory  string    `gorm:"column:stadium_category"`
	StadiumCost      float64   `gorm:"column:stadium_cost"`
	StadiumStartDate time.Time `gorm:"column:stadium_start_date"`
	StadiumEndDate   time.Time `gorm:"column:stadium_end_date"`
	Status           uint8     `gorm:"column:status"`
	CreatedDt        time.Time `gorm:"column:created_dt;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedDt        time.Time `gorm:"column:updated_dt;type:datetime;default:CURRENT_TIMESTAMP;autoUpdateTime"`
}

// TableName get sql table name players
func (m *Stadiums) TableName() string {
	return _table_stadium
}
