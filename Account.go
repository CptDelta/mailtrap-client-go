package mailtrap

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetAllAccounts() (*[]Account, error) {
	const resource = "/api/accounts"
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, resource), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var accounts []Account
	err = json.Unmarshal(body, &accounts)
	if err != nil {
		return nil, err
	}

	return &accounts, nil
}
