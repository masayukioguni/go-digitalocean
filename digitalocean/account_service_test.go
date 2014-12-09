package digitalocean

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("account.json"))
	})

	account, _, _ := mock.Client.AccountService.GetUserInformation()

	want := &Account{
		DropletLimit:  int(25),
		Email:         string("test@test.com"),
		UUID:          string("uuid"),
		EmailVerified: bool(true),
	}

	if !reflect.DeepEqual(account, want) {
		t.Errorf("AccountService.GetUserInformation returned %+v, want %+v", account, want)
	}
}

func TestAccountService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/account", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.AccountService.GetUserInformation()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusUnauthorized) {
		t.Errorf("AccountService.GetUserInformation returned %+v, want %+v", hr.StatusCode, http.StatusUnauthorized)
	}

	if err == nil {
		t.Errorf("AccountService.GetUserInformation returned %+v, want erespor message.", err)
	}

}
