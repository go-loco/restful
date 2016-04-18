package rest

import (
	"fmt"
	"net/http"
	"testing"
)

func TestResponseBytesAndString(t *testing.T) {
	resp := Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	if string(resp.Bytes()) != resp.String() {
		t.Fatal("Bytes() and String() are not equal")
	}

}

func TestDebug(t *testing.T) {

	defer Debug(false)
	Debug(true)

	resp := Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

	fmt.Print(resp.Debug())

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

	var rbXML = RequestBuilder{
		BaseURL:     server.URL,
		ContentType: XML,
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
