package response

type GetCraftingRotationResponse []Bundle

type ItemType struct {
	Name      string `json:"name"`
	Rarity    string `json:"rarity"`
	Asset     string `json:"asset"`
	RarityHex string `json:"rarityHex"`
}

type BundleContent struct {
	Item     string   `json:"item"`
	Cost     int      `json:"cost"`
	ItemType ItemType `json:"itemType"`
}

type Bundle struct {
	Bundle        string          `json:"bundle"`
	Start         int64           `json:"start,omitempty"`
	End           int64           `json:"end,omitempty"`
	StartDate     string          `json:"startDate,omitempty"`
	EndDate       string          `json:"endDate,omitempty"`
	BundleType    string          `json:"bundleType"`
	BundleContent []BundleContent `json:"bundleContent"`
}
