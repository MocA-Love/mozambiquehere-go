package response

type GetMapRotationResponse struct {
	BattleRoyale ModeRotation `json:"battle_royale"`
	Ranked       ModeRotation `json:"ranked"`
	LTM          ModeRotation `json:"ltm"`
}

type ModeRotation struct {
	Current MapDetails `json:"current"`
	Next    MapDetails `json:"next"`
}

type MapDetails struct {
	Start            int64  `json:"start"`
	End              int64  `json:"end"`
	ReadableDateStart string `json:"readableDate_start"`
	ReadableDateEnd   string `json:"readableDate_end"`
	Map              string `json:"map"`
	Code             string `json:"code"`
	DurationInSecs   int64  `json:"DurationInSecs"`
	DurationInMinutes int64  `json:"DurationInMinutes"`
	Asset            string `json:"asset"`
	RemainingSecs    int64  `json:"remainingSecs,omitempty"`
	RemainingMins    int64  `json:"remainingMins,omitempty"`
	RemainingTimer   string `json:"remainingTimer,omitempty"`
	IsActive         bool   `json:"isActive,omitempty"`
	EventName        string `json:"eventName,omitempty"`
}
