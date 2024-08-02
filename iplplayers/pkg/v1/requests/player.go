package requests

type PlayerAddRequest struct {
	PlayerName     string `json:"player_name" binding:"required,min=3,max=50"`
	PlayerDob      string `json:"player_dob" binding:"required,min=10,max=10"` // 12-01-2001
	PlayerCountry  string `json:"player_country" binding:"required,min=2,max=45"`
	PlayerCategory string `json:"player_category" binding:"required,min=4,max=45"`
}

type PlayerUpdateRequest struct {
	PlayerName     string `json:"player_name" binding:"required,min=3,max=50"`
	PlayerDob      string `json:"player_dob" binding:"required,min=10,max=10"` // 12-01-2001
	PlayerCountry  string `json:"player_country" binding:"required,min=2,max=45"`
	PlayerCategory string `json:"player_category" binding:"required,min=4,max=45"`
	PlayerStatus   uint8  `json:"player_status" binding:"required,min=0,max=8"`
}
