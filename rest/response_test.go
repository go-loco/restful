package rest_test

import (
	"net/http"
	"testing"

	"github.com/go-loco/restful/rest"
)

func TestResponseBytesAndString(t *testing.T) {
	resp := rest.Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	if string(resp.Bytes()) != resp.String() {
		t.Fatal("Bytes() and String() are not equal")
	}

}

func TestGetFillUpJSON(t *testing.T) {

	var u []User

	resp := rb.Get("/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	err := resp.FillUp(&u)
	if err != nil {
		t.Fatal("Json fill up failed. Error: " + err.Error())
	}

	for _, v := range users {
		if v.Name == "Hernan" {
			return
		}
	}

	t.Fatal("Couldn't found Hernan")
}

func TestGetFillUpXML(t *testing.T) {

	var u []User

	var rbXML = rest.RequestBuilder{
		BaseURL:     server.URL,
		ContentType: rest.XML,
	}

	resp := rbXML.Get("/xml/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	err := resp.FillUp(&u)
	if err != nil {
		t.Fatal("Json fill up failed. Error: " + err.Error())
	}

	for _, v := range users {
		if v.Name == "Hernan" {
			return
		}
	}

	t.Fatal("Couldn't found Hernan")
}
