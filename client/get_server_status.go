package client

import (
	"encoding/json"
	"net/http"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetServerStatus() (response.GetServerStatusResponse, error) {
	var resp response.GetServerStatusResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathServerStatus, map[string]string{})

	if err != nil {
		return resp, err
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
