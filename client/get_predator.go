package client

import (
	"encoding/json"
	"net/http"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetPredator() (response.GetPredatorResponse, error) {
	var resp response.GetPredatorResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathPredator, map[string]string{})

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
