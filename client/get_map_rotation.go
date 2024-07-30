package client

import (
	"encoding/json"
	"net/http"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetMapRotation() (response.GetMapRotationResponse, error) {
	var resp response.GetMapRotationResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathMapRotation, map[string]string{
		"version": "2",
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
