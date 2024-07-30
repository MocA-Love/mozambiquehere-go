package response

type ServerStatus struct {
	Status         string `json:"Status"`
	HTTPCode       int    `json:"HTTPCode"`
	ResponseTime   int    `json:"ResponseTime"`
	QueryTimestamp int64  `json:"QueryTimestamp"`
}

type SelfCoreTest struct {
	StatusWebsite   ServerStatus `json:"Status-website"`
	StatsAPI        ServerStatus `json:"Stats-API"`
	Overflow1       ServerStatus `json:"Overflow-#1"`
	Overflow2       ServerStatus `json:"Overflow-#2"`
	OriginAPI       ServerStatus `json:"Origin-API"`
	PlaystationAPI  ServerStatus `json:"Playstation-API"`
	XboxAPI         ServerStatus `json:"Xbox-API"`
}

type PlatformStatus struct {
	Status         string `json:"Status"`
	QueryTimestamp int64  `json:"QueryTimestamp"`
}

type GetServerStatusResponse struct {
	OriginLogin     map[string]ServerStatus    `json:"Origin_login"`
	EANovaFusion    map[string]ServerStatus    `json:"EA_novafusion"`
	EAAccounts      map[string]ServerStatus    `json:"EA_accounts"`
	ApexOauthCrossplay map[string]ServerStatus `json:"ApexOauth_Crossplay"`
	SelfCoreTest    SelfCoreTest               `json:"selfCoreTest"`
	OtherPlatforms  map[string]PlatformStatus  `json:"otherPlatforms"`
}
