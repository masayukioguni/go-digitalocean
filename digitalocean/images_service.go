package digitalocean

import (
	"fmt"
	"net/http"
)

// ImagesService Digital Ocean API docs: https://developers.digitalocean.com/#images
type ImagesService struct {
	client  *Client
	Page    int
	PerPage int
}

// Image https://developers.digitalocean.com/#list-all-images
type Image struct {
	ID           int      `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Distribution string   `json:"distribution,omitempty"`
	Slug         string   `json:"slug,omitempty"`
	Public       bool     `json:"public,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	MinDiskSize  int      `json:"min_disk_size,omitempty"`
}

// ImagesResponse https://developers.digitalocean.com/#images
type ImagesResponse struct {
	Images []Image `json:"images,omitempty"`
	Meta   Meta    `json:"meta,omitempty"`
}

// List https://developers.digitalocean.com/#list-all-images
func (s *ImagesService) List(url string) ([]Image, *http.Response, error) {
	u := fmt.Sprintf("%spage=%d&perpage=%d", url, s.Page, s.PerPage)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	ir := new(ImagesResponse)
	resp, err := s.client.Do(req, ir)
	if err != nil {
		return nil, resp, err
	}

	return ir.Images, resp, err
}

// ListAll https://developers.digitalocean.com/#list-all-images
func (s *ImagesService) ListAll() ([]Image, *http.Response, error) {
	return s.List("v2/images?")
}

// ListApplication https://developers.digitalocean.com/#list-all-application-images
func (s *ImagesService) ListApplication() ([]Image, *http.Response, error) {
	return s.List("v2/images?type=application&")
}

//ListDistribution https://developers.digitalocean.com/#list-all-distribution-images
func (s *ImagesService) ListDistribution() ([]Image, *http.Response, error) {
	return s.List("v2/images?type=distribution&")
}
