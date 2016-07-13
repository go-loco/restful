package rest

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	resp := Get(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestSlowGet(t *testing.T) {

	var f [100]*Response

	for i := range f {
		f[i] = rb.Get("/slow/user")

		if f[i].Response.StatusCode != http.StatusOK {
			t.Fatal("f Status != OK (200)")
		}

	}

}

func TestHead(t *testing.T) {
	resp := Head(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestPost(t *testing.T) {
	resp := Post(server.URL+"/user", &User{Name: "Matilda"})

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Status != OK (201)")
	}
}

func TestPostXML(t *testing.T) {

	rbXML := RequestBuilder{
		BaseURL:     server.URL,
		ContentType: XML,
	}

	resp := rbXML.Post("/xml/user", &User{Name: "Matilda"})

	if resp.StatusCode != http.StatusCreated {
		t.Fatal("Status != OK (201)")
	}
}

func TestPut(t *testing.T) {
	resp := Put(server.URL+"/user/3", &User{Name: "Pichucha"})

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestPatch(t *testing.T) {
	resp := Patch(server.URL+"/user/3", &User{Name: "Pichucha"})

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestDelete(t *testing.T) {
	resp := Delete(server.URL + "/user/4")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestOptions(t *testing.T) {
	resp := Options(server.URL + "/user")

	if resp.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200")
	}
}

func TestAsyncGet(t *testing.T) {

	r := <-AsyncGet(server.URL + "/user")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestAsyncHead(t *testing.T) {

	r := <-AsyncHead(server.URL + "/user")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestAsyncPost(t *testing.T) {

	r := <-AsyncPost(server.URL+"/user", &User{Name: "Matilda"})

	if r.StatusCode != http.StatusCreated {
		t.Fatal("Status != OK (201)")
	}

}

func TestAsyncPut(t *testing.T) {

	r := <-AsyncPut(server.URL+"/user/3", &User{Name: "Pichucha"})

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestAsyncPatch(t *testing.T) {

	r := <-AsyncPatch(server.URL+"/user/3", &User{Name: "Pichucha"})

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestAsyncDelete(t *testing.T) {

	r := <-AsyncDelete(server.URL + "/user/4")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestAsyncOptions(t *testing.T) {

	r := <-AsyncOptions(server.URL + "/user")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestHeaders(t *testing.T) {

	h := make(http.Header)
	h.Add("X-Test", "test")

	builder := RequestBuilder{
		BaseURL: server.URL,
		Headers: h,
	}

	r := builder.Get("/header")

	if r.StatusCode != http.StatusOK {
		t.Fatal("Status != OK (200)")
	}

}

func TestWrongURL(t *testing.T) {
	r := Get("foo")
	if r.Err == nil {
		t.Fatal("Wrong URL should get an error")
	}
}
