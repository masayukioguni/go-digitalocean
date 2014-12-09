package digitalocean

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSizesService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/sizes", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("size.json"))
	})

	sizes, _, _ := mock.Client.SizesService.List()

	wantLen := 9
	if !reflect.DeepEqual(len(sizes), wantLen) {
		t.Errorf("SizesService.List returned %+v, want %+v", len(sizes), wantLen)
	}

	wantSize := Size{
		Slug:         string("512mb"),
		Memory:       int(512),
		Vcpus:        int(1),
		Disk:         int(20),
		Transfer:     float32(1.0),
		PriceMonthly: float32(5.0),
		PriceHourly:  float32(0.00744),
		Regions: []string{"nyc1",
			"sgp1",
			"ams1",
			"ams2",
			"sfo1",
			"nyc2",
			"lon1",
			"nyc3",
			"ams3"},
	}

	if !reflect.DeepEqual(sizes[0], wantSize) {
		t.Errorf("SizesService.List returned %+v, want %+v", sizes[0], wantSize)
	}

}
