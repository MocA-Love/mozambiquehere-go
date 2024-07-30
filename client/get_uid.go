package client

import (
	"encoding/json"
	"net/http"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/models"
	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetUIDByName(platform models.Platform, username string) (response.GetUIDResponse, error) {
	var resp response.GetUIDResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathNametoUID, map[string]string{
		"platform": string(platform),
		"player":   username,
	})

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
