package digitalocean

import (
	"net/http"
)

// AccountService Digital Ocean API docs: https://developers.digitalocean.com/#account
type AccountService struct {
	client *Client
}

// Account represents a Digital Ocean acount.
type Account struct {
	DropletLimit  int    `json:"droplet_limit,omitempty"`
	Email         string `json:"email,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	EmailVerified bool   `json:"email_verified,omitempty"`
}

// AccountResponse https://developers.digitalocean.com/#account
type AccountResponse struct {
	Account Account `json:"account"`
}

// GetUserInformation https://developers.digitalocean.com/#account
func (s *AccountService) GetUserInformation() (*Account, *http.Response, error) {
	u := "v2/account"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ar := new(AccountResponse)
	resp, err := s.client.Do(req, ar)
	if err != nil {
		return nil, resp, err
	}
	return &ar.Account, resp, err
}
