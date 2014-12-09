package digitalocean

import (
	"net/http"
)

// SSHKeyService Digital Ocean API docs: https://developers.digitalocean.com/#ssh-keys
type SSHKeysService struct {
	client *Client
}

// SSHKey https://developers.digitalocean.com/#list-all-keys
type SSHKey struct {
	ID          int    `json:"id,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
	Name        string `json:"name,omitempty"`
}

// SSHKey https://developers.digitalocean.com/#list-all-keys
type SSHKeyListResponse struct {
	SSHKeys []SSHKey `json:"ssh_keys,omitempty"`
	Meta    Meta     `json:"meta,omitempty"`
}

// List List all SSHKeys
func (s *SSHKeysService) List() ([]SSHKey, *http.Response, error) {
	u := "/v2/account/keys"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	sr := new(SSHKeyListResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}
	return sr.SSHKeys, resp, err
}
