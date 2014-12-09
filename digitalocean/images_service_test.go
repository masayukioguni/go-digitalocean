package digitalocean

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestImagesService_ListAll(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/images", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer test")

		fmt.Fprint(w, mock.ReadJSON("list_all_images.json"))
	})

	images, _, _ := mock.Client.ImagesService.ListAll()

	wantImage := Image{
		ID:           int(9999999),
		Name:         string("test_server"),
		Distribution: string("Ubuntu"),
		Slug:         string(""),
		Public:       bool(false),
		Regions: []string{"nyc1",
			"nyc1"},
		CreatedAt:   string("2014-10-16T16:16:53Z"),
		MinDiskSize: int(30),
	}

	if !reflect.DeepEqual(images[0], wantImage) {
		t.Errorf("ImagesService.List returned %+v, want %+v", images[0], wantImage)
	}

}

func TestImagesService_ListApplication(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/images", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer test")

		fmt.Fprint(w, mock.ReadJSON("list_all_images_application.json"))
	})

	images, _, _ := mock.Client.ImagesService.ListApplication()

	wantImage := Image{
		ID:           int(6376601),
		Name:         string("Ruby on Rails on 14.04 (Nginx + Unicorn)"),
		Distribution: string("Ubuntu"),
		Slug:         string("ruby-on-rails"),
		Public:       bool(true),
		Regions: []string{"nyc1",
			"ams1",
			"sfo1",
			"nyc2",
			"ams2",
			"sgp1",
			"lon1",
			"nyc3",
			"ams3",
			"nyc1"},
		CreatedAt:   string("2014-09-26T20:20:24Z"),
		MinDiskSize: int(20),
	}

	if !reflect.DeepEqual(images[0], wantImage) {
		t.Errorf("ImagesService.ListApplication returned %+v, want %+v", images[0], wantImage)
	}

}

func TestImagesService_ListDistribution(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/images", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Content-Type", "application/json")
		testHeader(t, r, "Authorization", "Bearer test")

		fmt.Fprint(w, mock.ReadJSON("list_all_images_distribution.json"))
	})

	images, _, _ := mock.Client.ImagesService.ListDistribution()

	wantImage := Image{
		ID:           int(6370882),
		Name:         string("20 x64"),
		Distribution: string("Fedora"),
		Slug:         string("fedora-20-x64"),
		Public:       bool(true),
		Regions: []string{"nyc1",
			"ams1",
			"sfo1",
			"nyc2",
			"ams2",
			"sgp1",
			"lon1",
			"nyc3",
			"ams3",
			"nyc3"},
		CreatedAt:   string("2014-09-26T15:29:01Z"),
		MinDiskSize: int(20),
	}

	if !reflect.DeepEqual(images[0], wantImage) {
		t.Errorf("ImagesService.ListDistribution returned %+v, want %+v", images[0], wantImage)
	}
}
