package mailtrap

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetAllProjects(accountID string) (*[]Project, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/accounts/%s/projects", c.HostURL, accountID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(body, &projects)
	if err != nil {
		return nil, err
	}

	return &projects, nil
}

func (c *Client) GetProject(accountID string, projectID string) (*Project, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/accounts/%s/projects/%s", c.HostURL, accountID, projectID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) DeleteProject(accountID string, projectID string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/api/accounts/%s/projects/%s", c.HostURL, accountID, projectID), nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	if string(body) != "Deleted Project" {
		return errors.New(string(body))
	}

	return nil
}

func (c *Client) CreateProject(accountID string, p ProjectNew) (*Project, error) {
	rb, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/accounts/%s/projects", c.HostURL, accountID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *Client) UpdateProject(accountID string, p ProjectNew) (*Project, error) {
	rb, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/accounts/%s/projects", c.HostURL, accountID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	project := Project{}
	err = json.Unmarshal(body, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
