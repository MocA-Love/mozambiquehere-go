package response

type PlatformRP struct {
	FoundRank            int    `json:"foundRank"`
	Val                  int    `json:"val"`
	UID                  string `json:"uid"`
	UpdateTimestamp      int64  `json:"updateTimestamp"`
	TotalMastersAndPreds int    `json:"totalMastersAndPreds"`
}

type RPReport struct {
	PC     PlatformRP `json:"PC"`
	PS4    PlatformRP `json:"PS4"`
	X1     PlatformRP `json:"X1"`
	Switch PlatformRP `json:"SWITCH"`
}

type GetPredatorResponse struct {
	RP RPReport `json:"RP"`
}
