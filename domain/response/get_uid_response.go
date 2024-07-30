package response

type GetUIDResponse struct {
	Name   string `json:"name"`
	UID    string `json:"uid"`
	PID    string `json:"pid"`
	Avatar string `json:"avatar"`
}
