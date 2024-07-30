package client

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/models"
	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) AddMatchHistoryUsers(platform models.Platform, usernames []string) (response.GetMatchHistoryInfoResponse, error) {
	var resp response.GetMatchHistoryInfoResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathBridge, map[string]string{
		"platform": string(platform),
		"player":   strings.Join(usernames, ","),
		"action":   "add",
		"history":  "1",
	})

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
