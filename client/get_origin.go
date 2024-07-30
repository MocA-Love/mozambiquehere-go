package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MocA-Love/mozambiquehere-go/v4/domain/response"
)

func (c *clientImplementation) GetOrigin(username string) (response.GetUIDResponse, error) {
	var resp response.GetUIDResponse

	body, err := c.doEndpointRequest(http.MethodGet, pathOrigin, map[string]string{
		"player":   username,
	})

	if err != nil {
		return resp, err
	}

	fmt.Println(string(body))

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return resp, err
	}

	return resp, err
}
