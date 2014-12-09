package digitalocean

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSSHKeysService_List(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/account/keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, mock.ReadJSON("list_all_keys.json"))
	})

	keys, _, _ := mock.Client.SSHKeysService.List()

	wantSSHKey := SSHKey{
		ID:          int(123),
		Fingerprint: string("fingerprint"),
		PublicKey:   string("public_key"),
		Name:        string("My SSH Public Key"),
	}

	if !reflect.DeepEqual(keys[0], wantSSHKey) {
		t.Errorf("SSHKeysService.List returned %+v, want %+v", keys[0], wantSSHKey)
	}
}

func TestSSHKeysService_Unauthorized(t *testing.T) {
	mock := NewMockClient()
	defer mock.Close()

	mock.Mux.HandleFunc("/v2/account/keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, mock.ReadJSON("unauthorized.json"))
	})

	_, hr, err := mock.Client.SSHKeysService.List()

	if !reflect.DeepEqual(hr.StatusCode, http.StatusUnauthorized) {
		t.Errorf("SSHKeysService.List returned %+v, want %+v", hr.StatusCode, http.StatusUnauthorized)
	}

	if err == nil {
		t.Errorf("SSHKeysService.List returned %+v, want erespor message.", err)
	}

}
