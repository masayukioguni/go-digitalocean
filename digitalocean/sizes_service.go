package digitalocean

import (
	"net/http"
)

// SizesService Digital Ocean API docs: https://developers.digitalocean.com/#sizes
type SizesService struct {
	client *Client
}

// Size https://developers.digitalocean.com/#sizes
type Size struct {
	Slug         string   `json:"slug,omitempty"`
	Memory       int      `json:"memory,omitempty"`
	Vcpus        int      `json:"vcpus,omitempty"`
	Disk         int      `json:"disk,omitempty"`
	Transfer     float32  `json:"transfer,omitempty"`
	PriceMonthly float32  `json:"price_monthly,omitempty"`
	PriceHourly  float32  `json:"price_hourly,omitempty"`
	Regions      []string `json:"regions,omitempty"`
}

// SizesResponse https://developers.digitalocean.com/#list-all-sizes
type SizesResponse struct {
	Sizes []Size `json:"sizes,omitempty"`
	Meta  Meta   `json:"meta,omitempty"`
}

// List all Images
func (s *SizesService) List() ([]Size, *http.Response, error) {
	u := "v2/sizes"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	sr := new(SizesResponse)
	resp, err := s.client.Do(req, sr)
	if err != nil {
		return nil, resp, err
	}

	return sr.Sizes, resp, err
}
