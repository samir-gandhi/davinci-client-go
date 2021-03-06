package davinci

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Gets an array of all roles for a company
func (c *Client) GetRoles(companyId *string, args *Params) (*[]Role, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/company/%s/roles", c.HostURL, cIdString), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, args)
	if err != nil {
		return nil, err
	}
	fmt.Println("PRINTING BODY")
	fmt.Println(body)

	roles := []Role{}
	err = json.Unmarshal(body, &roles)
	if err != nil {
		return nil, err
	}

	return &roles, nil
}


//Create a bare role, policies can be added _after_ creation
func (c *Client) CreateRole(companyId *string, payload *RoleCreate) (*RoleCreateResponse, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if payload == nil {
		return nil, fmt.Errorf("payload not provided")
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/company/%s/roles", c.HostURL, cIdString), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	role := RoleCreateResponse{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

// Update a previously created role
func (c *Client) UpdateRole(companyId *string, roleName *string, payload *RoleUpdate) (*Role, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if roleName == nil {
		return nil, fmt.Errorf("roleName not provided")
	}

	uIdString := *roleName

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/company/%s/roles/%s", c.HostURL, cIdString, uIdString), strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	role := Role{}
	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}

	return &role, nil
}

//Delete a role from a company
func (c *Client) DeleteRole(companyId, roleName *string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	cIdString := *cIdPointer

	if roleName == nil {
		return nil, fmt.Errorf("customerId not provided")
	}

	uIdString := *roleName
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/company/%s/roles/%s", c.HostURL, cIdString, uIdString), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	msg := Message{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}